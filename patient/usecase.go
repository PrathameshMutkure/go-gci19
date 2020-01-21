package patient

import "github.com/PrathameshMutkure/go-gci19/models"

type Usecase interface {
	Fetch() ([]models.Patient, error)
	GetById(id int64) (models.Patient, error)
	New(patient models.Patient) error
}
