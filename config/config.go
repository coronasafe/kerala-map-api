package config

import (
	"github.com/vrischmann/envconfig"
)

var Config struct {
	Port int `envconfig:"default=3000"`
	DB   struct {
		URL      string `envconfig:"optional"`
		Host     string `envconfig:"optional"`
		Port     int    `envconfig:"optional"`
		User     string `envconfig:"optional"`
		Password string `envconfig:"optional"`
		Name     string `envconfig:"optional"`
	}
	JWT string
	KEY string
}

func init() {
	if err := envconfig.Init(&Config); err != nil {
		panic(err)
	}

	if Config.DB.URL == "" {
		if Config.DB.Host == "" ||
			Config.DB.Port == 0 ||
			Config.DB.User == "" ||
			Config.DB.Password == "" ||
			Config.DB.Name == "" {
			panic("Missing/Invalid Database ENV Vars")
		}
	}
}
