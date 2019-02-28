package model

import "time"

type ScoreLog struct {
    ID        uint64 `gorm:"primary_key"`
    CreatedAt time.Time

    UserId          uint64
    Score           int64
    ScoreCategoryId uint64
    Description     string
    User            User
    ScoreCategory   ScoreCategory
}

func (ScoreLog) Create(userId uint64, score int64, scoreCategoryId uint64) error {
    return Connect.Create(&ScoreLog{UserId: userId, Score: score, ScoreCategoryId: scoreCategoryId}).Error
}
