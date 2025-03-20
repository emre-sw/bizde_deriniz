package configs

import (
	"log"
	"time"

	"github.com/spf13/viper"
)

type Config struct {
	DBHost     string `mapstructure:"POSTGRES_HOST"`
	DBPort     string `mapstructure:"POSTGRES_PORT"`
	DBUser     string `mapstructure:"POSTGRES_USER"`
	DBPassword string `mapstructure:"POSTGRES_PASSWORD"`
	DBName     string `mapstructure:"POSTGRES_DB"`
	DBSSLMode  string `mapstructure:"POSTGRES_SSL_MODE"`

	KafkaBootstrapServers string `mapstructure:"KAFKA_BOOTSTRAP_SERVERS"`
	KafkaTopic            string `mapstructure:"KAFKA_TOPIC"`
	KafkaGroupID          string `mapstructure:"KAFKA_AUTH_GROUP_ID"`

	RedisHost     string `mapstructure:"REDIS_HOST"`
	RedisPort     string `mapstructure:"REDIS_PORT"`
	RedisPassword string `mapstructure:"REDIS_PASSWORD"`
	RedisDB       int    `mapstructure:"REDIS_DB"`

	JwtAccessSecret      string        `mapstructure:"JWT_ACCESS_SECRET"`
	JwtAccessExpiration  time.Duration `mapstructure:"JWT_ACCESS_EXPIRATION_TIME"`
	JwtRefreshSecret     string        `mapstructure:"JWT_REFRESH_SECRET"`
	JwtRefreshExpiration time.Duration `mapstructure:"JWT_REFRESH_EXPIRATION_TIME"`

	MaxLoginAttempts  int64         `mapstructure:"MAX_LOGIN_ATTEMPTS"`
	LoginAttemptsTime time.Duration `mapstructure:"LOGIN_ATTEMPTS_WINDOW"`
}

var cachedConfig *Config

func LoadConfig() (*Config, error) {
	if cachedConfig != nil {
		return cachedConfig, nil
	}

	viper.AutomaticEnv()
	viper.SetConfigFile(".env")

	if err := viper.ReadInConfig(); err != nil {
		log.Println("No .env file found, using only environment variables")
	}

	config := &Config{}
	if err := viper.Unmarshal(config); err != nil {
		log.Fatalf("Unable to decode into config struct: %v", err)
	}

	cachedConfig = config
	return config, nil
}

func GetString(key string) string {
	return viper.GetString(key)
}

func GetConfig() *Config {
	if cachedConfig == nil {
		config, err := LoadConfig()
		if err != nil {
			log.Fatal("Error loading config:", err)
		}
		cachedConfig = config
	}
	return cachedConfig
}
