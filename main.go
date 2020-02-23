package main

import (
	"encoding/json"
	"io/ioutil"
	"os"

	"github.com/anothrNick/email-notifications/config"
	"github.com/anothrNick/email-notifications/email"
	"github.com/go-redis/redis"
	"github.com/sirupsen/logrus"
)

const (
	// QueueEmailNotifications is the Redis queue to listen
	QueueEmailNotifications = "email_notifications_queue"
)

func main() {
	// load config
	configPath := os.Getenv("EMAIL_CONFIG_PATH")
	file, _ := ioutil.ReadFile(configPath)

	config := &config.App{}
	json.Unmarshal([]byte(file), &config)

	// initialize logger
	logger := logrus.New()
	logger.Formatter = &logrus.JSONFormatter{}

	logger.Info("starting email notifications")

	// initialize email client
	client := email.NewClient(config, logger)

	// create a new redis client
	queue := redis.NewClient(&redis.Options{
		Addr:     "cache:6379",
		Password: "", // no password set
		DB:       0,  // use default DB
	})

	// ping the redis server
	pong, err := queue.Ping().Result()
	// fail quickly if ping fails
	if err != nil {
		logger.WithFields(logrus.Fields{"pong": pong}).WithError(err).Error("error pinging redis")
		return
	}

	logger.Info("waiting for notifications...")
	for {
		// endlessly read from queue
		result, err := queue.BLPop(0, QueueEmailNotifications).Result()

		// exit on a read error
		if err != nil {
			logger.Error(err)
			return
		}

		// unmarshal event
		notification := &email.Notification{}
		if err := json.Unmarshal([]byte(result[1]), notification); err != nil {
			logger.Error(err)
			continue
		}

		logger.WithFields(logrus.Fields{"notification": notification}).Info("received notification")
		// send event
		if err := client.Send(notification); err != nil {
			logger.Error(err)
		}
	}
}
