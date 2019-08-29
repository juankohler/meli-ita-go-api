package main

import (
	"github.com/gin-gonic/gin"
	"../controllers"
	"log"
	"../utils"
)

const (
	port = ":8081"
)

var (
	router = gin.Default() // el router por defecto de gin

)


func main() {
	utils.NewCircuitBreaker(5,3,0, utils.CLOUSE)
	router.GET("/users/:userID", controllers.GetUserFromAPI) // GetUserFromAPI necesita un parametro pero lo saca de contesxto de :userID
	// path y handler (quien es el contorlador de ese)
	router.GET("/countries/:countryID", controllers.GetCountryFromAPI)
	router.GET("/sites/:siteID", controllers.GetSiteFromAPI)
	router.GET("/results/:userID", controllers.GetResultFromAPI)

	// Vamos a lanzar el gin gonic

	log.Fatal(router.Run(port))
}
