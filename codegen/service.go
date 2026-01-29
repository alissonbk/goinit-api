package codegen

import (
	"fmt"
	"github.com/alissonbk/goinit-api/model"
)

func GenerateServiceContent(cfg model.Configuration) string {
	return fmt.Sprintf(`
		package service

		import (
			"%s/app/model/entity"
			"%s/app/repository"
		)

		type ExampleService struct {
			repository *repository.ExampleRepository
		}

		func NewExampleService(repository *repository.ExampleRepository) *ExampleService {
			return &ExampleService{repository: repository}
		}

		func (s *ExampleService) GetAll() []*entity.Example {
			return s.repository.FindAllExample()
		}
		
	`, cfg.ModulePath, cfg.ModulePath)
}
