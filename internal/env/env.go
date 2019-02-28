package env

import (
	"coastal/config/constant"
)

type process struct {
	DB
	Config
	Email
	SMTPEmail
	Debug bool
}

var Process *process

type DB struct {
	DBHost     string
	DBPort     string
	DBName     string
	DBUser     string
	DBPassword string
}

type Config struct {
	NodeServerToken string
	ImageMinId      uint64
	ImageMaxId      uint64
}

type Email struct {
	Domain     string
	Sender     string
	PrivateKey string
	PublicKey  string
}

type SMTPEmail struct {
	Host     string
	User     string
	Password string
	Port     int
	From     string
}

func New() *process {
	parse()
	Process = &process{
		DB: DB{
			DBHost:     special.DbHost,
			DBPort:     special.DbPort,
			DBUser:     special.DbUser,
			DBPassword: special.DbPassword,
			DBName:     special.DbName,
		},
		Config: Config{
			NodeServerToken: special.NodeServerToken,
			ImageMinId:      uint64(special.ImageMin),
			ImageMaxId:      uint64(special.ImageMax),
		},
		Email: Email{
			Domain:     constant.EmailDomain,
			Sender:     constant.EmailSender,
			PrivateKey: special.MailGunPrivateKey,
			PublicKey:  special.MailGunPublicKey,
		},
		SMTPEmail: SMTPEmail{
			Host:     constant.SMTPHost,
			User:     constant.SMTPUser,
			Password: special.SMTPPassword,
			Port:     constant.SMTPPort,
			From:     constant.SMTPFrom,
		},
		Debug: special.Debug,
	}

	return Process
}

func GetDBConfig() *DB {
	return &DB{
		DBHost:     Process.DBHost,
		DBPort:     Process.DBPort,
		DBName:     Process.DBName,
		DBUser:     Process.DBUser,
		DBPassword: Process.DBPassword,
	}
}
