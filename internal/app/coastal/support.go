package coastal

import (
    "coastal/internal/app/coastal/types"
    "coastal/internal/env"
    "coastal/internal/pkg/model"
    "coastal/pkg/magic"
    "github.com/jinzhu/gorm"
)

func RandImages(s string, length int) ([]model.Image, error) {
    excepts, err := ImageSessionModel.ImageIds(s)
    if err != nil {
        return []model.Image{}, err
    }

    randN := magic.Num.RandIntArrButExcept(int(env.Process.ImageMinId), int(env.Process.ImageMaxId), length, magic.Num.Uint64ArrToIntArr(excepts), []int{})
    rands := magic.Num.IntArrToUint64(randN)

    images, err := ImageModel.ByIds(rands, true)

    if err != nil {
        return []model.Image{}, err
    }
    diff := length - len(images)
    if diff > 0 {
        // FIXME 如果已展示数据过大，可能导致where in sql过长异常
        patch, err := ImageModel.ByExcepts(append(excepts, rands...), uint64(diff), true)
        if err != nil {
            return []model.Image{}, err
        }
        images = append(images, patch...)
    }

    return images, nil
}

func saveImagesSession(s string, images []model.Image) {
    var ids []uint64
    for _, img := range images {
        ids = append(ids, img.ID)
    }
    ImageSessionModel.BatchSave(s, ids)
}

func sendVerificationEmail(recipient string, token string) error {
    code, err := server.Email.SendVerificationCodeEmail(recipient)
    if err != nil {
        return err
    }
    _, err = model.VerificationCode{}.SimpleCreate(recipient, code, token)
    return err
}

func favoriteIdsFilter(imageIds []uint64, albumId uint64) ([]uint64, error) {
    var ids []uint64
    err := server.DBConnect.Model(&model.AlbumImage{}).Where(
        "image_id IN (?)", imageIds,
    ).Where(
        "album_id=?", albumId,
    ).Pluck("image_id", &ids).Error
    return ids, err
}

func imageWithRelations(user model.User, images []model.Image) ([]types.ImageWithRelation, error) {
    var imageWithTags []types.ImageWithRelation

    type ImageTag struct {
        ID    uint64
        Name  string
        TagId uint64
    }

    var err error
    var ids []uint64
    var favoriteIds []uint64

    if user.ID > 0 && user.AlbumId > 0 {
        for _, image := range images {
            ids = append(ids, image.ID)
        }

        favoriteIds, err = favoriteIdsFilter(ids, user.AlbumId)
        if err != nil {
            return imageWithTags, err
        }
    }

    // FIXME 效率低
    for _, image := range images {
        var imageTags []ImageTag

        err := model.DB.Model(model.ImageTag{}).Select(
            " tags.name,image_tag.id,image_tag.tag_id",
        ).Joins(
            "left join tags on image_tag.tag_id = tags.id",
        ).Where(
            "image_tag.image_id=?", image.ID,
        ).Scan(&imageTags).Error
        if err != nil {
            return imageWithTags, err
        }

        var tagWithRelations []types.TagWithRelation
        for _, imageTag := range imageTags {
            var votedTagLog model.ImageTagVoteLog
            if user.ID > 0 {
                votedTagLog, err = model.ImageTagVoteLog{}.ByUserImageTag(user.ID, image.ID, imageTag.TagId)

                if err != nil && err != gorm.ErrRecordNotFound {
                    return imageWithTags, err
                }
            }

            tagWithRelations = append(tagWithRelations, types.TagWithRelation{
                TagID:      imageTag.TagId,
                TagName:    imageTag.Name,
                ImageTagId: imageTag.ID,
                Voted:      votedTagLog.Vote > 0,
            })
        }

        var withFavorite types.ImageWithFavorite
        withFavorite.Image = image
        withFavorite.Favorite = magic.Arr.HasUint64(image.ID, favoriteIds)

        imageWithTags = append(imageWithTags, types.ImageWithRelation{
            Image: withFavorite,
            Tags:  tagWithRelations,
        })
    }

    return imageWithTags, nil
}

func reviewContent(reviewCategoryId uint64, attr model.ReviewAttribute) (string, error) {
    var content string
    var err error
    var tag model.Tag
    switch reviewCategoryId {
    case model.ReviewCategoryIdUpdateImageName:
        content = attr.ImageName
        break
    case model.ReviewCategoryIdCreateImageTag:
        content = attr.ImageTagName
        break
    case model.ReviewCategoryIdDeleteImageTag:
        tag, err = model.Tag{}.ById(attr.TagId)
        content = tag.Name
        break
    }

    return content, err
}
