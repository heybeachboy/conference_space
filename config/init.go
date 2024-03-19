package config

import (
	"gopkg.in/yaml.v3"
	"log"
	"os"
)

var Config *config

func init() {
	//Config = new(config)
}

func InitConfig(path string) error {
	by, err := os.ReadFile(path)
	if err != nil {
		return err
	}
	Config = new(config)
	if err = yaml.Unmarshal(by, Config); err != nil {
		log.Printf("read config content :\n%s\n", by)
		return err
	}
	log.Printf("read config success : %v", Config)
	return nil
}
