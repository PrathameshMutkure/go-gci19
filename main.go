package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
	"log"
)

const (
	DB_USER     = "twvjbojy"
	DB_PASSWORD = "H0kJZUiDgpgweAWNWdLRgI6v6-IEyxyr"
	DB_NAME     = "twvjbojy host=john.db.elephantsql.com"
	DB_HOST     = "john.db.elephantsql.com"
)

var router *gin.Engine
var db *sqlx.DB

type Patient struct {
	ID         int    `json:"id" db:"id"`
	Name       string `json:"name" db:"name"`
	Illness    string `json:"illness" db:"illness"`
	BirthDate  string `json:"birth_date" db:"birth_date"`
	LocationId int    `json:"location_id" db:"location_id"`
}

type Location struct {
	ID         int    `json:"id" db:"id"`
	Name       string `json:"name" db:"name"`
	HospitalId int    `json:"hospital_id" db:"hospital_id"`
}

type Hospital struct {
	ID               int    `json:"id" db:"id"`
	Name             string `json:"name" db:"name"`
	MaxPatientAmount int    `json:"max_patient_amount" db:"max_patient_amount"`
}

func main() {
	initDatabase()
	initAPI()
	err := router.Run(":8080")
	if err != nil {
		log.Fatalln("Could not run the app")
	}
	defer db.Close()
}

func initDatabase() {

	psqlInfo := fmt.Sprintf("user=%s password=%s dbname=%s host=%s", DB_USER, DB_PASSWORD, DB_NAME, DB_HOST)
	var err error

	db, err = sqlx.Open("postgres", psqlInfo)
	if err != nil {
		panic(err)
	}

	err = db.Ping()
	if err != nil {
		panic(err)
	}
	log.Println("Successfully connected to the database")
}

// This function will initialize all the routes
func initAPI() {
	router = gin.Default()
	router.GET("/hello", helloHandler)
	router.GET("/greet/:name", greetHandler)
	router.GET("/patient/all", GetAllPatients)
	router.GET("/location/all", GetAllLocations)
	router.GET("/hospital/all", GetAllHospitals)
}
