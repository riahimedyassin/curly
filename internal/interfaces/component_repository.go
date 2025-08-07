package interfaces

import "github.com/riahimedyassin/curly/internal/dto"

type ComponentRepository interface {
	CreateComponent(payload *dto.CreateComponent) error
}
