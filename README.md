# email-notifications

Reads email notifications from a Redis queue and sends them via SendGrid API.

[![Stable Version](https://img.shields.io/github/v/tag/anothrNick/email-notifications)](https://img.shields.io/github/v/tag/anothrNick/email-notifications)

### Environment

Set local environment variables

```sh
echo "export SENDGRID_API_KEY='YOUR_API_KEY'" > dev.env
echo "export EMAIL_CONFIG_PATH='./config/sample-config.json'" >> dev.env
source ./dev.env
```

### Config

Reads a JSON Config file. Set the `EMAIL_CONFIG_PATH` environment variable with the path of the config file.

|Field|Description|
|-----|-----------|
|TemplateMap|A map of template names to HTML template file paths|
|SenderName|The name of the email sender|
|SenderEmail|The email of the sender|

### Run Locally

Make commands:

```sh
# build without cache
rebuild:
	docker-compose build --no-cache

# build image
build:
	docker-compose build

# spin up containers (postgres and api)
up:
	docker-compose up -d

# bring down running containers
down:
	docker-compose down

# stop running containers
stop:
	docker-compose stop

# remove containers and images
remove:
	docker-compose rm -f

# cleanup images and volumes
clean:
	docker-compose down --rmi all -v --remove-orphans
```
