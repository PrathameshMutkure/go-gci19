package hospital

import "github.com/PrathameshMutkure/go-gci19/models"

type Repository interface {
	Fetch() ([]models.Hospital, error)
	GetById(id int64) (models.Hospital, error)
	New(models.Hospital) error
}
