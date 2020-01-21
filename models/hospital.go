package models

type Hospital struct {
	ID               int    `json:"id" db:"id"`
	Name             string `json:"name" db:"name"`
	MaxPatientAmount int    `json:"maxPatientAmount" db:"max_patient_amount"`
}
