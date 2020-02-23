package email

import (
	"bytes"
	"html/template"
	"os"

	"github.com/anothrNick/email-notifications/config"
	"github.com/sendgrid/sendgrid-go"
	"github.com/sendgrid/sendgrid-go/helpers/mail"
	"github.com/sirupsen/logrus"
)

// Client wraps the email configuration and sending
type Client struct {
	config *config.App
	client *sendgrid.Client
	logger *logrus.Logger
}

// NewClient returns a `Client` with the default `SendEmail` func
func NewClient(config *config.App, logger *logrus.Logger) *Client {
	return &Client{
		config: config,
		client: sendgrid.NewSendClient(os.Getenv("SENDGRID_API_KEY")),
		logger: logger,
	}
}

// Send sends an email with the `Notification` contents
func (c *Client) Send(n *Notification) error {
	message, err := c.buildEmail(n)

	if err != nil {
		c.logger.WithFields(logrus.Fields{"meta": n.Meta}).WithError(err).Error("error buildEmail")
		return err
	}

	response, serr := c.client.Send(message)

	if serr != nil {
		c.logger.WithFields(logrus.Fields{"meta": n.Meta}).WithError(err).Error("error Send")
		return err
	}

	c.logger.WithFields(logrus.Fields{"status_code": response.StatusCode, "meta": n.Meta}).Info("sent notification")
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
		c.logger.WithFields(logrus.Fields{"meta": n.Meta}).WithError(err).Error("error parseTemplate")
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
