package email

// Notification contains the information in the queue for an email notification
type Notification struct {
	Template         string            `json:"template"`
	Subject          string            `json:"subject"`
	ReceiverName     string            `json:"receiver_name"`
	ReceiverEmail    string            `json:"receiver_email"`
	PlainTextContent string            `json:"plain_text_content"`
	Data             map[string]string `json:"data"`
}
