package mailgun

import (
	"gopkg.in/mailgun/mailgun-go.v1"
)

type Config struct {
	Domain              string
	PrivateAPIKey       string
	PublicValidationKey string
	Sender              string
}

type Client struct {
	Config *Config
	mg     mailgun.Mailgun
}

func New(config *Config) *Client {
	return &Client{
		Config: config,
		mg:     mailgun.NewMailgun(config.Domain, config.PrivateAPIKey, config.PublicValidationKey),
	}
}

func (c *Client) Send(subject, body, recipient string) (string, error) {
	message := c.mg.NewMessage(c.Config.Sender, subject, body, recipient)
	_, id, err := c.mg.Send(message)

	return id, err
}
