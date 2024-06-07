package provider

type Options struct {
	ApiURL     string
	ConnectURL string
	Port       int

	Username string
	Password string
}

func (o *Options) Fill() {
	if len(o.ApiURL) == 0 {
		o.ApiURL = DefaultApiURL
	}

	if len(o.ConnectURL) == 0 {
		o.ConnectURL = DefaultConnectURL
	}

	if o.Port == 0 {
		o.Port = 80
	}
}
