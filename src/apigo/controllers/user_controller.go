package controllers

import (
	"github.com/gin-gonic/gin"
	"../services"
	"net/http"
	"strconv"
	"../utils"
)

const (
	paramUserID = "userID"
)

// la idea es q haga las resquest

func GetUserFromAPI (c *gin.Context){

	userID :=	c.Param(paramUserID)

	id, err := strconv.Atoi(userID)
	if err != nil {
		apiErr := &utils.ApiError{
			Message: err.Error(),
			Status: http.StatusBadRequest,
		}
		c.JSON(apiErr.Status, apiErr)
		return
	}

	response ,apiErr := services.GetUser(id)
	if apiErr != nil {
		c.JSON(apiErr.Status, apiErr)
		return
	}

	c.JSON(http.StatusOK, response)
}