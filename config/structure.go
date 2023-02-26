package config

type Config struct {
	Ignore           string  `yaml:"ignore,omitempty"`
	FieldsMax        int     `yaml:"fields-max,omitempty"`
	CharLengthMax    int     `yaml:"char-length-max,omitempty"`
	VarcharLengthMax int     `yaml:"varchar-mength-max,omitempty"`
	Tables           []Table `yaml:"tables,omitempty"`
}

type Table struct {
	Name   string `yaml:"name"`
	Ignore string `yaml:"ignore"`
}
