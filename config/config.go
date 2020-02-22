package config

// App holds all of the email-notifications configuration
type App struct {
	TemplateMap map[string]string // a map of key -> html file
	SenderName  string
	SenderEmail string
}
