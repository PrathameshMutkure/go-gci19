package location

import "github.com/PrathameshMutkure/go-gci19/models"

type Usecase interface {
	Fetch() ([]models.Location, error)
	GetById(id int64) (models.Location, error)
	New(models.Location) error
}
