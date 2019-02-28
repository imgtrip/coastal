package model

import (
    "time"
)

type ImageTagVoteLog struct {
    ID         uint64 `gorm:"primary_key"`
    CreatedAt  time.Time
    ImageTagId uint64
    ImageId    uint64
    TagId      uint64
    UserId     uint64
    Vote       uint64
}

func (ImageTagVoteLog) IsExist(imageId uint64, tagId uint64, userId uint64) (bool, error) {
    count := 0
    err := DB.Model(&ImageTagVoteLog{}).Where(
        "image_id=?", imageId,
    ).Where(
        "tag_id=?", tagId,
    ).Where(
        "user_id=?", userId,
    ).Count(&count).Error
    return count > 0, err
}

func (ImageTagVoteLog) Update(imageId uint64, tagId uint64, userId uint64, vote uint64) error {
    return DB.Model(ImageTagVoteLog{}).Where(&ImageTagVoteLog{
        ImageId: imageId,
        TagId:   tagId,
        UserId:  userId,
    }).Update("vote", vote).Error
}

func (ImageTagVoteLog) Create(log ImageTagVoteLog) error {
    imageTag, err := ImageTag{}.ByImageAndTag(log.ImageId, log.TagId)
    if err != nil {
        return err
    }
    log.CreatedAt = time.Now().Local()
    log.ImageTagId = imageTag.ID

    return DB.Create(&log).Error
}

func (ImageTagVoteLog) ByUserImageTag(userId uint64, imageId uint64, tagId uint64) (ImageTagVoteLog, error) {
    var log ImageTagVoteLog
    err := DB.Where(
        "user_id=?", userId,
    ).Where(
        "image_id=?", imageId,
    ).Where(
        "tag_id=?", tagId,
    ).First(&log).Error

    return log, err
}
