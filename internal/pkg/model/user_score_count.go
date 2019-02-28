package model

import "time"

type UserScoreCount struct {
    ID              uint64 `gorm:"primary_key"`
    CreatedAt       time.Time
    UpdatedAt       time.Time
    ScoreCategoryId uint64
    UserId          uint64
}
