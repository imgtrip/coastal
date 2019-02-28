package model

import "coastal/config/constant"

type ImageVote struct {
    CommonFields
    ImageId uint64
    UserId  uint64
    Vote    int64
}

func (ImageVote) TableName() string {
    return "image_votes"
}

func (ImageVote) UpdateVote(userId uint64, imageId uint64, vote int64) error {
    count := 0
    err := Connect.Model(&ImageVote{}).Where(&ImageVote{UserId: userId, ImageId: imageId}).Count(&count).Error
    if err != nil {
        return err
    }

    exist := count > 0

    if exist {
        err = Connect.Model(&ImageVote{}).Where(&ImageVote{UserId: userId, ImageId: imageId}).Update("vote", vote).Error
    } else {
        err = Connect.Create(&ImageVote{UserId: userId, ImageId: imageId, Vote: vote}).Error
        if vote > 0 {
            image, err := Image{}.ById(imageId)
            if err != nil {
                return err
            }
            err = User{}.UpdateScore(image.UploaderId, constant.ImageVoteUpUploaderScore, ScoreCategoryImageUp)
        }
    }

    return err
}

func (ImageVote) ByUserId(userId uint64) ([]ImageVote, error) {
    imageVotes := []ImageVote{}

    err := Connect.Where("user_id=?", userId).Find(&imageVotes).Error
    return imageVotes, err
}

func (ImageVote) SumUp(imageId uint64) (uint64, error) {
    sum := struct {
        Sum int
    }{}
    err := Connect.Model(&ImageVote{}).Select("sum(vote) as sum").Where(&ImageVote{ImageId: imageId}).Where("vote > ?", 0).Scan(&sum).Error
    if err != nil {
        return 0, err
    }
    return uint64(sum.Sum), nil
}

func (ImageVote) SumDown(imageId uint64) (uint64, error) {
    sum := struct {
        Sum int
    }{}
    err := Connect.Model(&ImageVote{}).Select("sum(vote) as sum").Where(&ImageVote{ImageId: imageId}).Where("vote < ?", 0).Scan(&sum).Error
    if err != nil {
        return 0, err
    }
    return uint64(-sum.Sum), nil
}

func (i ImageVote) Sync(imageId uint64) error {
    up, err := i.SumUp(imageId)
    if err != nil {
        return err
    }

    down, err := i.SumDown(imageId)
    if err != nil {
        return err
    }
    err = Image{}.UpdateVotes(imageId, up, down)

    return err
}
