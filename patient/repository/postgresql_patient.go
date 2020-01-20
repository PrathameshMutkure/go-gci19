package repository

import (
	"github.com/PrathameshMutkure/go-gci19/models"
	"github.com/PrathameshMutkure/go-gci19/patient"
	"github.com/jmoiron/sqlx"
)

type patientRepository struct {
	db *sqlx.DB
}

func NewPatientRepository(db *sqlx.DB) patient.Repository {
	return &patientRepository{db}
}

func (m *patientRepository) Fetch() ([]models.Patient, error) {
	var patients []models.Patient
	err := m.db.Select(&patients, "SELECT * FROM  patient")
	return patients, err
}

func (m *patientRepository) GetById(id int64) (models.Patient, error) {
	var _patient models.Patient
	err := m.db.Get(&_patient, "SELECT * FROM patient WHERE id = $1", id)
	return _patient, err
}

func (m *patientRepository) New(patient models.Patient) error {
	_, err := m.db.NamedExec("INSERT INTO patient(name, illness, birth_date, location_id) VALUES (:name, :illness, :birth_date, :location_id)", patient)
	return err
}
