package email

// Notification contains the information in the queue for an email notification
type Notification struct {
	Template         string            `json:"template"`
	Subject          string            `json:"subject"`
	ReceiverName     string            `json:"receiver_name"`
	ReceiverEmail    string            `json:"receiver_email"`
	PlainTextContent string            `json:"plain_text_content"` // plain text in the case email html can't render
	Data             map[string]string `json:"data"`               // data for template
	Meta             map[string]string `json:"meta"`               // meta data for logging
}
