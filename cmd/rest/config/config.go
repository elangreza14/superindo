package config

import "github.com/spf13/viper"

type Config struct {
	ENV               string   `mapstructure:"ENV"`
	HTTP_PORT         string   `mapstructure:"HTTP_PORT"`
	POSTGRES_USER     string   `mapstructure:"POSTGRES_USER"`
	POSTGRES_PASSWORD string   `mapstructure:"POSTGRES_PASSWORD"`
	POSTGRES_HOSTNAME string   `mapstructure:"POSTGRES_HOSTNAME"`
	POSTGRES_PORT     string   `mapstructure:"POSTGRES_PORT"`
	POSTGRES_DB       string   `mapstructure:"POSTGRES_DB"`
	POSTGRES_SSL      string   `mapstructure:"POSTGRES_SSL"`
	ALLOWED_ORIGINS   []string `mapstructure:"ALLOWED_ORIGINS"`
	REDIS_PORT        string   `mapstructure:"REDIS_PORT"`
}

func LoadConfig() (*Config, error) {
	viper.SetConfigFile(".env")
	viper.ReadInConfig()
	cfg := Config{}
	err := viper.Unmarshal(&cfg)
	if err != nil {
		return nil, err
	}

	return &cfg, nil
}
