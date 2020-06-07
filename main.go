package main

import (
	"book-system/config"
	"book-system/routes"

	"github.com/gin-gonic/gin"
)

func main() {

	r := gin.Default()

	r.POST("/reserve", routes.MakeReservation)

	r.Run(config.Conf.Port)
}
