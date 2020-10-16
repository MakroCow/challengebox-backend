package environment

import (
	"io/ioutil"
	"path/filepath"

	"github.com/go-yaml/yaml"
)

type Config struct {
	Host string `yaml:"host"`
	Port int `yaml:"port"`
	User string `yaml:"user"`
	Password string `yaml:"password"`
	DbName string `yaml:"dbname"`
}

var config = &Config{}

func Parse() *Config {
	if (*config != Config{}){
		return config
	} else {
		filepath, _ := filepath.Abs("./environment/config.yaml")
		configContent, err := ioutil.ReadFile(filepath)
		if err != nil {
			panic(err)
		}

		if err := yaml.Unmarshal(configContent, config); err != nil {
			panic (err)
		}
	}
	return config
}