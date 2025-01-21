package config

import (
	"log"

	"github.com/spf13/viper"
)

type Config struct {
	Logger       LoggerConf   `mapstructure:"logging"`
	Db           DBConfig     `mapstructure:"db"`
	StorageType  string       `mapstructure:"storage_type"`
	ServerConfig ServerConfig `mapstructure:"server_config"`
}

type LoggerConf struct {
	Level string `mapstructure:"log_level"`
}

type DBConfig struct {
	Host     string `mapstructure:"db_host"`
	Port     string `mapstructure:"db_port"`
	Login    string `mapstructure:"db_login"`
	Pass     string `mapstructure:"db_password"`
	Database string `mapstructure:"db_database"`
	Schema   string `mapstructure:"db_schema"`
}

type ServerConfig struct {
	Host string `mapstructure:"server_host"`
	Port string `mapstructure:"server_port"`
}

func NewConfig() Config {
	return Config{}
}

func (c *Config) Load(path string) error {
	viper.SetConfigFile(path)
	if err := viper.ReadInConfig(); err != nil {
		log.Fatal(err)
		return err
	}

	log.Printf("Viper keys:")
	for _, key := range viper.AllKeys() {
		log.Printf("  %s: %v", key, viper.Get(key))
	}

	if err := viper.Unmarshal(&c); err != nil {
		log.Println(err)
		return err
	}
	log.Printf("config: %#v", c)
	return nil
}
