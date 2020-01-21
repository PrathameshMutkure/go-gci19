package repository

import (
	"github.com/PrathameshMutkure/go-gci19/hospital"
	"github.com/PrathameshMutkure/go-gci19/models"
	"github.com/jmoiron/sqlx"
)

type hospitalRepository struct {
	db *sqlx.DB
}

func NewHospitalRepository(db *sqlx.DB) hospital.Repository {
	return &hospitalRepository{db}
}

func (m *hospitalRepository) Fetch() ([]models.Hospital, error) {
	var hospitals []models.Hospital
	err := m.db.Select(&hospitals, "SELECT * FROM  hospital")
	return hospitals, err
}

func (m *hospitalRepository) GetById(id int64) (models.Hospital, error) {
	var _hospital models.Hospital
	err := m.db.Get(&_hospital, "SELECT * FROM hospital WHERE id = $1", id)
	return _hospital, err
}

func (m *hospitalRepository) New(hospital models.Hospital) error {
	_, err := m.db.NamedExec("INSERT INTO hospital(name, max_patient_amount) VALUES (:name, :max_patient_amount)", hospital)
	return err
}
