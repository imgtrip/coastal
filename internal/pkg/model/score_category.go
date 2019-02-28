package model

import "time"

type ScoreCategory struct {
    ID        uint64 `gorm:"primary_key"`
    CreatedAt time.Time
    UpdatedAt time.Time
    Name      string
    Symbol    string
}

const (
    ScoreCategoryImageUp              = 1
    ScoreCategoryDownload             = 2
    ScoreCategoryUpdateImageName      = 3
    ScoreCategoryCreateImageTag       = 4
    ScoreCategoryDeleteImageTag       = 5
    ScoreCategoryAgreeImageTag        = 6
    ScoreCategoryReviewCreateImageTag = 7
    ScoreCategoryReviewDeleteImageTag = 8
    ScoreCategoryReviewImageDown      = 9
    ScoreCategoryImageDown            = 10
    ScoreCategoryAlbumCollected       = 11
    ScoreCategoryFlagExistImage       = 12
)

