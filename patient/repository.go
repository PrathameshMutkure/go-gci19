package patient

import "github.com/PrathameshMutkure/go-gci19/models"

type Repository interface {
	Fetch() ([]models.Patient, error)
	GetById(id int64) (models.Patient, error)
	New(patient models.Patient) error
}
