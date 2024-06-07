package provider

import (
	"context"
	"errors"
	"fmt"

	gojwt "github.com/golang-jwt/jwt/v5"

	"bringyour.com/connect"
	"bringyour.com/protocol"
)

const DefaultApiURL = "https://api.bringyour.com"
const DefaultConnectURL = "wss://connect.bringyour.com"
const LocalVersion = "0.0.0-local"

func New(ctx context.Context, o Options) (out *Provider, err error) {
	var p Provider
	p.o = o
	p.o.Fill()

	fmt.Println("Opts", p.o)
	if err = p.authenticate(ctx); err != nil {
		return
	}

	p.instanceID = connect.NewId()
	clientOob := connect.NewApiOutOfBandControl(ctx, p.byClientJWT, o.ApiURL)
	p.client = connect.NewClientWithDefaults(ctx, p.clientID, clientOob)

	auth := &connect.ClientAuth{
		ByJwt:      p.byClientJWT,
		InstanceId: p.instanceID,
		AppVersion: RequireVersion(),
	}

	connect.NewPlatformTransportWithDefaults(ctx, o.ConnectURL, auth, p.client.RouteManager())
	p.localNAT = connect.NewLocalUserNatWithDefaults(ctx, p.clientID.String())
	p.remoteNAT = connect.NewRemoteUserNatProviderWithDefaults(p.client, p.localNAT)

	pm := provideModes{
		protocol.ProvideMode_Public:  true,
		protocol.ProvideMode_Network: true,
	}
	p.client.ContractManager().SetProvideModes(pm)
	out = &p
	return
}

type Provider struct {
	o Options

	byClientJWT string
	clientID    connect.Id
	instanceID  connect.Id

	localNAT  *connect.LocalUserNat
	remoteNAT *connect.RemoteUserNatProvider
	client    *connect.Client
}

func (p *Provider) ClientID() string {
	return p.clientID.String()
}

func (p *Provider) InstanceID() string {
	return p.instanceID.String()
}

func (p *Provider) Close() (err error) {
	p.remoteNAT.Close()
	p.localNAT.Close()
	p.client.Cancel()
	return
}

func (p *Provider) authenticate(ctx context.Context) (err error) {
	api := connect.NewBringYourApiWithContext(ctx, p.o.ApiURL)
	loginCallback, loginChannel := connect.NewBlockingApiCallback[*connect.AuthLoginWithPasswordResult]()
	loginArgs := &connect.AuthLoginWithPasswordArgs{
		UserAuth: p.o.Username,
		Password: p.o.Password,
	}

	api.AuthLoginWithPassword(loginArgs, loginCallback)
	loginResult := <-loginChannel

	switch {
	case loginResult.Error != nil:
		return loginResult.Error
	case loginResult.Result.Error != nil:
		return errors.New(loginResult.Result.Error.Message)
	}

	api.SetByJwt(loginResult.Result.Network.ByJwt)
	authClientCallback, authClientChannel := connect.NewBlockingApiCallback[*connect.AuthNetworkClientResult]()
	authClientArgs := &connect.AuthNetworkClientArgs{
		Description: fmt.Sprintf("provider %s", RequireVersion()),
		DeviceSpec:  "",
	}

	api.AuthNetworkClient(authClientArgs, authClientCallback)
	authClientResult := <-authClientChannel

	switch {
	case authClientResult.Error != nil:
		return loginResult.Error
	case authClientResult.Result.Error != nil:
		return errors.New(authClientResult.Result.Error.Message)
	}

	p.byClientJWT = authClientResult.Result.ByClientJwt
	parser := gojwt.NewParser()
	token, _, err := parser.ParseUnverified(p.byClientJWT, gojwt.MapClaims{})
	if err != nil {
		err = fmt.Errorf("error parsing client JWT: %v", err)
		return
	}

	claims := token.Claims.(gojwt.MapClaims)
	p.clientID, err = connect.ParseId(claims["client_id"].(string))
	return
}
