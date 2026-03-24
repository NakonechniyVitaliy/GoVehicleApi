package config

import (
	"log"
	"os"
	"time"

	"github.com/ilyakaznacheev/cleanenv"
)

type Config struct {
	Env          string `yaml:"env" env:"ENV" env-default:"local"`
	StoragePath  string `yaml:"storage_path" env-required:"true"`
	MongoURI     string `yaml:"mongo_uri" env-required:"true"`
	DataBase     string `yaml:"database" env-default:"sqlite"`
	SecretJwtKey string `yaml:"secret_jwt_key" env-required:"true"`
	AutoriaKey   string `yaml:"autoria_key" env-required:"true"`
	HTTPServer   `yaml:"http_server"`
	Redis        `yaml:"redis"`
}

type HTTPServer struct {
	Address     string        `yaml:"address" env-required:"true"`
	Timeout     time.Duration `yaml:"timeout" env-default:"4s"`
	IdleTimeout time.Duration `yaml:"idle_timeout" env-default:"60s"`
}

type Redis struct {
	Address   string        `yaml:"address" env-required:"true"`
	Password  string        `yaml:"password" env-required:"true"`
	DB        int           `yaml:"db" env-default:"0"`
	CacheTime time.Duration `yaml:"cache_time" env-default:"5m"`
}

func MustLoad() *Config {
	configPath := os.Getenv("CONFIG_PATH")
	if configPath == "" {
		log.Fatal("CONFIG_PATH is not set")
	}

	// check if file exists
	if _, err := os.Stat(configPath); os.IsNotExist(err) {
		log.Fatalf("CONFIG_PATH does not exist: %s", configPath)
	}

	var cfg Config

	if err := cleanenv.ReadConfig(configPath, &cfg); err != nil {
		log.Fatalf("cannot read config: %s", err)
	}

	return &cfg

}
