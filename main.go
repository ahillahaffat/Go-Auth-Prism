package main

import (
	"Go-Auth/config"
	"Go-Auth/routes"

	"github.com/gin-gonic/gin"
)

func main(){
	r := gin.Default()
	config.ConnectDatabase()

	routes.AuthRoutes(r)

	r.Run(":8080")
}