package configuration

import (
	"fmt"
	"gopkg.in/yaml.v3"
	"log"
	"os"
)

type Database struct {
	Host     string `yaml:"host"`
	Port     int    `yaml:"port"`
	User     string `yaml:"user"`
	Password string `yaml:"password"`
	Name     string `yaml:"name"`
}
type App struct {
	Port  int  `yaml:"port"`
	Debug bool `yaml:"debug"`
}

type Configuration struct {
	App App      `yaml:"app"`
	DB  Database `yaml:"database"`
}

const defaultPath = "./dependency-injection/config.yaml"

func Load(filename string) (*Configuration, error) {
	var c Configuration
	file, err := os.ReadFile(filename)
	if err != nil {
		log.Fatal(err)
	}
	if err := yaml.Unmarshal(file, &c); err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Configuration: %#v\n", &c)
	return &c, nil
}

func New() (*Configuration, error) {
	return Load(defaultPath)
}
