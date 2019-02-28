package model

import (
    "fmt"
    _ "github.com/go-sql-driver/mysql"
    "github.com/jinzhu/gorm"
    "time"
)

type Manager struct {
    Config  *Config
    Connect *gorm.DB
}

type Config struct {
    Host     string `ini:"host"`
    Port     string `ini:"port"`
    DBName   string `ini:"db_name"`
    User     string `ini:"user"`
    Password string `ini:"password"`
}

type CommonFields struct {
    ID        uint64 `gorm:"primary_key"`
    CreatedAt time.Time
    UpdatedAt time.Time
    DeletedAt *time.Time `sql:"index"`
}

var Connect *gorm.DB
var DB *gorm.DB

func New() *Manager {
    return &Manager{}
}

func (m *Manager) SetConfig(config *Config) *Manager {
    m.Config = config
    return m.setConnect()
}

func (m *Manager) setConnect() *Manager {
    source := fmt.Sprintf(
        "%s:%s@tcp(%s:%s)/%s?charset=utf8&parseTime=True&loc=Local",
        m.Config.User,
        m.Config.Password,
        m.Config.Host,
        m.Config.Port,
        m.Config.DBName,
    )

    conn, err := gorm.Open("mysql", source)
    if err != nil {
        panic(err)
    }
    Connect = conn
    DB = conn
    m.Connect = conn
    return m
}
