package configs

import (
	"io/ioutil"
	"log"

	"gopkg.in/yaml.v2"
)

type Config struct {
	Server struct {
		Host string `yaml:"host"`
		Port string `yaml:"port"`
	} `yaml:"server"`
	Database struct {
		Host     string `yaml:"host"`
		Port     string `yaml:"port"`
		User     string `yaml:"user"`
		Password string `yaml:"password"`
		Name     string `yaml:"name"`
	} `yaml:"database"`
}

var Conf *Config

func GetConfig() {
	c := &Config{}
	yamlFile, err := ioutil.ReadFile("config.yml")
	if err != nil {
		log.Println("Error in reading config file")
	}

	err = yaml.Unmarshal(yamlFile, c)
	if err != nil {
		log.Println("Error in unmarshalling config file")
	}

	Conf = c
}

func (c *Config) GetServerHost() string {
	return c.Server.Host
}

func (c *Config) GetDatabaseHost() string {
	return c.Database.Host
}

func (c *Config) GetServerPort() string {
	return c.Server.Port
}

func (c *Config) GetDatabasePort() string {
	return c.Database.Port
}

func (c *Config) GetDatabaseName() string {
	return c.Database.Name
}

func (c *Config) GetDatabaseUser() string {
	return c.Database.User
}

func (c *Config) GetDatabasePassword() string {
	return c.Database.Password
}
