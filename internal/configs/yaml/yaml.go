package yaml

import (
	"errors"
	"io/ioutil"
	"log"
	"path/filepath"

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

var ErrUnknownFileExtension = errors.New("unknown file extension")

func GetConfig(path string) (configs.ConfigImp, error) {

	switch filepath.Ext(path) {
	case ".yaml", ".yml":
		return ParsYaml(path), nil
	default:
		return nil, ErrUnknownFileExtension
	}
}

func ParsYaml(path string) configs.ConfigImp {
	c := &Config{}
	yamlFile, err := ioutil.ReadFile(path)
	if err != nil {
		log.Println("Error in reading config file")
	}

	err = yaml.Unmarshal(yamlFile, c)
	if err != nil {
		log.Println("Error in unmarshalling config file")
	}

	return c
}
