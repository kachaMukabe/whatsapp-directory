package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func getEnvVariable(key string) string {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatalf("Error loading .env file")
	}

	return os.Getenv("VERIFYTOKEN")
}

func verifyWebhook(c *gin.Context) {
	mode := c.Query("hub.mode")
	challenge := c.Query("hub.challenge")
	verifyToken := c.Query("hub.verify_token")
	token := getEnvVariable("VERIFYTOKEN")

	if mode == "subscribe" && verifyToken == token {
		c.String(http.StatusOK, "%s", challenge)
	}
}

func handleMessageNotification(c *gin.Context) {
	var notification Notification
	if err := c.BindJSON(&notification); err != nil {
		return
	}
	fmt.Println(notification)
	fmt.Printf("%#v\n", notification)

	c.JSON(http.StatusOK, notification.Entry[0].Changes[0].Value.Messages[0].Text.Body)

}

func main() {
	router := gin.Default()
	router.GET("/webhook", verifyWebhook)
	router.POST("/webhook", handleMessageNotification)

	router.Run("localhost:8000")
}
