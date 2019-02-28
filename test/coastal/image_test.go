package main

import (
    "coastal/config/constant"
    "coastal/internal/pkg/model"
    "coastal/internal/pkg/pb"
    "coastal/test/coastal/util"
    "github.com/icrowley/fake"
    "github.com/stretchr/testify/assert"
    "testing"
)

func TestImages(t *testing.T) {
    _, err := testServer.Images(touristCtx(), &pb.ImagesReq{})
    assert.Nil(t, err)

    _, err = testServer.Images(touristCtx(), &pb.ImagesReq{Start: 0, Limit: 10})
    assert.Nil(t, err)

    _, _ = testServer.Images(authCtx(), &pb.ImagesReq{Start: 0, Limit: 10})
    assert.Nil(t, err)

    _, err = testServer.Images(touristCtx(), &pb.ImagesReq{Start: 0, Limit: 10, Random: true})
    assert.Nil(t, err)

    _, err = testServer.Images(touristCtx(), &pb.ImagesReq{Start: 0, Limit: 10, Random: false})
    assert.Nil(t, err)

    _, err = testServer.Images(touristCtx(),
        &pb.ImagesReq{
            Start:  0,
            Limit:  10,
            Random: false,
            Sort:   pb.ImageSorts_ID_DESC,
        })
    assert.Nil(t, err)
}

func TestDownloadImage(t *testing.T) {
    // free download false
    ctx := authCtx()

    err := dbConnect.Model(&model.User{}).Where("id=?", ctx.GetUserId()).Updates(map[string]interface{}{"daily_free_download_number": 0, "score": 0}).Error
    assert.Nil(t, err)

    res, err := testServer.DownloadImage(ctx, &pb.DownloadImageReq{ImageId: 1})
    assert.Nil(t, err)
    assert.Equal(t, false, res.Status)

    // free download true
    ctx = authCtx()
    err = dbConnect.Model(model.User{}).Where("id=?", ctx.GetUserId()).Updates(map[string]interface{}{"daily_free_download_number": 1, "score": 0}).Error
    assert.Nil(t, err)
    res, err = testServer.DownloadImage(ctx, &pb.DownloadImageReq{ImageId: 1})
    assert.Nil(t, err)
    assert.Equal(t, true, res.Status)

    // tourist false
    touristCtx := touristCtx()
    res, err = testServer.DownloadImage(touristCtx, &pb.DownloadImageReq{ImageId: 1})
    assert.NotNil(t, err)

    // score false
    ctx = authCtx()
    err = dbConnect.Model(model.User{}).Where("id=?", ctx.GetUserId()).Updates(map[string]interface{}{"daily_free_download_number": 0, "score": constant.DownloadImageScoreRequired - 1}).Error
    assert.Nil(t, err)
    res, err = testServer.DownloadImage(ctx, &pb.DownloadImageReq{ImageId: 1})
    assert.Nil(t, err)
    assert.Equal(t, false, res.Status)

    // score true
    ctx = authCtx()
    err = dbConnect.Model(model.User{}).Where("id=?", ctx.GetUserId()).Updates(map[string]interface{}{"daily_free_download_number": 0, "score": constant.DownloadImageScoreRequired}).Error
    assert.Nil(t, err)
    res, err = testServer.DownloadImage(ctx, &pb.DownloadImageReq{ImageId: 1})
    assert.Nil(t, err)
    assert.Equal(t, true, res.Status)
    var user model.User
    err = db.Model(model.User{}).Where("id=?", ctx.GetUserId()).First(&user).Error
    assert.Nil(t, err)
    assert.Equal(t, uint64(0), user.Score)
}

func TestUpdateImageVote(t *testing.T) {
    _, err := testServer.UpdateImageVote(touristCtx(), &pb.UpdateImageVoteReq{ImageId: util.ImageId(), Vote: 1})
    assert.NotNil(t, err)

    ctx := authCtx()
    _, err = testServer.UpdateImageVote(ctx, &pb.UpdateImageVoteReq{Vote: 1})
    assert.NotNil(t, err)

    // vote up
    image := model.Image{}
    err = dbConnect.Model(&model.Image{}).First(&image).Error
    assert.Nil(t, err)
    ctx = authCtx()
    _, err = testServer.UpdateImageVote(ctx, &pb.UpdateImageVoteReq{ImageId: image.ID, Vote: 1})
    assert.Nil(t, err)
    imageVote := model.ImageVote{}
    err = dbConnect.Where("image_id=?", image.ID).Where("user_id=?", ctx.GetUserId()).First(&imageVote).Error
    assert.Nil(t, err)
    assert.Equal(t, imageVote.Vote, int64(1))

    // vote down
    image = model.Image{}
    err = dbConnect.Model(&model.Image{}).First(&image).Error
    assert.Nil(t, err)
    ctx = authCtx()
    _, err = testServer.UpdateImageVote(ctx, &pb.UpdateImageVoteReq{ImageId: image.ID, Vote: -1})
    assert.Nil(t, err)
    imageVote = model.ImageVote{}
    err = dbConnect.Where("image_id=?", image.ID).Where("user_id=?", ctx.GetUserId()).First(&imageVote).Error
    assert.Nil(t, err)
    assert.Equal(t, imageVote.Vote, int64(-1))
}

func TestUpdateImageName(t *testing.T) {
    image := model.Image{}
    err := dbConnect.Model(&model.Image{}).First(&image).Error
    assert.Nil(t, err)

    // tourist
    _, err = testServer.UpdateImageName(touristCtx(), &pb.UpdateImageNameReq{ImageId: image.ID, Name: fake.ProductName()})
    assert.NotNil(t, err)

    // success
    ctx := authCtx()
    newName := fake.ProductName()
    err = dbConnect.Model(&model.User{}).Where("id=?", ctx.GetUserId()).Update("score", constant.UpdateImageNameScoreRequired).Error
    assert.Nil(t, err)
    _, err = testServer.UpdateImageName(ctx, &pb.UpdateImageNameReq{ImageId: image.ID, Name: newName})
    assert.Nil(t, err)

    // test review
    review := model.Review{}
    err = dbConnect.Model(&model.Review{}).Where(model.Review{
        ImageId:          image.ID,
        ReviewCategoryId: model.ReviewCategoryIdUpdateImageName,
    }).Order("id DESC").First(&review).Error

    assert.Nil(t, err)
    reviewAttribute := model.ReviewAttribute{}
    err = dbConnect.Model(&model.ReviewAttribute{}).Where("review_id=?", review.ID).First(&reviewAttribute).Error
    assert.Nil(t, err)
    assert.Equal(t, newName, reviewAttribute.ImageName)
}

func TestCreateImageTag(t *testing.T) {
    image := model.Image{}
    err := dbConnect.Model(&model.Image{}).First(&image).Error
    assert.Nil(t, err)

    // tourist
    _, err = testServer.CreateImageTag(touristCtx(), &pb.CreateImageTagReq{ImageId: image.ID, Tag: fake.UserName()})
    assert.NotNil(t, err)

    // success
    ctx := authCtx()
    err = dbConnect.Model(&model.User{}).Where("id=?", ctx.GetUserId()).Update("score", constant.CreateImageTagScoreRequired).Error
    assert.Nil(t, err)

    tagName := fake.UserName()
    _, err = testServer.CreateImageTag(ctx, &pb.CreateImageTagReq{ImageId: image.ID, Tag: tagName})
    assert.Nil(t, err)

    review := model.Review{}
    err = dbConnect.Model(&model.Review{}).Where(model.Review{
        ImageId:          image.ID,
        ReviewCategoryId: model.ReviewCategoryIdCreateImageTag,
    }).Order("id DESC").First(&review).Error
    assert.Nil(t, err)
    reviewAttribute := model.ReviewAttribute{}
    err = dbConnect.Model(&model.ReviewAttribute{}).Where("review_id=?", review.ID).First(&reviewAttribute).Error
    assert.Nil(t, err)
    assert.Equal(t, tagName, reviewAttribute.ImageTagName)

}

func TestDeleteImageTag(t *testing.T) {
    image := model.Image{}
    err := dbConnect.Model(&model.Image{}).First(&image).Error
    assert.Nil(t, err)

    tag := model.Tag{Name: fake.ProductName()}
    err = dbConnect.Create(&tag).Error
    assert.Nil(t, err)

    imageTag := model.ImageTag{ImageId: image.ID, TagId: tag.ID, Confidence: 100}
    err = dbConnect.Create(&imageTag).Error
    assert.Nil(t, err)

    // tourist
    _, err = testServer.DeleteImageTag(touristCtx(), &pb.DeleteImageTagReq{
        ImageId: imageTag.ImageId,
        TagId:   imageTag.TagId,
    })
    assert.NotNil(t, err)

    // success
    ctx := authCtx()
    err = dbConnect.Model(&model.User{},
    ).Where("id=?", ctx.GetUserId(),
    ).Update("score", constant.DeleteImageTagScoreRequired,
    ).Error
    assert.Nil(t, err)

    _, err = testServer.DeleteImageTag(ctx, &pb.DeleteImageTagReq{
        ImageId: imageTag.ImageId,
        TagId:   imageTag.TagId,
    })
    assert.Nil(t, err)

    review := model.Review{}
    err = dbConnect.Model(&model.Review{}).Where(model.Review{
        ImageId:          image.ID,
        ReviewCategoryId: model.ReviewCategoryIdDeleteImageTag,
    }).Order("id DESC").First(&review).Error

    assert.Nil(t, err)
    reviewAttribute := model.ReviewAttribute{}
    err = dbConnect.Model(&model.ReviewAttribute{},
    ).Where("review_id=?", review.ID,
    ).First(&reviewAttribute,
    ).Error
    assert.Nil(t, err)
    assert.Equal(t, imageTag.ImageId, reviewAttribute.ImageId)
    assert.Equal(t, imageTag.TagId, reviewAttribute.TagId)
}

func TestUpdateImageTagVoteReq(t *testing.T) {
    image := model.Image{}
    err := dbConnect.Model(&model.Image{}).First(&image).Error
    assert.Nil(t, err)

    tag := model.Tag{Name: fake.ProductName()}
    err = dbConnect.Create(&tag).Error
    assert.Nil(t, err)

    imageTag := model.ImageTag{ImageId: image.ID, TagId: tag.ID, Confidence: 100}
    err = dbConnect.Create(&imageTag).Error
    assert.Nil(t, err)

    _, err = tServer.UpdateImageTagVote(touristCtx(), &pb.UpdateImageTagVoteReq{
        ImageId: imageTag.ImageId,
        TagId:   imageTag.TagId,
        Vote:    uint64(1),
    })
    assert.NotNil(t, err)

    ctx := authCtx()
    err = db.Model(model.ImageTagVoteLog{},
    ).Where("user_id=?", ctx.GetUserId(),
    ).Update("user_id", 0,
    ).Error
    assert.Nil(t, err)
    _, err = tServer.UpdateImageTagVote(ctx, &pb.UpdateImageTagVoteReq{
        ImageId: imageTag.ImageId,
        TagId:   imageTag.TagId,
        Vote:    uint64(1),
    })
    assert.Nil(t, err)

    count := 0
    err = db.Model(model.ImageTagVoteLog{}).Where(
        "user_id=?", ctx.GetUserId(),
    ).Where("image_id=?", image.ID).Count(&count).Error
    assert.Nil(t, err)
    assert.Equal(t, 1, count)
}
