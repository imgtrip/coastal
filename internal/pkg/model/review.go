package model

import (
    "coastal/config/constant"
    "coastal/internal/pkg/pb"
    "time"
)

type Review struct {
    CommonFields
    CreatorId        uint64
    ImageId          uint64
    ReviewCategoryId uint64
    AgreeCount       uint64
    DisagreeCount    uint64
    EndAt            time.Time
    ReviewCategory   ReviewCategory
    Closed           bool
    Creator          User
    Image            Image
    ReviewAttribute  ReviewAttribute
    ReviewLog        ReviewLog
}

// func (Review) Paginate(offset uint64, limit uint64, order string, categoryIds []uint64, args ...uint64) ([]Review, error) {
//
//     if len(args) == 1 {
//
//     }
//
//     var reviews []Review
//     err := DB.Preload("ReviewCategory",
//     ).Preload("Creator",
//     ).Preload("Image",
//     ).Preload("ReviewAttribute",
//     ).Preload("ReviewLog",
//     ).Where("review_category_id IN (?)", categoryIds,
//     ).Offset(offset,
//     ).Limit(limit,
//     ).Order(" DESC",
//     ).Find(&reviews,
//     ).Error
//
//     return reviews, err
// }

func (Review) SaveUpdateImageName(imageId uint64, creatorId uint64, imageName string) (Review, error) {
    review := Review{
        CreatorId:        creatorId,
        ImageId:          imageId,
        ReviewCategoryId: ReviewCategoryIdUpdateImageName,
        EndAt:            time.Now().Add(time.Second * constant.UpdateImageNameReviewExpireSeconds),
    }

    err := DB.Create(&review).Error
    if err != nil {
        return Review{}, err
    }

    err = ReviewAttribute{}.SaveUpdateImageNameAttribute(review.ID, imageName, imageId)
    if err != nil {
        return Review{}, err
    }

    return review, err
}

func (Review) SaveCreateImageTag(imageId uint64, creatorId uint64, tagName string) (Review, error) {
    review := &Review{
        CreatorId:        creatorId,
        ImageId:          imageId,
        ReviewCategoryId: ReviewCategoryIdCreateImageTag,
        EndAt:            time.Now().Add(time.Second * constant.CreateImageTagReviewExpireSeconds),
    }

    err := DB.Create(review).Error
    if err != nil {
        return Review{}, err
    }

    err = ReviewAttribute{}.SaveCreateImageTagAttribute(review.ID, tagName, imageId)
    if err != nil {
        return Review{}, err
    }

    return *review, err
}

func (Review) SaveDeleteImageTag(creatorId uint64, imageId uint64, tagId uint64) (Review, error) {
    review := &Review{
        CreatorId:        creatorId,
        ImageId:          imageId,
        ReviewCategoryId: ReviewCategoryIdDeleteImageTag,
        EndAt:            time.Now().Add(time.Second * constant.DeleteImageTagReviewExpireSeconds),
    }

    err := DB.Create(review).Error
    if err != nil {
        return Review{}, err
    }

    err = ReviewAttribute{}.SaveDeleteImageTagAttribute(review.ID, tagId, imageId)
    if err != nil {
        return Review{}, err
    }

    return *review, err
}

func (Review) CanReviewByScore(score uint64, reviewId uint64) (bool, error) {
    ids, err := ReviewCategory{}.GetIdsByRequiredScore(score)
    if err != nil {
        return false, err
    }

    if len(ids) == 0 {
        return false, nil
    }

    count := 0
    err = DB.Model(Review{}).Where("id=?", reviewId).Where("review_category_id IN (?)", ids).Count(&count).Error
    if err != nil {
        return false, err
    }

    return count > 0, nil
}

func (Review) SyncOpinionsCount(reviewId uint64) error {
    agreeCount, err := ReviewLog{}.CountById(reviewId, pb.ReviewOpinions_AGREE)
    if err != nil {
        return err
    }

    disagreeCount, err := ReviewLog{}.CountById(reviewId, pb.ReviewOpinions_DISAGREE)
    if err != nil {
        return err
    }

    return DB.Model(Review{}).Where("id=?", reviewId,
    ).Updates(map[string]interface{}{
        "agree_count":    agreeCount,
        "disagree_count": disagreeCount,
    }).Error
}
