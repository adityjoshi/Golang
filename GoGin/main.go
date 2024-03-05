package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	// this is a handler fucntion
	r := gin.New()
	r.Use(gin.Logger())
	r.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "pong",
			"hey":     "they",
		})
	})
	r.GET("/getdata", getdata)
	r.POST("/getdatapost", getpost)
	r.GET("/query",getquery)
	r.Run()
}

func getdata(c *gin.Context) {
	c.JSON(200, gin.H{
		"hey": "i am a function",
	})

}
func getpost(c *gin.Context) {
	c.JSON(200, gin.H{
		"hey": "post is working fine",
	})
func getquery(c *gin.Context) {
	
}


}
