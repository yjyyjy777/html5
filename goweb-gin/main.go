package main

import (



	"net/http"

	"github.com/gin-gonic/gin"
)
func main() {
    r :=gin.Default()
	// r :=gin.Default()
	// r.GET("/index", func(c *gin.Context) {
	// 	c.JSON(http.StatusOK,gin.H{
	// 		"message":"ok",
	// 	})

	// })
	r.GET("/index",func(c *gin.Context){
		c.Redirect(http.StatusMovedPermanently,"http://www.baidu.com")
	
		
	})
	r.Run(":9000")




}

