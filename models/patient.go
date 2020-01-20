package models

type Patient struct {
	Id         int    `json:"id" db:"id"`
	Name       string `json:"name" db:"name"`
	Illness    string `json:"illness" db:"illness"`
	BirthDate  string `json:"birthDate" db:"birth_date"`
	LocationId int    `json:"locationId" db:"location_id"`
}
