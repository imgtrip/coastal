package model

import "time"

type ReviewCategory struct {
    ID                    uint64 `gorm:"primary_key"`
    CreatedAt             time.Time
    UpdatedAt             time.Time
    Name                  string
    Description           string
    CreatorScore          int64
    ReviewerScore         int64
    ReviewerScoreRequired uint64
}

const (
    ReviewCategoryIdUpdateImageName      = 1
    ReviewCategoryIdCreateImageTag       = 2
    ReviewCategoryIdDeleteImageTag       = 3
    ReviewCategoryIdDeleteImage          = 4
    // ReviewCategoryIdDeleteDuplicateImage = 5
)

func (ReviewCategory) GetIdsByRequiredScore(score uint64) ([]uint64, error) {
    var ids []uint64
    err := DB.Model(&ReviewCategory{}).Where("reviewer_score_required<=?", score).Pluck("id", &ids).Error
    return ids, err
}


