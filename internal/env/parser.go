package env

import (
	"github.com/kelseyhightower/envconfig"
	"log"
)

type Special struct {
	Debug             bool   `envconfig:"debug" default:"true"`
	DbHost            string `envconfig:"db_host" default:"localhost"`
	DbPort            string `envconfig:"db_port" default:"3306"`
	DbName            string `envconfig:"db_name"`
	DbUser            string `envconfig:"db_user"`
	DbPassword        string `envconfig:"db_password"`
	ImageMin          int    `envconfig:"image_min" default:"1"`
	ImageMax          int    `envconfig:"image_max" default:"100"`
	NodeServerToken   string `envconfig:"node_server_token" default:"NodeJs0000"`
	MailGunPrivateKey string `envconfig:"mailgun_private_key"`
	MailGunPublicKey  string `envconfig:"mailgun_public_key"`
	SMTPPassword      string `envconfig:"smtp_password"`
}

var special Special

func parse() {
	err := envconfig.Process("imgtrip", &special)
	if err != nil {
		log.Fatal(err.Error())
	}
}
