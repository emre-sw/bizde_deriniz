package configs

import (
	"fmt"
	"log"

	"github.com/spf13/viper"
)

type Config struct {
	KafkaTopics           string `mapstructure:"KAFKA_TOPICS"`
	KafkaBootstrapServers string `mapstructure:"KAFKA_BOOTSTRAP_SERVERS"`
	KafkaGroupID          string `mapstructure:"KAFKA_GROUP_ID"`
	KafkaZookeeperConnect string `mapstructure:"KAFKA_ZOOKEEPER_CONNECT"`

	EmailHost     string `mapstructure:"EMAIL_HOST"`
	EmailPort     string `mapstructure:"EMAIL_PORT"`
	EmailUsername string `mapstructure:"EMAIL_USERNAME"`
	EmailPassword string `mapstructure:"EMAIL_PASSWORD"`
	EmailFrom     string `mapstructure:"EMAIL_FROM"`
}

var cachedConfig *Config

func LoadConfig(envPath ...string) (*Config, error) {
	viper.AutomaticEnv()

	if len(envPath) > 0 && envPath[0] != "" {
		viper.SetConfigFile(envPath[0])
		if err := viper.ReadInConfig(); err != nil {
			fmt.Printf("Error reading .env file: %s\n", err)
		} else {
			fmt.Println("Using .env file:", viper.ConfigFileUsed())
		}
	}

	config := &Config{}
	if err := viper.Unmarshal(config); err != nil {
		return nil, err
	}

	cachedConfig = config

	return config, nil
}

func GetString(key string) string {
	return viper.GetString(key)
}

func GetConfig() *Config {
	if cachedConfig == nil {
		log.Fatal("Config not loaded")
	}
	return cachedConfig
}
