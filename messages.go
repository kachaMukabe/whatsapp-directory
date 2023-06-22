package main

type MessageObject struct {
	MessagingProduct string      `json:"messaging_product"`
	RecipientType    string      `json:"recipient_type"`
	Text             *TextObject `json:"text"`
	To               string      `json:"to"`
	Type             string      `json:"type"`
}

type TextObject struct {
	Body       string `json:"body"`
	PreviewUrl bool   `json:"preview_url"`
}
