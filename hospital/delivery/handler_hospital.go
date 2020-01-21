package delivery

import (
	"github.com/PrathameshMutkure/go-gci19/hospital"
	"github.com/PrathameshMutkure/go-gci19/models"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

type hospitalHandler struct {
	hospitalUsecase hospital.Usecase
}

func NewHospitalHandler(group *gin.RouterGroup, usecase hospital.Usecase) {
	handler := hospitalHandler{usecase}

	group.GET("/", func(context *gin.Context) {
		hospitalsRetrieved, err := handler.hospitalUsecase.Fetch()
		if err != nil {
			context.JSON(http.StatusInternalServerError, models.ResponseError{Message: err.Error()})
			return
		}

		context.JSON(http.StatusOK, hospitalsRetrieved)
	})

	group.GET("/:id", func(context *gin.Context) {
		idInt, err := strconv.Atoi(context.Param("id"))
		if err != nil {
			context.JSON(http.StatusInternalServerError, models.ResponseError{Message: err.Error()})
			return
		}
		id := int64(idInt)

		_hospital, err := handler.hospitalUsecase.GetById(id)
		if err != nil {
			context.JSON(http.StatusNotFound, models.ResponseError{Message: err.Error()})
			return
		}

		context.JSON(http.StatusOK, _hospital)
	})

	group.POST("/", func(context *gin.Context) {
		var _hospital models.Hospital
		err := context.BindJSON(&_hospital)
		if err != nil {
			context.JSON(http.StatusInternalServerError, models.ResponseError{Message: err.Error()})
			return
		}

		err = handler.hospitalUsecase.New(_hospital)
		if err != nil {
			context.JSON(http.StatusInternalServerError, models.ResponseError{Message: err.Error()})
			return
		}

		context.String(http.StatusOK, "Successfully inserted hospital object into database")
	})
}
