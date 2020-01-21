package hospital

import "github.com/PrathameshMutkure/go-gci19/models"

type Usecase interface {
	Fetch() ([]models.Hospital, error)
	GetById(id int64) (models.Hospital, error)
	New(hospital models.Hospital) error
}
