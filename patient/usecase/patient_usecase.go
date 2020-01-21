package usecase

import (
	"github.com/PrathameshMutkure/go-gci19/models"
	"github.com/PrathameshMutkure/go-gci19/patient"
)

type patientUsecase struct {
	repository patient.Repository
}

func NewPatientUsecase(repository patient.Repository) patient.Usecase {
	return &patientUsecase{repository}
}

func (m *patientUsecase) Fetch() (patients []models.Patient, err error) {
	return m.repository.Fetch()
}

func (m *patientUsecase) GetById(id int64) (models.Patient, error) {
	return m.repository.GetById(id)
}

func (m *patientUsecase) New(patient models.Patient) error {
	return m.repository.New(patient)
}
