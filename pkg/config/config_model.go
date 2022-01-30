package config

// Config is central container for configuration object
type Config struct {
	Server Server `yaml:"server"`
}

// Server is the container for server-related configuration
type Server struct {
	Env       string     `yaml:"environment"`
	Ports     []Port     `yaml:"ports"`
	Endpoints []Endpoint `yaml:"endpoints"`
	Timeouts  []Timeout  `yaml:"timeouts"`
}

// Port is an entry in Ports for a named-port
type Port struct {
	Name string `yaml:"name"`
	Port int64  `yaml:"port"`
}

// Endpoint is an entry in Endpoints for qualified endpoints
type Endpoint struct {
	Name       string `yaml:"name"`
	TLSEnabled bool   `yaml:"tlsEnabled"`
	EnvType    string `yaml:"envType"`
}

// Timeout is an entry in Timeouts for server-timeouts
type Timeout struct {
	Name  string `yaml:"name"`
	Value int    `yaml:"value"`
}
