package validator

import (
    "coastal/internal/pkg/errors"
    "coastal/internal/pkg/model"
    "coastal/internal/pkg/pb"
)

func Reviews(req *pb.ReviewsReq, authId uint64) error {
    if authId == 0 {
        return errors.PermissionDenied()
    }

    if err := Structure(req); err != nil {
        return errors.InvalidArgument(err.Error())
    }

    return nil
}

func reviewOpinionPermission(authId uint64, reviewId uint64) error {
    user, err := model.User{}.ById(authId)
    if err != nil {
        return errors.Internal(err.Error())
    }

    can, err := model.Review{}.CanReviewByScore(user.Score, reviewId)
    if err != nil {
        return errors.Internal(err.Error())
    }

    if !can {
        return errors.PermissionDenied()
    }

    return nil
}


func UpdateReviewOpinion(req *pb.UpdateReviewOpinionReq, authId uint64) error {
    if authId == 0 {
        return errors.PermissionDenied()
    }

    if err := Structure(req); err != nil {
        return errors.InvalidArgument(err.Error())
    }

    return reviewOpinionPermission(authId, req.ReviewId)
}
