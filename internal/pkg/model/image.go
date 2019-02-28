package model

import (
    "coastal/pkg/magic"
    "github.com/jinzhu/gorm"
)

const defaultVoteStep = 1

type Image struct {
    CommonFields
    Name       string
    Src        string
    Downloads  uint64
    VoteUp     uint64
    VoteDown   uint64
    UploaderId uint64
    Published  bool
}

func (Image) Paginate(offset uint64, limit uint64, order string, arg ...bool) ([]Image, error) {
    var images []Image
    query := Connect.Offset(offset).Order(order).Limit(limit)
    if len(arg) > 0 && arg[0] {
        query = query.Where("published=?", arg[0])
    }
    e := query.Find(&images).Error

    return images, e
}

func (Image) ByIds(ids []uint64, arg ...bool) ([]Image, error) {
    var images []Image
    query := DB.Where("id IN (?)", ids)

    if len(arg) > 0 && arg[0] {
        query = query.Where("published=?", arg[0])
    }

    e := query.Find(&images).Error

    return images, e
}

func (Image) ByExcepts(excepts []uint64, limit uint64, published ...bool) ([]Image, error) {
    var images []Image
    query := DB.Where("id NOT IN (?)", excepts).Offset(magic.Num.RandInt(1, 2000))

    if len(published) > 0 {
        query = query.Where("published=?", published[0])
    }

    e := query.Limit(limit).Find(&images).Error
    return images, e
}

func (Image) UpdateVote(imageId uint64, up bool) error {
    expr := ""
    column := ""
    if up {
        expr = "vote_up + ?"
        column = "vote_up"
    } else {
        expr = "vote_down + ?"
        column = "vote_down"
    }

    err := DB.Model(Image{}).Where("id=?", imageId).Update(column, gorm.Expr(expr, defaultVoteStep)).Error

    return err
}

func (Image) UpdateVotes(imageId uint64, up uint64, down uint64) error {
    err := DB.Model(&Image{}).Where("id=?", imageId).Updates(map[string]uint64{"vote_up": up, "vote_down": down}).Error
    if err != nil {
        return err
    }

    return nil
}

func (Image) ByUploaderId(uploader uint64) (Image, error) {
    image := Image{}
    err := DB.Where("uploader_id", uploader).First(&image).Error
    return image, err
}

func (Image) ById(id uint64) (Image, error) {
    image := Image{}
    err := DB.Where("id=?", id).First(&image).Error
    return image, err
}
