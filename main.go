package main

import (
	HospitalDelivery "github.com/PrathameshMutkure/go-gci19/hospital/delivery"
	HospitalRepo "github.com/PrathameshMutkure/go-gci19/hospital/repository"
	HospitalUsecase "github.com/PrathameshMutkure/go-gci19/hospital/usecase"
	LocationDelivery "github.com/PrathameshMutkure/go-gci19/location/delivery"
	LocationRepo "github.com/PrathameshMutkure/go-gci19/location/repository"
	LocationUsecase "github.com/PrathameshMutkure/go-gci19/location/usecase"
	PatientDelivery "github.com/PrathameshMutkure/go-gci19/patient/delivery"
	PatientRepo "github.com/PrathameshMutkure/go-gci19/patient/repository"
	PatientUsecase "github.com/PrathameshMutkure/go-gci19/patient/usecase"
	"github.com/gin-gonic/gin"
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

func main() {

	db := sqlx.MustConnect("postgres", "postgres://twvjbojy:H0kJZUiDgpgweAWNWdLRgI6v6-IEyxyr@john.db.elephantsql.com:5432/twvjbojy")
	router := gin.Default()

	// Repository Instances
	var hospitalRepo = HospitalRepo.NewHospitalRepository(db)
	var locationRepo = LocationRepo.NewLocationRepository(db)
	var patientRepo = PatientRepo.NewPatientRepository(db)

	// Usecase Instances
	var hospitalUsecase = HospitalUsecase.NewHospitalUsecase(hospitalRepo)
	var locationUsecase = LocationUsecase.NewLocationUsecase(locationRepo)
	var patientUsecase = PatientUsecase.NewPatientUsecase(patientRepo)

	// Endpoint Groups
	var hospitalGroup = router.Group("/hospitals")
	var locationGroup = router.Group("/locations")
	var patientGroup = router.Group("/patients")

	// API Routers
	HospitalDelivery.NewHospitalHandler(hospitalGroup, hospitalUsecase)
	LocationDelivery.NewLocationHandler(locationGroup, locationUsecase)
	PatientDelivery.NewPatientHandler(patientGroup, patientUsecase)

	router.Run(":8080")
	defer db.Close()
}
