package main

import (
    "coastal/internal/pkg/model"
    "coastal/pkg/magic"
    "coastal/test/coastal/util"
    "github.com/stretchr/testify/assert"
    "testing"
)

func TestImageVoteUpdateVote(t *testing.T) {

}

func TestImageVoteSumUp(t *testing.T) {

}

func TestImageVoteSumDown(t *testing.T) {

}

func TestImageVoteSync(t *testing.T) {
    image := model.Image{}
    err := dbConnect.Model(&model.Image{}).First(&image).Error
    assert.Nil(t, err)
    imageId := image.ID

    err = dbConnect.Where("image_id =?", imageId).Delete(&model.ImageVote{}).Error
    assert.Nil(t, err)

    up1 := int64(magic.Num.RandInt(1, 99))
    err = dbConnect.Create(&model.ImageVote{ImageId: imageId, Vote: up1, UserId: util.UserId()}).Error
    assert.Nil(t, err)
    up2 := int64(magic.Num.RandInt(1, 99))
    err = dbConnect.Create(&model.ImageVote{ImageId: imageId, Vote: up2, UserId: util.UserId()}).Error
    assert.Nil(t, err)
    up := up1 + up2

    down1 := -int64(magic.Num.RandInt(1, 99))
    err = dbConnect.Create(&model.ImageVote{ImageId: imageId, Vote: down1, UserId: util.UserId()}).Error
    down2 := -int64(magic.Num.RandInt(1, 99))
    err = dbConnect.Create(&model.ImageVote{ImageId: imageId, Vote: down2, UserId: util.UserId()}).Error
    assert.Nil(t, err)
    down := down1 + down2

    err = model.ImageVote{}.Sync(imageId)
    assert.Nil(t, err)

    image = model.Image{}
    err = dbConnect.Model(&model.Image{}).Where("id=?", imageId).First(&image).Error
    assert.Nil(t, err)
    assert.Equal(t, image.VoteUp, uint64(up))
    assert.Equal(t, image.VoteDown, uint64(-down))
}
