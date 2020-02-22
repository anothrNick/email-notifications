package main

import (
	"encoding/json"
	"io/ioutil"
	"os"

	"github.com/anothrNick/email-notifications/config"
	"github.com/anothrNick/email-notifications/email"
)

func main() {
	// load config
	configPath := os.Getenv("EMAIL_CONFIG_PATH")
	file, _ := ioutil.ReadFile(configPath)

	config := &config.App{}
	json.Unmarshal([]byte(file), &config)

	// test notification
	// this will be read from redis queue
	notification := &email.Notification{
		Subject:          "Test Email from Local",
		ReceiverName:     "Nick",
		ReceiverEmail:    "nick@machinable.io",
		PlainTextContent: "Email confirmation code: 0000",
		Template:         "default",
		Data: map[string]string{
			"Name":    "Nick Sjostrom",
			"Site":    "Machinable",
			"URL":     "https://machinable.io/login",
			"Company": "Machinable",
		},
	}

	client := email.NewClient(config)
	client.Send(notification)
}
