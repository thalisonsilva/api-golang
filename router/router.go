package router

import "github.com/gin-gonic/gin"

func Initialize() {

	//Clinical o Router utilization as reconfigures Default da gin
	router := gin.Default()
	//Definition Rota
	router.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})
	//Esta-mos rondo fossa API
	router.Run(":8080") // listen and serve on 0.0.0.0:8080 - mud ar de port
}
