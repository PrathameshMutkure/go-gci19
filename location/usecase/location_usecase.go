package usecase

import (
	"github.com/PrathameshMutkure/go-gci19/location"
	"github.com/PrathameshMutkure/go-gci19/models"
)

type locationUsecase struct {
	repository location.Repository
}

func NewLocationUsecase(repository location.Repository) location.Usecase {
	return &locationUsecase{repository}
}

func (m *locationUsecase) Fetch() (locations []models.Location, err error) {
	return m.repository.Fetch()
}

func (m *locationUsecase) GetById(id int64) (models.Location, error) {
	return m.repository.GetById(id)
}

func (m *locationUsecase) New(location models.Location) error {
	return m.repository.New(location)
}
