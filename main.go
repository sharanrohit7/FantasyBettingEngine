package main

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/sharanrohit7/FantasyBettingEngine/config"
)

func init() {
	config.Dbconn()
	if config.DB != nil {
		fmt.Println("Database connection is ready to use.")

		// You can now use config.DB to interact with the database
	} else {
		fmt.Println("Failed to connect to the database.")
	}

}
func main() {

	router := gin.Default()
	router.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "pong",
		})
	})

	router.Run("localhost:8080")
}
