package config

import (
	"fmt"
	"strings"

	"github.com/spf13/viper"
)

type Config struct {
	GRPC      GRPCConf
	Logger    LoggerConf
	Redis     RedisConf
	RateLimit RateLimitConf
}

type LoggerConf struct {
	Level string
}

type RedisConf struct {
	DSN string
}

type GRPCConf struct {
	Address string
}

type RateLimitConf struct {
	LoginPerMinute    int
	PasswordPerMinute int
	IPPerMinute       int
}

func NewConfig(configFile string) (*Config, error) {
	viper.SetConfigFile(configFile)
	viper.AutomaticEnv()
	viper.SetEnvKeyReplacer(strings.NewReplacer(".", "_"))

	if err := viper.ReadInConfig(); err != nil {
		return nil, fmt.Errorf("%v: %w", ErrFailedReadConfigFile, err)
	}

	config := new(Config)

	err := viper.Unmarshal(config)
	if err != nil {
		return nil, fmt.Errorf("unable to decode into struct, %w", err)
	}

	return config, nil
}
