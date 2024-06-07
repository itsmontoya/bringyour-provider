package plugin

import (
	"context"
	"fmt"
	"log"

	"github.com/vroomy/vroomy"

	"github.com/itsmontoya/bringyour-provider/libs/provider"
)

var p Plugin

func init() {
	if err := vroomy.Register("provider", &p); err != nil {
		log.Fatal(err)
	}
}

type Plugin struct {
	vroomy.BasePlugin

	// Backend
	provider *provider.Provider
}

// Load will load up the underlying controller
func (p *Plugin) Load(env vroomy.Environment) (err error) {
	var o provider.Options
	o.Username = env["username"]
	o.Password = env["password"]
	if o.Port, err = parseInt(env["port"]); err != nil {
		err = fmt.Errorf("error parsing port")
		return
	}

	if p.provider, err = provider.New(context.Background(), o); err != nil {
		return
	}

	fmt.Printf("Client ID: %s\n", p.provider.ClientID())
	fmt.Printf("Instance ID: %s\n", p.provider.InstanceID())
	return
}

// Backend exposes this plugin's backend layer to other plugins
func (p *Plugin) Backend() interface{} {
	return p.provider
}

// Close is called when vroomy service is terminating
func (p *Plugin) Close() (err error) {
	return p.provider.Close()
}
