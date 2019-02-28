package main

import (
    "coastal/internal/pkg/model"
    "coastal/internal/pkg/pb"
    "coastal/pkg/magic"
    "github.com/icrowley/fake"
    "github.com/stretchr/testify/assert"
    "testing"
    "time"
)

func TestReviews(t *testing.T) {
    // tourist false
    _, err := testServer.Reviews(touristCtx(), &pb.ReviewsReq{Start: 0, Limit: 10})
    assert.NotNil(t, err)

    // create fake review rows
    ctx := authCtx()
    var image model.Image
    err = db.Model(model.Image{}).First(&image).Error
    assert.Nil(t, err)

    for i := 1; i <= 20; i++ {
        reviewCategoryId := magic.Num.RandInt(model.ReviewCategoryIdUpdateImageName, model.ReviewCategoryIdDeleteImage)
        review := model.Review{
            ImageId:          image.ID,
            ReviewCategoryId: uint64(reviewCategoryId),
            CreatorId:        ctx.GetUserId(),
            EndAt:            time.Now().Local(),
        }
        err = db.Create(&review).Error
        assert.Nil(t, err)

        if reviewCategoryId == model.ReviewCategoryIdDeleteImageTag {
            tag := model.Tag{
                Name: fake.ProductName(),
            }
            err = db.Create(&tag).Error
            assert.Nil(t, err)

            err = model.ReviewAttribute{}.SaveDeleteImageTagAttribute(review.ID,
                tag.ID,
                review.ImageId,
            )
            assert.Nil(t, err)
        }

    }

    // request success
    var reviewCategory model.ReviewCategory
    err = db.Model(model.ReviewCategory{}).Order("reviewer_score_required DESC").First(&reviewCategory).Error
    assert.Nil(t, err)
    err = db.Model(model.User{}).Where("id=?", ctx.GetUserId()).Update("score", reviewCategory.ReviewerScoreRequired).Error
    assert.Nil(t, err)

    res, err := tServer.Reviews(ctx, &pb.ReviewsReq{Start: 0, Limit: 10})
    assert.Nil(t, err)
    assert.Equal(t, true, len(res.Items) > 0)
    assert.Equal(t, true, res.Items[0].Image.Id > 0)

    // only score satisfied
    err = db.Model(model.ReviewCategory{}).Order("reviewer_score_required ASC").First(&reviewCategory).Error
    assert.Nil(t, err)
    err = db.Model(model.User{}).Where("id=?", ctx.GetUserId()).Update("score", reviewCategory.ReviewerScoreRequired).Error
    assert.Nil(t, err)
    err = db.Model(model.Review{}).Update("review_category_id", 0).Error
    assert.Nil(t, err)

    satisfiedRows := 1
    err = db.Model(model.Review{}).Limit(satisfiedRows).Update("review_category_id", reviewCategory.ID).Error
    assert.Nil(t, err)
    res, err = tServer.Reviews(ctx, &pb.ReviewsReq{Start: 0, Limit: 10})
    assert.Nil(t, err)
    assert.Equal(t, satisfiedRows, len(res.Items))
}

func TestUpdateReviewOpinion(t *testing.T) {
    var review model.Review
    err := db.Model(model.Review{}).First(&review).Error
    assert.Nil(t, err)

    // tourist false
    _, err = tServer.UpdateReviewOpinion(touristCtx(), &pb.UpdateReviewOpinionReq{ReviewId: review.ID, Opinion: pb.ReviewOpinions_AGREE})
    assert.NotNil(t, err)

    // fail
    ctx := authCtx()
    err = db.Model(model.User{}).Where("id=?", ctx.GetUserId()).Update("score", 0).Error
    assert.Nil(t, err)
    _, err = tServer.UpdateReviewOpinion(ctx, &pb.UpdateReviewOpinionReq{ReviewId: review.ID, Opinion: pb.ReviewOpinions_AGREE})
    assert.NotNil(t, err)

    // create success
    var reviewCategory model.ReviewCategory
    err = db.Model(review).Related(&reviewCategory).Error
    assert.Nil(t, err)
    err = db.Model(model.User{}).Where("id=?", ctx.GetUserId()).Update("score", reviewCategory.ReviewerScoreRequired).Error
    assert.Nil(t, err)
    _, err = tServer.UpdateReviewOpinion(ctx, &pb.UpdateReviewOpinionReq{ReviewId: review.ID, Opinion: pb.ReviewOpinions_AGREE})
    assert.Nil(t, err)

    // success
    err = db.Model(review).Related(&reviewCategory).Error
    assert.Nil(t, err)
    err = db.Model(model.User{}).Where("id=?", ctx.GetUserId()).Update("score", reviewCategory.ReviewerScoreRequired).Error
    assert.Nil(t, err)

    err = db.Model(model.ReviewLog{}).Where("user_id=?", ctx.GetUserId()).Update("user_id", 0).Error
    assert.Nil(t, err)

    reviewLog := model.ReviewLog{ReviewId: review.ID, UserId: ctx.GetUserId(), Opinion: pb.ReviewOpinions_AGREE}
    err = db.Create(&reviewLog).Error
    assert.Nil(t, err)

    _, err = tServer.UpdateReviewOpinion(ctx, &pb.UpdateReviewOpinionReq{ReviewId: review.ID, Opinion: pb.ReviewOpinions_DISAGREE})
    assert.Nil(t, err)

    err = db.Where("id=?", reviewLog.ID).First(&reviewLog).Error
    assert.Nil(t, err)
    assert.Equal(t, pb.ReviewOpinions_DISAGREE, reviewLog.Opinion)
}
