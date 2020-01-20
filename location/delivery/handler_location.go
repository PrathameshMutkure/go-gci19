package delivery

import (
	"github.com/PrathameshMutkure/go-gci19/location"
	"github.com/PrathameshMutkure/go-gci19/models"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

type locationHandler struct {
	locationUsecase location.Usecase
}

func NewLocationHandler(group *gin.RouterGroup, usecase location.Usecase) {
	handler := locationHandler{usecase}

	group.GET("/", func(context *gin.Context) {
		locationRetrieved, err := handler.locationUsecase.Fetch()
		if err != nil {
			context.JSON(http.StatusInternalServerError, models.ResponseError{Message: err.Error()})
			return
		}

		context.JSON(http.StatusOK, locationRetrieved)
	})

	group.GET("/:id", func(context *gin.Context) {
		idInt, err := strconv.Atoi(context.Param("id"))
		if err != nil {
			context.JSON(http.StatusInternalServerError, models.ResponseError{Message: err.Error()})
			return
		}
		id := int64(idInt)

		_location, err := handler.locationUsecase.GetById(id)
		if err != nil {
			context.JSON(http.StatusNotFound, models.ResponseError{Message: err.Error()})
			return
		}

		context.JSON(http.StatusOK, _location)
	})

	group.POST("/", func(context *gin.Context) {
		var _location models.Location
		err := context.BindJSON(&_location)
		if err != nil {
			context.JSON(http.StatusInternalServerError, models.ResponseError{Message: err.Error()})
			return
		}

		err = handler.locationUsecase.New(_location)
		if err != nil {
			context.JSON(http.StatusInternalServerError, models.ResponseError{Message: err.Error()})
			return
		}

		context.String(http.StatusOK, "Successfully inserted location object into database")
	})
}
