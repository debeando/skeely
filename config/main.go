package config

import (
	"errors"
	"fmt"
	"io/ioutil"
	"os"

	"mylinter/common"

	"github.com/go-yaml/yaml"
)

func (c *Config) Load() error {
	source, err := ioutil.ReadFile(".mylinter.yaml")
	if err != nil {
		return nil
	}

	source = []byte(os.ExpandEnv(string(source)))

	if err := yaml.Unmarshal(source, &c); err != nil {
		errors.New(fmt.Sprintf("Imposible to parse config file - %s", err))
	}

	return nil
}

func (c *Config) IgnoreCodes(table_name string) (errors []int) {
	errors = append(errors, c.getIgnoreByTable(table_name)...)
	errors = append(errors, c.getIgnoreGeneral()...)
	errors = common.UnduplicateArrayInt(errors)

	return errors
}

func (c *Config) getIgnoreByTable(table_name string) (errors []int) {
	for index := range c.Tables {
		if c.Tables[index].Name == table_name {
			return common.StringToArrayInt(c.Tables[index].Ignore)
		}
	}

	return errors
}

func (c *Config) getIgnoreGeneral() []int {
	return common.StringToArrayInt(c.Ignore)
}
