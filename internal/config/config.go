package config

import "github.com/spf13/viper"

type Config struct {
	TargetLang string `mapstructure:"target_lang"`
}

// LoadConfig ищет config.yaml в корне и загружает его.
func LoadConfig() (*Config, error) {
	viper.SetConfigName("config")
	viper.SetConfigType("yaml")
	viper.AddConfigPath(".") // Ищем в корне запуска

	// Значение по умолчанию, если файла вообще нет
	viper.SetDefault("target_lang", "en")

	if err := viper.ReadInConfig(); err != nil {
		// Если файл не найден, это не критично — сработает Default или Флаг
		if _, ok := err.(viper.ConfigFileNotFoundError); !ok {
			return nil, err
		}
	}

	var cfg Config
	if err := viper.Unmarshal(&cfg); err != nil {
		return nil, err
	}

	return &cfg, nil
}
