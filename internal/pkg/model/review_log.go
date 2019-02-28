package model

import (
    "coastal/internal/pkg/pb"
    "time"
)

type ReviewLog struct {
    ID        uint64 `gorm:"primary_key"`
    CreatedAt time.Time
    ReviewId  uint64
    UserId    uint64
    Opinion   pb.ReviewOpinions
}

func (ReviewLog) Create(reviewId uint64, userId uint64, opinion pb.ReviewOpinions) (ReviewLog, error) {
    reviewLog := ReviewLog{ReviewId: reviewId, UserId: userId, Opinion: opinion}
    err := Connect.Create(&reviewLog).Error
    return reviewLog, err
}

func (ReviewLog) Update(reviewId uint64, userId uint64, opinion pb.ReviewOpinions) error {
    return Connect.Model(&ReviewLog{}).Where(ReviewLog{ReviewId: reviewId, UserId: userId}).Update("opinion", opinion).Error
}

func (ReviewLog) IsExist(reviewId uint64, userId uint64) (bool, error) {
    count := 0
    err := Connect.Model(&ReviewLog{}).Where(ReviewLog{ReviewId: reviewId, UserId: userId}).Count(&count).Error
    return count > 0, err
}

func (ReviewLog) CountById(reviewId uint64, opinion pb.ReviewOpinions) (int, error) {
    count := 0
    err := Connect.Model(&ReviewLog{}).Where(ReviewLog{ReviewId: reviewId, Opinion: opinion}).Count(&count).Error
    return count, err
}

func (r ReviewLog) CreateOrUpdate(reviewId uint64, userId uint64, opinion pb.ReviewOpinions) error {
    isExist, err := r.IsExist(reviewId, userId)
    if err != nil {
        return err
    }

    if isExist {
        err = r.Update(reviewId, userId, opinion)
    } else {
        _, err = r.Create(reviewId, userId, opinion)
    }

    return err
}
