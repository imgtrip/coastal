package model

import (
    "github.com/jinzhu/gorm"
)

type AlbumImage struct {
    CommonFields
    ImageId uint64
    AlbumId uint64
}

func (AlbumImage) TableName() string {
    return "album_images"
}

func (AlbumImage) Paginate(albumId uint64, offset uint64, limit int64, order string) ([]Image, error) {
    var imageIds []uint64

    query := Connect.Model(&AlbumImage{}).Select("image_id").Where("album_id = ?", albumId).Offset(offset)
    if limit >= 0 {
        query = query.Limit(limit)
    }
    err := query.Order(order).Pluck("image_id", &imageIds).Error

    if err != nil {
        return []Image{}, err
    }

    isOfficial, err := Album{}.IsOfficialAlbum(albumId)
    if err != nil {
        return []Image{}, err
    }

    return Image{}.ByIds(imageIds, isOfficial)
}

func (AlbumImage) ImageIdsByAlbumId(albumId uint64) ([]uint64, error) {
    var imageIds []uint64

    if e := Connect.Model(&AlbumImage{}).Select("image_id").Where("album_id = ?", albumId).Pluck("image_id", &imageIds).Error; e != nil {
        return []uint64{}, e
    }

    return imageIds, nil
}

func (a AlbumImage) Create(albumId uint64, imageId uint64) (AlbumImage, error) {
    albumImage := AlbumImage{AlbumId: albumId, ImageId: imageId}

    exist, err := a.ExistWithUnScoped(albumId, imageId)
    if err != nil {
        return AlbumImage{}, err
    }

    if exist {
        err = Connect.Model(&AlbumImage{}).Unscoped().Where("album_id=?", albumId).Where("image_id=?", imageId).Update("deleted_at", gorm.Expr("NULL")).Error
    } else {
        err = Connect.Create(&albumImage).Error
    }

    if err != nil {
        return AlbumImage{}, err
    }

    err = Album{}.Increment(albumId, AlbumAmounts)
    if err != nil {
        return albumImage, err
    }

    return albumImage, err
}

func (AlbumImage) Delete(albumId uint64, imageId uint64) error {
    err := Connect.Where("album_id=?", albumId).Where("image_id=?", imageId).Delete(&AlbumImage{}).Error
    if err != nil {
        return err
    }

    err = Album{}.Decrement(albumId, AlbumAmounts)
    if err != nil {
        return err
    }
    return nil
}

func (AlbumImage) ExistWithUnScoped(albumId uint64, imageId uint64) (bool, error) {
    count := 0
    err := Connect.Model(&AlbumImage{}).Unscoped().Where("album_id=?", albumId).Where("image_id=?", imageId).Count(&count).Error
    if err != nil {
        return false, err
    }
    return count > 0, nil
}
