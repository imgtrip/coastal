package coastal

import (
    "coastal/internal/app/coastal/response"
    "coastal/internal/app/coastal/types"
    "coastal/internal/app/coastal/validator"
    "coastal/internal/pkg/errors"
    "coastal/internal/pkg/model"
    "coastal/internal/pkg/pb"
    "golang.org/x/net/context"
)

func (*Server) Reviews(c context.Context, req *pb.ReviewsReq) (*pb.ReviewsRes, error) {
    res := &pb.ReviewsRes{}
    auth, err := Auth(c, cachedUser)
    if err != nil {
        return res, err
    }

    reviewCategoryIds, err := model.ReviewCategory{}.GetIdsByRequiredScore(auth.Score)
    if err != nil {
        return res, err
    }

    if err := validator.Reviews(req, auth.ID); err != nil {
        return res, err
    }

    var reviews []model.Review
    query := model.DB.Where("review_category_id IN (?)", reviewCategoryIds)

    order := "updated_at DESC"

    switch req.Order {
    case pb.ReviewsOrder_CREATE:
        order = "created_at DESC"
        break
    case pb.ReviewsOrder_ENDING:
        query = query.Where("closed=?", false)
        order = "end_at ASC"
        break
    case pb.ReviewsOrder_CREATOR_HISTORY:
        query = query.Where("creator_id=?", auth.ID)
        order = "created_at DESC"
        break
    case pb.ReviewsOrder_ENDED:
        query = query.Where("closed=?", true)
        order = "end_at DESC"
        break
        // FIXME
        // case pb.ReviewsOrder_REVIEWER_HISTORY:
    }

    err = query.Preload("ReviewCategory",
    ).Preload("Creator",
    ).Preload("Image",
    ).Preload("ReviewAttribute",
    ).Preload("ReviewLog",
    ).Offset(req.Start,
    ).Limit(req.Limit,
    ).Order(order,
    ).Find(&reviews,
    ).Error

    if err != nil {
        return res, err
    }
    var reviewRelations []types.ReviewWithImageRelation

    for _, item := range reviews {

        var images []model.Image
        imageWithRelations, err := imageWithRelations(auth, append(images, item.Image))
        if err != nil {
            return res, errors.Internal(err.Error())
        }

        content, err := reviewContent(item.ReviewCategoryId, item.ReviewAttribute)
        if err != nil {
            return res, errors.Internal(err.Error())
        }

        reviewRelations = append(reviewRelations, types.ReviewWithImageRelation{
            Review:            item,
            Content:           content,
            ImageWithRelation: imageWithRelations[0],
        })
    }

    total := 0
    if req.Total {
        err = model.DB.Model(model.Review{}).Where("review_category_id IN (?)", reviewCategoryIds,
        ).Count(&total,
        ).Error
        if err != nil {
            return res, errors.Internal(err.Error())
        }
    }

    return response.Reviews(reviewRelations, total), nil
}

func (*Server) UpdateReviewOpinion(c context.Context, req *pb.UpdateReviewOpinionReq) (*pb.UpdateReviewOpinionRes, error) {
    res := &pb.UpdateReviewOpinionRes{}
    auth, err := Auth(c, cachedUser)
    if err != nil {
        return res, err
    }

    if err := validator.UpdateReviewOpinion(req, auth.ID); err != nil {
        return res, err
    }

    err = model.ReviewLog{}.CreateOrUpdate(req.ReviewId, auth.ID, req.Opinion)

    if err != nil {
        return res, errors.Internal(err.Error())
    }

    err = model.Review{}.SyncOpinionsCount(req.ReviewId)

    return res, err
}
