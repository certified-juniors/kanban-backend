package config

type HTTPServer struct {
	Address      string `yaml:"address"`
	ReadTimeout  uint   `yaml:"read_timeout"`
	WriteTimeout uint   `yaml:"write_timeout"`
	IdleTimeout  uint   `yaml:"idle_timeout"`
}
