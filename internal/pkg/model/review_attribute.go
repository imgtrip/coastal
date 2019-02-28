package model

import "time"

type ReviewAttribute struct {
    ID           uint64 `gorm:"primary_key"`
    CreatedAt    time.Time
    UpdatedAt    time.Time
    ReviewId     uint64
    ImageName    string
    ImageTagName string
    ImageId      uint64
    TagId        uint64
}

func (ReviewAttribute) SaveUpdateImageNameAttribute(reviewId uint64, imageName string, imageId uint64) error {
    return Connect.Create(&ReviewAttribute{
        ReviewId:  reviewId,
        ImageName: imageName,
        ImageId:   imageId,
    }).Error
}

func (ReviewAttribute) SaveCreateImageTagAttribute(reviewId uint64, imageTagName string, imageId uint64) error {
    return Connect.Create(&ReviewAttribute{
        ReviewId:     reviewId,
        ImageTagName: imageTagName,
        ImageId:      imageId,
    }).Error
}

func (ReviewAttribute) SaveDeleteImageTagAttribute(reviewId uint64, tagId uint64, imageId uint64) error {
    return Connect.Create(&ReviewAttribute{
        ReviewId: reviewId,
        TagId:    tagId,
        ImageId:  imageId,
    }).Error
}
