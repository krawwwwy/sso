package config

import (
	"log"
	"os"
	"time"

	"github.com/ilyakaznacheev/cleanenv"
)

type Config struct {
	Env         string        `yaml:"env" env-default:"local"`
	StoragePath string        `yaml:"storage_path" env-required:"true"`
	TokenTTL    time.Duration `yaml:"token_ttl" env-default:"1h"`
	GRPC        GRPCConfig    `yaml:"grpc"`
}

type GRPCConfig struct {
	Port    int           `yaml:"port"`
	Timeout time.Duration `yaml:"timeout" env-default:"2m"`
}

func MustLoad() *Config {
	//Создаем конфигурационный путь
	congifPath := os.Getenv("CONFIG_PATH")
	if congifPath == "" {
		log.Fatal("CONFIG_PATH is not set")
	}

	//Проверяем существование конфигурационного файла
	if _, err := os.Stat(congifPath); os.IsNotExist(err) {
		log.Fatalf("config file does not exist: %s", congifPath)
	}

	var cfg Config

	if err := cleanenv.ReadConfig(congifPath, &cfg); err != nil {
		log.Fatalf("error reading config: %s", err)
	}

	return &cfg

}
