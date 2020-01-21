package repository

import (
	"github.com/PrathameshMutkure/go-gci19/location"
	"github.com/PrathameshMutkure/go-gci19/models"
	"github.com/jmoiron/sqlx"
)

type locationRepository struct {
	db *sqlx.DB
}

func NewLocationRepository(db *sqlx.DB) location.Repository {
	return &locationRepository{db}
}

func (m *locationRepository) Fetch() ([]models.Location, error) {
	var locations []models.Location
	err := m.db.Select(&locations, "SELECT * FROM  location")
	return locations, err
}

func (m *locationRepository) GetById(id int64) (models.Location, error) {
	var _location models.Location
	err := m.db.Get(&_location, "SELECT * FROM location WHERE id = $1", id)
	return _location, err
}

func (m *locationRepository) New(location models.Location) error {
	_, err := m.db.NamedExec("INSERT INTO location(name, hospital_id) VALUES (:name, :max_patient_amount)", location)
	return err
}
