package config

import (
	"github.com/ilyakaznacheev/cleanenv"
	"log"
	"sync"
)

type Config struct {
	App struct {
		Repo string `yaml:"repo" env:"APP_REPO"`
		Rest RestConfig
		GRPC GRPCConfig
	} `yaml:"app"`
	DB     DBConfig
	Broker BrokerConfig
}

type RestConfig struct {
	Port string `yaml:"port" env:"APP_PORT"`
}

type GRPCConfig struct {
	Port string `yaml:"port" env:"APP_PORT"`
}

type DBConfig struct {
	Host     string `yaml:"host" env:"DB_HOST"`
	Port     string `yaml:"port" env:"DB_PORT"`
	Username string `yaml:"username" env:"DB_USERNAME"`
	Password string `yaml:"password" env:"DB_PASSWORD"`
	Database string `yaml:"database" env:"DB_DATABASE"`
}

type BrokerConfig struct {
	Host     string `yaml:"host" env:"BROKER_HOST"`
	Port     string `yaml:"port" env:"BROKER_PORT"`
	Username string `yaml:"username" env:"BROKER_USERNAME"`
	Password string `yaml:"password" env:"BROKER_PASSWORD"`
}

var instance *Config
var once sync.Once

func GetConfig() *Config {
	once.Do(func() {
		instance = &Config{}
		if err := cleanenv.ReadConfig("./config/config.yaml", instance); err != nil {
			log.Fatal(err)
		}
	})
	return instance
}
