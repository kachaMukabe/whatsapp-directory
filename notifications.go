package main

type Notification struct {
	Object string   `json:"object"`
	Entry  []*Entry `json:"entry"`
}

type Entry struct {
	ID      string     `json:"id"`
	Changes []*Changes `json:"changes"`
}

type Changes struct {
	Value *Value `json:"value"`
}

type Value struct {
	Contacts         []*Contacts    `json:"contacts"`
	Errors           []*ValueErrors `json:"errors"`
	MessagingProduct string         `json:"messaging_product"`
	Messages         []*Messages    `json:"messages"`
	Metadata         *Metadata      `json:"metadata"`
	Statuses         []*Statuses    `json:"statuses"`
}

type Messages struct {
	Audio       *Audio          `json:"audio"`
	Button      *Button         `json:"button"`
	Context     *MessageContext `json:"context"`
	Document    *Document       `json:"document"`
	Errors      *ValueErrors    `json:"errors"`
	From        string          `json:"from"`
	ID          string          `json:"id"`
	Identity    *Identity       `json:"identity"`
	Image       *Image          `json:"image"`
	Interactive *Interactive    `json:"interactive"`
	//Order *messageOrder `json:"order"`
	//Referral *referral `json:"referral"`
	//Sticker *sticker `json:"sticker"`
	//System *system `json:"system"`
	Text      *MessageText `json:"text"`
	Timestamp string       `json:"timestamp"`
	Type      string       `json:"type"`
	//Video     *video       `json:"video"`
}

type Audio struct {
	ID       string `json:"id"`
	MimeType string `json:"mime_type"`
}

type Button struct {
	Payload string `json:"payload"`
	Text    string `json:"text"`
}

type MessageContext struct {
	Forwarded           bool             `json:"forwarded"`
	FrequentlyForwarded bool             `json:"frequently_forwarded"`
	From                string           `json:"from"`
	ID                  string           `json:"id"`
	ReferredProduct     *ReferredProduct `json:"referred_product"`
}

type ReferredProduct struct {
	CatalogID         string `json:"catalog_id"`
	ProductRetailerID string `json:"product_retailer_id"`
}

type Document struct {
	Caption  string `json:"caption"`
	Filename string `json:"filename"`
	Sha256   string `json:"sha256"`
	MimeType string `json:"mime_type"`
	ID       string `json:"id"`
}

type Identity struct {
	Acknowledged     string `json:"acknowledged"`
	CreatedTimestamp string `json:"created_timestamp"`
	Hash             string `json:"hash"`
}

type Image struct {
	Caption  string `json:"caption"`
	Sha256   string `json:"sha256"`
	MimeType string `json:"mime_type"`
	ID       string `json:"id"`
}

type Interactive struct {
	Type        string       `json:"type"`
	ButtonReply *ButtonReply `json:"button_reply"`
	ListReply   *ListReply   `json:"list_reply"`
}

type ButtonReply struct {
	ID    string `json:"id"`
	Title string `json:"title"`
}

type ListReply struct {
	ID          string `json:"id"`
	Title       string `json:"title"`
	Description string `json:"description"`
}

type MessageText struct {
	Body string `json:"body"`
}

type Contacts struct {
	WaId    string   `json:"wa_id"`
	Profile *Profile `json:"profile"`
}

type Profile struct {
	Name string `json:"name"`
}

type ValueErrors struct {
	Code      int64      `json:"code"`
	Type      string     `json:"type"`
	Message   string     `json:"message"`
	ErrorData *ErrorData `json:"error_data"`
	Details   string     `json:"details"`
}

type ErrorData struct {
	MessagingProduct string `json:"messaging_product"`
	Details          string `json:"details"`
}

type Metadata struct {
	DisplayPhoneNumber string `json:"display_phone_number"`
	PhoneNumberId      string `json:"phone_number_id"`
}

type Statuses struct {
	Conversation *Conversation  `json:"conversation"`
	Errors       []*ValueErrors `json:"errors"`
	ID           string         `json:"id"`
	Pricing      *Pricing       `json:"pricing"`
	RecipientID  string         `json:"recipient_id"`
	Status       string         `json:"status"`
	Timestamp    string         `json:"timestamp"`
}

type Conversation struct {
	Id                  string  `json:"id"`
	Origin              *Origin `json:"origin"`
	ExpirationTimestamp string  `json:"expiration_timestamp"`
}

type Origin struct {
	Type string `json:"type"`
}

type Pricing struct {
	Billable     bool   `json:"billable"`
	PricingModel string `json:"pricing_model"`
	Category     string `json:"user_initiated"`
}
