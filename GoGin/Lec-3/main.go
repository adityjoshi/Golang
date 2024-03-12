package main

import (
	"GoGin/Lec-3/middleware"
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.New()
	// here the middle ware is applied to all the routes
	r.Use(middleware.Authenticate)
	r.Use(gin.Logger())
	r.GET("/getData1", firstData)
	r.GET("/getData2", secondData)
	// to apply middle ware to specific route we can do this => r.GET("/getData3", middleware.Authenticate, secondData)
	r.GET("/getData3", thirdData)

	server := &http.Server{
		Addr:    ":2426",
		Handler: r,
	}
	server.ListenAndServe()

}

func firstData(c *gin.Context) {
	c.JSON(200, gin.H{
		"data": "data 1 ",
	})
}
func secondData(c *gin.Context) {
	c.JSON(200, gin.H{
		"data": "data 1 ",
	})
}
func thirdData(c *gin.Context) {
	c.JSON(200, gin.H{
		"data": "data 1 ",
	})
}
