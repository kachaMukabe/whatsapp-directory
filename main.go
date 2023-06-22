package main

import (
	"bytes"
	"encoding/json"
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
	//fmt.Println(notification)
	//fmt.Printf("%#v\n", notification)
	sendMessage(notification.Entry[0].Changes[0].Value.Messages[0].Text.Body)

	c.JSON(http.StatusOK, notification)

}

func sendMessage(text string) {
	body := TextObject{Body: text, PreviewUrl: true}
	message := MessageObject{Text: &body, To: "+260966581925", Type: "text", RecipientType: "individual", MessagingProduct: "whatsapp"}
	phoneNumber := os.Getenv("PHONENUMBER")
	apiToken := os.Getenv("FBTOKEN")
	b, err := json.Marshal(message)
	if err != nil {
		log.Println(err)
	}

	client := &http.Client{}

	req, err := http.NewRequest("POST", "https://graph.facebook.com/v17.0/"+phoneNumber+"/messages", bytes.NewBuffer(b))
	if err != nil {
		log.Fatal(err)
	}
	req.Header.Add("Authorization", "Bearer "+apiToken)
	req.Header.Add("Content-Type", "application/json")

	resp, err := client.Do(req)
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()
	log.Println(resp.Status)

	//resp, err := http.Post("", "application/json", b)

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
