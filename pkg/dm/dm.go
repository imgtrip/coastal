package dm

import (
    "gopkg.in/gomail.v2"
)

type Config struct {
    Host     string
    User     string
    Password string
    Port     int
    From     string
}

type Client struct {
    Config Config
    Dialer *gomail.Dialer
}

func New(c Config) *Client {
    return &Client{
        Config: c,
        Dialer: gomail.NewPlainDialer(c.Host, c.Port, c.User, c.Password),
    }
}

func (c *Client) Send(to string, subject string, body string) error {
    m := gomail.NewMessage()
    m.SetHeader("From", c.Config.From)
    m.SetHeader("To", to)
    m.SetHeader("Subject", subject)
    m.SetBody("text/html", body)

    if err := c.Dialer.DialAndSend(m); err != nil {
        return err
    }

    return nil
}
