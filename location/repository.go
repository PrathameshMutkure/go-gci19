package location

import "github.com/PrathameshMutkure/go-gci19/models"

type Repository interface {
	Fetch() ([]models.Location, error)
	GetById(id int64) (models.Location, error)
	New(models.Location) error
}
