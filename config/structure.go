package config

type Config struct {
	Ignore string  `yaml:"ignore,omitempty"`
	Tables []Table `yaml:"tables,omitempty"`
}

type Table struct {
	Name   string `yaml:"name"`
	Ignore string `yaml:"ignore"`
}
