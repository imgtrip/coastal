package model

import (
    "time"
)

type ImageTag struct {
    ID         uint64 `gorm:"primary_key"`
    CreatedAt  time.Time
    UpdatedAt  time.Time
    ImageId    uint64
    TagId      uint64
    Confidence uint64
    VoteUp     uint64
    Tag        Tag
}

func (ImageTag) TableName() string {
    return "image_tag"
}

func (ImageTag) Find(imageId uint64, tagId uint64) (ImageTag, error) {
    var imageTag ImageTag
    err := DB.Where("image_id=?", imageId).Where("tag_id=?", tagId).First(&imageTag).Error
    return imageTag, err
}

func (ImageTag) ById(id uint64) (ImageTag, error) {
    var imageTag ImageTag
    err := DB.Where("id=?", id).First(&imageTag).Error
    return imageTag, err
}

func (ImageTag) ByImageAndTag(imageId uint64, tagId uint64) (ImageTag, error) {
    var imageTag ImageTag
    err := DB.Where(
        "image_id=?", imageId,
    ).Where(
        "tag_id=?", tagId,
    ).First(&imageTag).Error
    return imageTag, err
}

func (ImageTag) GetTagById(id uint64) (Tag, error) {
    var imageTag ImageTag
    err := DB.Preload("Tag").Where("id=?", id).First(&imageTag).Error
    return imageTag.Tag, err
}

func (ImageTag) SyncVote(imageId uint64, tagId uint64) error {
    type Result struct {
        VoteCount int
    }
    var result Result
    err := DB.Model(ImageTagVoteLog{}).Select(" sum(vote) AS vote_count",
    ).Where("image_id=?", imageId,
    ).Where("tag_id=?", tagId,
    ).Where("vote>?", 0,
    ).Scan(&result).Error

    if err != nil {
        return err
    }

    return DB.Model(ImageTag{}).Where("image_id=?", imageId,
    ).Where("tag_id=?", tagId,
    ).Update("vote_up", result.VoteCount,
    ).Error
}
