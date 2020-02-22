package email

import (
	"bytes"
	"fmt"
	"html/template"
	"log"
	"os"

	"github.com/anothrNick/email-notifications/config"
	"github.com/sendgrid/sendgrid-go"
	"github.com/sendgrid/sendgrid-go/helpers/mail"
)

// Client wraps the email configuration and sending
type Client struct {
	config *config.App
	client *sendgrid.Client
}

// NewClient returns a `Client` with the default `SendEmail` func
func NewClient(config *config.App) *Client {
	return &Client{
		config: config,
		client: sendgrid.NewSendClient(os.Getenv("SENDGRID_API_KEY")),
	}
}

// Send sends an email with the `Notification` contents
func (c *Client) Send(n *Notification) error {
	message, err := c.buildEmail(n)

	if err != nil {
		return err
	}

	response, serr := c.client.Send(message)

	if serr != nil {
		log.Println(err)
		return err
	}

	fmt.Println(response.StatusCode)
	fmt.Println(response.Body)
	fmt.Println(response.Headers)
	return nil
}

func (c *Client) parseTemplate(templateFileName string, data interface{}) (string, error) {
	t, err := template.ParseFiles(templateFileName)

	if err != nil {
		return "", err
	}

	buf := new(bytes.Buffer)
	if err = t.Execute(buf, data); err != nil {
		return "", err
	}
	msgBody := buf.String()

	return msgBody, nil
}

func (c *Client) buildEmail(n *Notification) (*mail.SGMailV3, error) {
	template, ok := c.config.TemplateMap[n.Template]
	if !ok {
		// default
		template = "./templates/default.html"
	}
	htmlContent, err := c.parseTemplate(template, n.Data)

	if err != nil {
		return nil, err
	}

	from := mail.NewEmail(c.config.SenderName, c.config.SenderEmail)
	to := mail.NewEmail(n.ReceiverName, n.ReceiverEmail)

	return mail.NewSingleEmail(
		from,
		n.Subject,
		to,
		n.PlainTextContent,
		htmlContent), nil
}
