package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
)

func verifyWebhook(c *gin.Context) {
	mode := c.Query("hub.mode")
	challenge := c.Query("hub.challenge")
	verifyToken := c.Query("hub.verify_token")
	token := os.Getenv("VERIFYTOKEN")
	log.Print(token)

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
	fmt.Println(notification.Entry[0].Changes[0].Value.Messages)

	c.JSON(http.StatusOK, notification.Entry[0].Changes[0].Value.Messages[0].Text.Body)

}

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}
	router := gin.Default()
	router.GET("/webhook", verifyWebhook)
	router.POST("/webhook", handleMessageNotification)

	router.Run(":" + port)
}
