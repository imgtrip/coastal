package model

import (
    "coastal/config/constant"
    "time"
)

type FreeDownloadLog struct {
    ID        uint64 `gorm:"primary_key"`
    CreatedAt time.Time
    ImageId   uint64
    UserId    uint64
    Date      string
}

func (FreeDownloadLog) CountBy(userId uint64, date time.Time) (int, error) {
    count := 0
    err := Connect.Model(&FreeDownloadLog{}).Where(FreeDownloadLog{UserId: userId, Date: date.Format(constant.DateFormat)}).Count(&count).Error
    return count, err
}

func (FreeDownloadLog) Create(userId uint64, imageId uint64) error {
    return Connect.Create(&FreeDownloadLog{UserId: userId, ImageId: imageId, Date: time.Now().Local().Format(constant.DateFormat)}).Error
}
