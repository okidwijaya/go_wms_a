package main

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/okidwijaya/go_wms_a/config"
	"github.com/okidwijaya/go_wms_a/controllers"
)

func main() {
	fmt.Println("Starting...")
	config.DbConn()
	router := gin.Default()
	defer config.DB.Close()

	router.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"Message": "Ping Connection Successfully"})
	})

	router.POST("/receiver", controllers.ReceiverController)
	router.POST("/dispatching", controllers.DispatchingController)

	router.Run(":8080")
}
