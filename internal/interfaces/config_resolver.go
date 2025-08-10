package interfaces

import "github.com/riahimedyassin/curly/internal/config"

type ConfigResolver interface {
	Load() (*ConfigResolver, error)
	Resolve() (*config.Config, error)
}
