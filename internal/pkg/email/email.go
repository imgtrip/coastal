package email

import (
    "coastal/pkg/dm"
    "coastal/pkg/magic"
    "fmt"
)

const (
    VerificationEmailSubject = "旅图网验证码: %s "
    VerificationEmailContent = "<html><body>旅图网验证码: <b>%s</b></body></html>"
)

type Client struct {
    Config  dm.Config
    Postman *dm.Client
}

func New(conf dm.Config) *Client {
    return &Client{
        Config:  conf,
        Postman: dm.New(conf),
    }
}

func (c *Client) SendVerificationCodeEmail(to string) (string, error) {
    code := magic.Num.IntToString(magic.Num.RandInt(100000, 999999))
    subject := fmt.Sprintf(VerificationEmailSubject, code)
    content := fmt.Sprintf(VerificationEmailContent, code)
    err := c.Postman.Send(to, subject, content)
    return code, err
}
