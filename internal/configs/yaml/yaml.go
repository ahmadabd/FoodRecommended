package yaml

import (
	"io/ioutil"
	"log"

	"github.com/ahmadabd/FoodRecommended.git/internal/configs"
	"gopkg.in/yaml.v2"
)

type (
	Config struct {
		Server   Server   `yaml:"server"`
		Database Database `yaml:"database"`
	}

	Server struct {
		Host string `yaml:"host"`
		Port string `yaml:"port"`
	}

	Database struct {
		Host     string `yaml:"host"`
		Port     string `yaml:"port"`
		User     string `yaml:"user"`
		Password string `yaml:"password"`
		Name     string `yaml:"name"`
	}
)

func GetConfig() configs.ConfigImp {
	c := &Config{}
	yamlFile, err := ioutil.ReadFile("config.yml")
	if err != nil {
		log.Println("Error in reading config file")
	}

	err = yaml.Unmarshal(yamlFile, c)
	if err != nil {
		log.Println("Error in unmarshalling config file")
	}

	return c
}
