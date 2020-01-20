package models

type Location struct {
	ID         int    `json:"id" db:"id"`
	Name       string `json:"name" db:"name"`
	HospitalId int    `json:"hospitalId" db:"hospital_id"`
}
