package config

type Config struct {
	Server Server `yaml:"server"`
}

type Server struct {
	Env       string     `yaml:"environment"`
	Ports     []Port     `yaml:"ports"`
	Endpoints []Endpoint `yaml:"endpoints"`
	Timeouts  []Timeout  `yaml:"timeouts"`
}

type Port struct {
	Name string `yaml:"name"`
	Port int64  `yaml:"port"`
}

type Endpoint struct {
	Name       string `yaml:"name"`
	TlsEnabled bool   `yaml:"tlsEnabled"`
	EnvType    string `yaml:"envType"`
}

type Timeout struct {
	Name  string `yaml:"name"`
	Value int  `yaml:"value"`
}
