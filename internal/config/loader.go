package config

import (
	"github.com/spf13/viper"
)

type ConfigLoader struct {
	primaryConfig *primaryConfig
}

func NewConfigLoader() *ConfigLoader {
	return &ConfigLoader{}
}

func (l *ConfigLoader) Load() (*ConfigLoader, error) {
	teamConfig, err := l.loadTeamConfig()
	if err != nil {
		return nil, err
	}
	templateConfig, err := l.loadTemplateConfig()
	if err != nil {
		return nil, err
	}
	config := &primaryConfig{
		Team:     *teamConfig,
		Template: *templateConfig,
	}
	return &ConfigLoader{
		primaryConfig: config,
	}, nil
}

func (l *ConfigLoader) loadTeamConfig() (*TeamConfig, error) {
	v := viper.New()
	v.SetConfigName("curly-team")
	v.SetConfigType("yaml")
	v.AddConfigPath(".")
	if err := v.ReadInConfig(); err != nil {
		return nil, err
	}
	var teamConfig TeamConfig
	if err := v.Unmarshal(&teamConfig); err != nil {
		return nil, err
	}
	return &teamConfig, nil
}

func (l *ConfigLoader) loadTemplateConfig() (*TemplateConfig, error) {
	v := viper.New()
	v.SetConfigName("template")
	v.SetConfigType("yml")
	v.AddConfigPath("templates/react")
	if err := v.ReadInConfig(); err != nil {
		return nil, err
	}
	var templateConfig TemplateConfig
	if err := v.Unmarshal(&templateConfig); err != nil {
		return nil, err
	}
	return &templateConfig, nil
}
