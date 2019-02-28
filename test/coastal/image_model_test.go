package main

import (
    "coastal/internal/pkg/model"
    "coastal/pkg/magic"
    "github.com/stretchr/testify/assert"
    "testing"
)

func TestImageByIds(t *testing.T) {
    image := model.Image{}
    err := dbConnect.First(&image).Error
    assert.Nil(t, err)
    err = dbConnect.Model(&model.Image{}).Where("id=?", image.ID).Update("published", false).Error
    assert.Nil(t, err)

    // published only
    images, err := model.Image{}.ByIds([]uint64{image.ID}, true)
    assert.Nil(t, err)
    assert.Equal(t, 0, len(images))

    // include not published
    images, err = model.Image{}.ByIds([]uint64{image.ID}, false)
    assert.Nil(t, err)
    assert.Equal(t, 1, len(images))
}

func TestImageUpdateVote(t *testing.T) {
    changeVote := func(up bool) {
        image := model.Image{}
        err := dbConnect.Model(&model.Image{}).First(&image).Error
        assert.Nil(t, err)
        err = model.Image{}.UpdateVote(image.ID, up)
        assert.Nil(t, err)
        imageAssert := model.Image{}
        err = dbConnect.Model(&model.Image{}).First(&imageAssert).Error
        assert.Nil(t, err)

        var changedVote uint64
        if up {
            changedVote = image.VoteUp + 1
            assert.Equal(t, changedVote, imageAssert.VoteUp)
        } else {
            changedVote = image.VoteDown + 1
            assert.Equal(t, changedVote, imageAssert.VoteDown)
        }
    }

    changeVote(true)
    changeVote(false)
}

func TestImageUpdateVotes(t *testing.T) {
    image := model.Image{}
    err := dbConnect.Model(&model.Image{}).First(&image).Error
    assert.Nil(t, err)

    up := uint64(magic.Num.RandInt(1, 999))
    down := uint64(magic.Num.RandInt(1, 999))

    err = model.Image{}.UpdateVotes(image.ID, up, down)
    assert.Nil(t, err)

    assertImage := model.Image{}
    err = dbConnect.Model(&model.Image{}).Where("id=?", image.ID).First(&assertImage).Error
    assert.Nil(t, err)
    assert.Equal(t, assertImage.VoteUp, up)
    assert.Equal(t, assertImage.VoteDown, down)
}
