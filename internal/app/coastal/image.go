package coastal

import (
    "coastal/config/constant"
    "coastal/internal/app/coastal/response"
    "coastal/internal/app/coastal/validator"
    "coastal/internal/env"
    "coastal/internal/pkg/errors"
    "coastal/internal/pkg/model"
    "coastal/internal/pkg/pb"
    "context"
    "time"
)

func (*Server) Images(c context.Context, req *pb.ImagesReq) (*pb.ImagesRes, error) {
    res := &pb.ImagesRes{}
    auth, err := Auth(c, cachedUser)
    if err != nil {
        return res, err
    }

    limit := req.Limit
    if limit > 50 {
        limit = 50
    }

    var images []model.Image
    if req.Random {
        token, err := GetTokenFromContext(c)
        if err != nil {
            return res, err
        }

        images, err = RandImages(token, int(limit))
        if err != nil {
            return res, errors.Internal(err.Error())
        }

        if !env.Process.Debug {
            saveImagesSession(token, images)
        }
    } else {
        sorts := map[pb.ImageSorts]string{pb.ImageSorts_ID_ASC: "id ASC", pb.ImageSorts_ID_DESC: "id DESC"}

        images, err = model.Image{}.Paginate(req.Start, limit, sorts[req.Sort])
        if err != nil {
            return res, errors.Internal(err.Error())
        }
    }

    imageWithRelations, err := imageWithRelations(auth, images)
    if err != nil {
        return res, errors.Internal(err.Error())
    }

    total := 0
    if req.Total {
        err = model.DB.Model(model.Image{}).Where("published=?", true).Count(&total).Error
        if err != nil {
            return res, errors.Internal(err.Error())
        }
    }

    return response.Images(imageWithRelations, total), nil
}

func (*Server) DownloadImage(c context.Context, req *pb.DownloadImageReq) (*pb.DownloadImageRes, error) {
    res := &pb.DownloadImageRes{}
    auth, err := Auth(c, cachedUser)
    if err != nil {
        return res, err
    }

    if err := validator.DownloadImage(req, auth.ID); err != nil {
        return res, err
    }

    freeStatus, err := model.User{}.CanFreeDownload(auth.ID)
    if err != nil {
        return res, err
    }

    scoreStatus := false
    if !freeStatus {
        scoreStatus, err = model.User{}.IsScoreGreater(auth.ID, constant.DownloadImageScoreRequired)
        if err != nil {
            return res, err
        }
    }

    // FIXME 优先扣费可能导致事务失败，却扣费成功
    if freeStatus {
        err = model.FreeDownloadLog{}.Create(auth.ID, req.ImageId)
    } else if scoreStatus {
        err = model.User{}.UpdateScore(auth.ID, -constant.DownloadImageScoreRequired, model.ScoreCategoryDownload)
    }

    if err != nil {
        return res, err
    }

    return response.DownloadImage(freeStatus || scoreStatus), nil

}

func (*Server) UpdateImageVote(c context.Context, req *pb.UpdateImageVoteReq) (*pb.UpdateImageVoteRes, error) {
    res := &pb.UpdateImageVoteRes{}
    auth, err := Auth(c, cachedUser)
    if err != nil {
        return res, err
    }

    if err := validator.UpdateImageVote(req, auth.ID); err != nil {
        return res, err
    }

    err = model.ImageVote{}.UpdateVote(auth.ID, req.ImageId, req.Vote)
    if err != nil {
        return res, errors.Internal(err.Error())
    }

    go func() {
        err := model.ImageVote{}.Sync(req.ImageId)
        if err != nil {
            // Log error
        }
    }()

    return res, nil
}

func (*Server) UpdateImageName(c context.Context, req *pb.UpdateImageNameReq) (*pb.UpdateImageNameRes, error) {
    res := &pb.UpdateImageNameRes{}
    auth, err := Auth(c, cachedUser)
    if err != nil {
        return res, err
    }

    if err := validator.UpdateImageName(req, auth.ID); err != nil {
        return res, err
    }

    _, err = model.Review{}.SaveUpdateImageName(req.ImageId, auth.ID, req.Name)
    if err != nil {
        return res, errors.Internal(err.Error())
    }

    return res, nil
}

func (*Server) CreateImageTag(c context.Context, req *pb.CreateImageTagReq) (*pb.CreateImageTagRes, error) {
    res := &pb.CreateImageTagRes{}
    auth, err := Auth(c, cachedUser)
    if err != nil {
        return res, err
    }

    if err := validator.CreateImageTag(req, auth.ID); err != nil {
        return res, err
    }

    _, err = model.Review{}.SaveCreateImageTag(req.ImageId, auth.ID, req.Tag)
    if err != nil {
        return res, errors.Internal(err.Error())
    }

    return res, nil
}

func (*Server) DeleteImageTag(c context.Context, req *pb.DeleteImageTagReq) (*pb.DeleteImageTagRes, error) {
    res := &pb.DeleteImageTagRes{}
    auth, err := Auth(c, cachedUser)
    if err != nil {
        return res, err
    }

    if err := validator.DeleteImageTag(req, auth.ID); err != nil {
        return res, err
    }

    var reviews []model.Review
    err = model.DB.Model(model.Review{},
    ).Where("image_id=?", req.ImageId,
    ).Where("end_at > ?", time.Now(),
    ).Where("review_category_id=?", model.ReviewCategoryIdDeleteImageTag,
    ).Find(&reviews).Error
    if err != nil {
        return res, errors.Internal(err.Error())
    }

    if len(reviews) == 0 {
        _, err = model.Review{}.SaveDeleteImageTag(auth.ID, req.ImageId, req.TagId)
        if err != nil {
            return res, errors.Internal(err.Error())
        }
    } else {
        var reviewIds []uint64
        for _, item := range reviews {
            reviewIds = append(reviewIds, item.ID)
        }

        if len(reviewIds) > 0 {
            count := 0
            err := model.DB.Model(model.ReviewAttribute{},
            ).Where("review_id IN (?)", reviewIds,
            ).Where("tag_id=?", req.TagId,
            ).Count(&count).Error

            if err != nil {
                return res, errors.Internal(err.Error())
            }

            if count > 0 {
                return res, errors.AlreadyExists()
            }

            _, err = model.Review{}.SaveDeleteImageTag(auth.ID, req.ImageId, req.TagId)
            if err != nil {
                return res, errors.Internal(err.Error())
            }
        } else {
            // data error
        }
    }

    return res, nil
}

func (*Server) UpdateImageTagVote(c context.Context, req *pb.UpdateImageTagVoteReq) (*pb.UpdateImageTagVoteRes, error) {
    res := &pb.UpdateImageTagVoteRes{}
    auth, err := Auth(c, cachedUser)
    if err != nil {
        return res, err
    }

    if err := validator.UpdateImageTagVote(req, auth.ID); err != nil {
        return res, err
    }

    isExist, err := model.ImageTagVoteLog{}.IsExist(req.ImageId, req.TagId, auth.ID)
    if err != nil {
        return res, errors.Internal(err.Error())
    }

    if isExist {
        err = model.ImageTagVoteLog{}.Update(req.ImageId, req.TagId, auth.ID, req.Vote)
    } else {
        err = model.ImageTagVoteLog{}.Create(model.ImageTagVoteLog{
            ImageId: req.ImageId,
            TagId:   req.TagId,
            UserId:  auth.ID,
            Vote:    req.Vote,
        })

        if err != nil {
            return res, errors.Internal(err.Error())
        }

        err = model.User{}.UpdateScore(auth.ID, constant.ImageTagVoteUpScore, model.ScoreCategoryAgreeImageTag)
    }
    if err != nil {
        return res, errors.Internal(err.Error())
    }

    go func() {
        err = model.ImageTag{}.SyncVote(req.ImageId, req.TagId)
        if err != nil {
            // log
        }
    }()

    return res, nil

}
