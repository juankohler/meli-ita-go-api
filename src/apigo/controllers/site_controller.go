package controllers

import (
	"github.com/gin-gonic/gin"
	"../services"
	"net/http"
)

const (
	paramSiteID = "siteId"
)

// la idea es q haga las resquest

func GetSiteFromAPI (c *gin.Context){

	siteID :=	c.Param(paramSiteID)
	response ,err := services.GetSite(siteID)
	if err != nil {
		c.JSON(err.Status, err)
		return
	}

	c.JSON(http.StatusOK, response)
}