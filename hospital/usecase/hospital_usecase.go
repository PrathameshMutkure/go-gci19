package usecase

import (
	"github.com/PrathameshMutkure/go-gci19/hospital"
	"github.com/PrathameshMutkure/go-gci19/models"
)

type hospitalUsecase struct {
	repository hospital.Repository
}

func NewHospitalUsecase(repository hospital.Repository) hospital.Usecase {
	return &hospitalUsecase{repository}
}

func (m *hospitalUsecase) Fetch() (hospitals []models.Hospital, err error) {
	return m.repository.Fetch()
}

func (m *hospitalUsecase) GetById(id int64) (models.Hospital, error) {
	return m.repository.GetById(id)
}

func (m *hospitalUsecase) New(hospital models.Hospital) error {
	return m.repository.New(hospital)
}
