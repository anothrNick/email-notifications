# email-notifications

Reads email notifications from a Redis queue and sends them via SendGrid API.

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

### TODO

* Read notifications from redis
