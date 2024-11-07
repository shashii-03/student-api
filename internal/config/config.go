package config

import (
	"flag"
	"log"
	"os"

	"github.com/ilyakaznacheev/cleanenv"
)

type HttpServers struct{
	Addr string
}
type Config struct{
	Env string		`yaml:"env" env:"ENV" env-required:"true" env-default:"production"`
	StoragePath	string `yaml:"storage_path" env-required:"true"`
	HttpServers		`yaml:"http_server"`
}

func MustLoad() *Config{


	configPath := os.Getenv("CONFIG_PATH")

	if configPath ==""{
		flags:= flag.String("config", "","path to configuration file")
		flag.Parse()

		configPath := *flags
		 
		if configPath==""{
			log.Fatal("config path not set")
		}

	}
	if _, err := os.Stat(configPath); os.IsNotExist(err){
		log.Fatalf("config file does not exist at %s", configPath)
	}

	var cfg Config

	err :=cleanenv.ReadConfig(configPath, &cfg)

	if err !=nil{
		log.Fatal("cannot read config file")
	}

	return &cfg

}
