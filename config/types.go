package config

// Config TODO
type Config struct {
	Stacks []Stack `yaml:"stacks"`
}

// Stack TODO
type Stack struct {
	Name           string       `yaml:"name"`
	Gateway        string       `yaml:"gateway"`
	GatewayAuth    EndpointAuth `yaml:"gatewayauth,omitempty"`
	Prometheus     string       `yaml:"prometheus"`
	PrometheusAuth EndpointAuth `yaml:"prometheusauth,omitempty"`
}

// EndpointAuth TODO
type EndpointAuth struct {
	Token string `yaml:"token,omitempty"`
}
