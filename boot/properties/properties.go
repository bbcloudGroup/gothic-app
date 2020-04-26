package properties

import (
	"github.com/bbcloudGroup/gothic/config"
	"gopkg.in/yaml.v3"
)

type Properties struct {
	Author	string	`yaml:"Author"`
}

func (p *Properties) Register(properties *map[string]interface{}) {
	(*properties)["Author"] = p.Author
}


func BootStrap(env string) map[string]interface{} {

	properties := make(map[string]interface{})

	c := struct {
		Properties Properties `yaml:"Properties"`
	}{}

	data, err := config.ReadYAML(env)
	if err == nil {
		if err := yaml.Unmarshal(data, &c); err == nil {
			c.Properties.Register(&properties)
		}
	}
	return properties
}