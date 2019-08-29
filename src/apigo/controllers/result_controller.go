package controllers

import (
	"github.com/gin-gonic/gin"
	"strconv"
	"net/http"
	"../services"
	"../utils"
	"fmt"
	"time"
)

func GetResultFromAPI(c *gin.Context) {
	/*
		userID := c.Param(paramUserID)
		id, err := strconv.Atoi(userID)
		if err != nil {
			apiErr := &utils.ApiError{
				Message: err.Error(),
				Status:  http.StatusBadRequest,
			}
			c.JSON(apiErr.Status, apiErr)
			return
		}
		response, apiErr := services.GetResult(id)

	if apiErr != nil {
			c.JSON(apiErr.Status, apiErr)
			return
		}else {
			c.JSON(http.StatusOK, response)
		}


	*/
		statusCB := utils.CB.State

		if statusCB == utils.CLOUSE{
			userID := c.Param(paramUserID)
			id, err := strconv.Atoi(userID)
			if err != nil {
				apiErr := &utils.ApiError{
					Message: "Intente nuevamente",
					Status:  http.StatusBadRequest,
				}
				c.JSON(apiErr.Status, apiErr)
				return
			}
			response, apiErr := services.GetResult(id)
			if apiErr != nil {
				utils.CB.AddCountRetries();
				if utils.CB.CountRetries >= utils.CB.MaxRetries {
					utils.CB.SetState(utils.OPEN)
					go verificarApi()
					println("Estado cambió a OPEN")
				}
				apiErr := &utils.ApiError{
					Message: "Intente nuevamente",
					Status:  http.StatusBadRequest,
				}
				c.JSON(apiErr.Status, apiErr)
				return
			}else {
				c.JSON(http.StatusOK, response)
			}
		} else if (statusCB == utils.OPEN) || (statusCB == utils.HALFOPEN){
			apiErr := &utils.ApiError{
				Message: "Por favor intente mas tarde",
				Status:  http.StatusBadRequest,
			}
			c.JSON(apiErr.Status, apiErr)
			return

		} 

}

func verificarApi()  {


	for  {
		time.Sleep(5 * time.Second)
		utils.CB.SetState(utils.HALFOPEN)
		url := fmt.Sprintf("%s", utils.UrlPing)
		println(url)
		response, err := http.Get(url)
		if err != nil {
			println("Error - API CAIDA")
			utils.CB.SetState(utils.OPEN)
		}
		if response != nil {
			println("Response - API FUNCIONANDO")
			utils.CB.SetState(utils.CLOUSE)
			break
		}
	}

}