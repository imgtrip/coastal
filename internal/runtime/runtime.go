package runtime

import (
	"coastal/internal/env"
	"coastal/internal/pkg/email"
	"coastal/internal/pkg/model"
	"coastal/pkg/dm"
	"github.com/jinzhu/gorm"
)

type Runtime struct {
	DB    *model.Manager
	Email *email.Client
}

var runtime *Runtime

func New() *Runtime {
	dmConfig := email.New(dm.Config{
		Host:     env.Process.SMTPEmail.Host,
		User:     env.Process.SMTPEmail.User,
		Password: env.Process.SMTPEmail.Password,
		Port:     env.Process.SMTPEmail.Port,
		From:     env.Process.SMTPEmail.From,
	})

	runtime = &Runtime{
		DB: model.New().SetConfig(&model.Config{
			Host:     env.Process.DBHost,
			User:     env.Process.DBUser,
			Password: env.Process.DBPassword,
			DBName:   env.Process.DBName,
			Port:     env.Process.DBPort,
		}),
		Email: dmConfig,
	}
	return runtime
}

func DBConnect() *gorm.DB {
	return runtime.DB.Connect
}

func Email() *email.Client {
	return runtime.Email
}
