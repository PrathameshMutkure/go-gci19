package delivery

import (
	"github.com/PrathameshMutkure/go-gci19/models"
	"github.com/PrathameshMutkure/go-gci19/patient"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

type patientHandler struct {
	patientUsecase patient.Usecase
}

func NewPatientHandler(group *gin.RouterGroup, usecase patient.Usecase) {
	handler := patientHandler{usecase}

	group.GET("/", func(context *gin.Context) {
		patientsRetrieved, err := handler.patientUsecase.Fetch()
		if err != nil {
			context.JSON(http.StatusInternalServerError, models.ResponseError{Message: err.Error()})
			return
		}

		context.JSON(http.StatusOK, patientsRetrieved)
	})

	group.GET("/:id", func(context *gin.Context) {
		idInt, err := strconv.Atoi(context.Param("id"))
		if err != nil {
			context.JSON(http.StatusInternalServerError, models.ResponseError{Message: err.Error()})
			return
		}
		id := int64(idInt)

		_patient, err := handler.patientUsecase.GetById(id)
		if err != nil {
			context.JSON(http.StatusNotFound, models.ResponseError{Message: err.Error()})
			return
		}

		context.JSON(http.StatusOK, _patient)
	})

	group.POST("/", func(context *gin.Context) {
		var _patient models.Patient
		err := context.BindJSON(&_patient)
		if err != nil {
			context.JSON(http.StatusInternalServerError, models.ResponseError{Message: err.Error()})
			return
		}

		err = handler.patientUsecase.New(_patient)
		if err != nil {
			context.JSON(http.StatusInternalServerError, models.ResponseError{Message: err.Error()})
			return
		}

		context.String(http.StatusOK, "Successfully inserted Patient object into database")
	})
}
