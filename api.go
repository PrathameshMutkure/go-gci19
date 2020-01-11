package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

/*
 * Returns a HelloWorld json message
 */
func helloHandler(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"message": "Hello World!"})
}

/*
 * Returns a greeting json message
 */
func greetHandler(c *gin.Context) {
	name := c.Param("name")
	c.JSON(http.StatusOK, gin.H{"greet-message": "Hello " + name + "!"})
}

/*
 * Returns a json array of all the products
 */
func GetAllPatients(c *gin.Context) {

	var patients []Patient
	err := db.Select(&patients, "SELECT * FROM patient")

	if err != nil {
		panic(err)
	}

	c.JSON(http.StatusOK, patients)
}

/*
 * Returns a json array of all the locations
 */
func GetAllLocations(c *gin.Context) {

	var locations []Location
	err := db.Select(&locations, "SELECT * FROM location")

	if err != nil {
		panic(err)
	}

	c.JSON(http.StatusOK, locations)
}

/*
 * Returns json array of all the hospitals
 */
func GetAllHospitals(c *gin.Context) {

	var hospitals []Hospital
	err := db.Select(&hospitals, "SELECT * FROM hospital")

	if err != nil {
		panic(err)
	}

	c.JSON(http.StatusOK, hospitals)
}
