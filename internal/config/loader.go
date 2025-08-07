package config

import (
	"os"

	"gopkg.in/yaml.v3"
)

type ConfigLoader struct {
	Config *Config
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
	config := &Config{
		Team:     *teamConfig,
		Template: *templateConfig,
	}
	return &ConfigLoader{
		Config: config,
	}, nil
}

// todo : Implement config resolver
func (l *ConfigLoader) Resolve() {

}

func (l *ConfigLoader) loadTeamConfig() (*TeamConfig, error) {
	teamConfig, err := os.ReadFile("./curly-team.yml")
	if err != nil {
		return nil, err
	}
	var parsedTeamConfig TeamConfig
	if err := yaml.Unmarshal(teamConfig, &parsedTeamConfig); err != nil {
		return nil, err
	}
	return &parsedTeamConfig, nil
}

func (l *ConfigLoader) loadTemplateConfig() (*TemplateConfig, error) {
	teamplteConfig, err := os.ReadFile("templates/react/template.yml")
	if err != nil {
		return nil, err
	}
	var parsedTeamplteConfig TemplateConfig
	if err := yaml.Unmarshal(teamplteConfig, &parsedTeamplteConfig); err != nil {
		return nil, err
	}
	return &parsedTeamplteConfig, nil
}
