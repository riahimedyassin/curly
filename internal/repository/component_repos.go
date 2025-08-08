package repository

import (
	"github.com/riahimedyassin/curly/internal/dto"
	"github.com/riahimedyassin/curly/internal/interfaces"
)

type ComponentRepository struct {
	fs interfaces.FileSystem
}

func NewComponentRepository(fs interfaces.FileSystem) *ComponentRepository {
	return &ComponentRepository{
		fs: fs,
	}
}

func (r *ComponentRepository) CreateComponent(payload *dto.CreateComponent) error {
	if err := r.fs.Write(payload.Path, []byte(payload.Content)); err != nil {
		return err
	}
	if payload.Test != nil {
		if err := r.createTestFile(payload); err != nil {
			return err
		}
	}
	if payload.Style != nil {
		if err := r.createStyleFile(payload); err != nil {
			return err
		}
	}
	return nil
}

func (r *ComponentRepository) createTestFile(payload *dto.CreateComponent) error {
	return r.fs.Write(payload.Test.Path, []byte(payload.Test.Content))
}

func (r *ComponentRepository) createStyleFile(payload *dto.CreateComponent) error {
	return r.fs.Write(payload.Style.Path, []byte(payload.Style.Content))
}
