package services

import "github.com/riahimedyassin/curly/internal/interfaces"

type ComponentService struct {
	repos interfaces.ComponentRepository
}

func NewComponentService(repos interfaces.ComponentRepository) *ComponentService {
	return &ComponentService{
		repos: repos,
	}
}
