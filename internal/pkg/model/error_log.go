package model

import "time"

type ErrorLog struct {
    ID          uint64 `gorm:"primary_key"`
    CreatedAt   time.Time
    UpdatedAt   time.Time
    Code        uint64
    Message     string
    Url         string
    Payload     string
    Environment string
    Header      string
    Cookie      string
}
