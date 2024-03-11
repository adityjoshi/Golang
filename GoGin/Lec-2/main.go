package main

import (
	"net/http"
	"time"

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
	//r.GET("/getdata", getdata)
	//r.POST("/getdatapost", getpost)
	//r.GET("/query", getquery)
	r.GET("/urlData/:name/:age", getUrlData)
	// to run it in custom port no => http.ListenAndServe(":2426", r)

	// Router Grouping

	admin := r.Group("/admin")
	{
		admin.GET("/getdata", getdata)
		admin.POST("/getdatapost", getpost)
	}

	client := r.Group("/client")
	{
		client.GET("/query", getquery)
	}

	server := &http.Server{
		Addr:         ":2426",
		Handler:      r,
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 10 * time.Second,
	}
	server.ListenAndServe()
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
}

// http://localhost:8080/query?name=adi&age=19
func getquery(c *gin.Context) {
	name := c.Query("name")
	age := c.Query("age ")
	c.JSON(200, gin.H{
		"hey":  "Query String is working fine",
		"name": name,
		"age":  age,
	})
}

// http://localhost:8080/urlData/19

func getUrlData(c *gin.Context) {
	name := c.Param("name")
	age := c.Param("age")
	c.JSON(200, gin.H{
		"hey":  "Url data working",
		"name": name,
		"age":  age,
	})

}
