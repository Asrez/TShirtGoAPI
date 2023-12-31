package config

import (
	"errors"
	"log"
	"os"
	"time"

	viper "github.com/spf13/viper"
)

type Config struct {
	Server ServerConfig
	Postgres PostgresConfig
	Redis RedisConfig
	Logger LoggerConfig
}

type ServerConfig struct {
	Port string
}

type PostgresConfig struct {
	Host string
	Port string
	User string
	Password string
	DbName string
	SSLMode string
}

type LoggerConfig struct {
	FilePath string
	Encoding string
	Level string
	Logger string
}

type RedisConfig struct {
	Host               string
	Port               string
	Password           string
	Db                 string
	DialTimeout        time.Duration
	ReadTimeout        time.Duration
	WriteTimeout       time.Duration
	IdleCheckFrequency time.Duration
	PoolSize           int
	PoolTimeout        time.Duration
}

func GetConfig() *Config {
	configPath :=  getConfigPath(os.Getenv("APP_ENV"))
	v , err := LoadConfig(configPath , "yml")
	if err != nil {
		log.Fatal(err)
	}
	config , err := ParseConfig(v)

	if err != nil {
		log.Fatal("Error in Parse Config" , err)
	}
	return config

}

func ParseConfig(v *viper.Viper) (*Config, error) {
	var config Config

	err := v.Unmarshal(&config)
	if err != nil {
		log.Printf("Unable to parse config: %v", err)
	}
	return &config, nil
}

func LoadConfig(filename string, fileType string) (*viper.Viper, error) {
	v := viper.New()
	v.SetConfigType(fileType)
	v.SetConfigName(filename)
	v.AddConfigPath(".")
	v.AutomaticEnv()

	err := v.ReadInConfig()
	if err != nil {
		log.Printf("Unable to read config: %v", err)
		if _, ok := err.(viper.ConfigFileNotFoundError); ok {
			return nil, errors.New("config file not found")
		}
		return nil, err
	}
	return v, nil
}

func getConfigPath(env string) string {
	if env == "production" {
		return "../config/config-production"
	}else {
		return "../config/config-development"
	}
}