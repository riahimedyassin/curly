package interfaces

import "github.com/riahimedyassin/curly/internal/config"

type ConfigResolver interface {
	// Called at a root level to init the config
	Load() (*ConfigResolver, error)
	// Called to retrieve the config at child level.
	Resolve() (*config.Config, error)
}
