package services

import (
	"github.com/riahimedyassin/curly/internal/dto"
	"github.com/riahimedyassin/curly/internal/interfaces"
)

// todo : Get the user args.
// Get the config.
// Validate the user args structure. -> Moved to validator package concern
// Validate the user args against the global config. -> Moved to validator package concern
// Get templates and parse them.
// -> Pass everything to the repos.
type ComponentService struct {
	repos          interfaces.ComponentRepository
	cmdParser      interfaces.Parser[dto.ComponentArgs]
	configResolver interfaces.ConfigResolver
}

func NewComponentService(repos interfaces.ComponentRepository, cmdParser interfaces.Parser[dto.ComponentArgs], configResolver interfaces.ConfigResolver) *ComponentService {
	return &ComponentService{
		repos:          repos,
		cmdParser:      cmdParser,
		configResolver: configResolver,
	}
}

func (s *ComponentService) Execute() error {
	_, err := s.cmdParser.Parse()
	if err != nil {
		return err
	}
	_, err = s.configResolver.Resolve()
	if err != nil {
		return err
	}
	return nil
}

func (s *ComponentService) getComponentTemplate() {

}

func (s *ComponentService) getComponentPath() {
}
