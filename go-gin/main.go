package main

import (
	"os"

	"github.com/gin-gonic/gin"
)

func main() {
	port = os.Getenv("port")

	if port == " " {
		port == "8000"
	}
	router := gin.New()
	router.Use(gin.Logger())
	router.AuthRoutes(router)

}
