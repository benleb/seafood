package utils

import "github.com/spf13/viper"

var Conf *Config

type Config struct {
	EtherScanApiKey string `mapstructure:"ETHERSCAN_API_KEY"`

	PostgresHost     string `mapstructure:"POSTGRES_HOST"`
	PostgresUser     string `mapstructure:"POSTGRES_USER"`
	PostgresPassword string `mapstructure:"POSTGRES_PASSWORD"`
	PostgresDatabase string `mapstructure:"POSTGRES_DATABASE"`

	BarkURL string `mapstructure:"BARK_URL"`
	BarkKey string `mapstructure:"BARK_Key"`

	TelegramBotToken  string `mapstructure:"TELEGRAM_BOT_TOKEN"`
	TelegramBotAPIURL string `mapstructure:"TELEGRAM_BOT_API_URL"`

	OpenSeaMyAddr   string `mapstructure:"OPENSEA_MY_ADDR"`
	OpenSeaUsername string `mapstructure:"OPENSEA_USERNAME"`
}

// loadConfig reads the config from file or environment.
func LoadConfig(path string) (config Config, err error) {
	viper.AddConfigPath(path)
	viper.SetConfigName("app")
	viper.SetConfigType("env")

	viper.AutomaticEnv()

	err = viper.ReadInConfig()
	if err != nil {
		return
	}
	err = viper.Unmarshal(&config)
	return
}
