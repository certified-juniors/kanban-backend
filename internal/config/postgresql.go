package config

type PostgresConfig struct {
	URL         string `yaml:"url"`
	AutoMigrate bool   `yaml:"auto_migrate"`
	Migrations  string `yaml:"migrations"`
}
