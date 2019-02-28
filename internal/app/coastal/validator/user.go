package validator

import (
    "coastal/internal/env"
    "coastal/internal/pkg/errors"
    "coastal/internal/pkg/model"
    "coastal/internal/pkg/pb"
    "google.golang.org/grpc/codes"
)

func CreateUser(req *pb.CreateUserReq) error {
    if err := Structure(req); err != nil {
        return errors.InvalidArgument(err.Error())
    }

    nameExisted, err := model.User{}.IsNameExisted(req.Name)
    if err != nil {
        return errors.Internal(err.Error())
    }

    if nameExisted {
        return errors.New(codes.AlreadyExists, "{\"name\":\"name already exist\"}")
    }

    emailExisted, err := model.User{}.IsEmailExisted(req.Email)
    if err != nil {
        return errors.Internal(err.Error())
    }

    if emailExisted {
        return errors.New(codes.AlreadyExists, "{\"email\":\"email already exist\"}")
    }

    isValidCode, err := model.VerificationCode{}.IsValidCode(req.Email, req.VerificationCode)
    if err != nil {
        return errors.Internal(err.Error())
    }

    if !isValidCode {
        return errors.PermissionDenied("verification code is invalid")
    }

    if req.ReferrerEmail != "" {
        refererEmailExisted, err := model.User{}.IsEmailExisted(req.ReferrerEmail)
        if err != nil {
            return errors.Internal(err.Error())
        }

        if !refererEmailExisted {
            return errors.New(codes.AlreadyExists, "{\"referer_email\":\"not found referer email\"}")
        }
    }

    return nil
}

func UpdateUser(req *pb.UpdateUserReq, userId uint64) error {
    if err := Structure(req); err != nil {
        return errors.InvalidArgument(err.Error())
    }

    if err := MustOwner(req.Id, userId); err != nil {
        return errors.PermissionDenied()
    }

    return nil
}

func ShowUser(req *pb.ShowUserReq, userId uint64) error {
    if req.Id == 0 {
        return errors.InvalidArgument("id is required")
    }
    return nil
}

func AuthUser(req *pb.AuthUserReq) error {
    if err := Structure(req); err != nil {
        return errors.InvalidArgument(err.Error())
    }

    return nil
}

func AuthToken(req *pb.AuthTokenReq) error {
    if err := Structure(req); err != nil {
        return errors.InvalidArgument(err.Error())
    }

    return nil
}

func CreateToken(token string) error {
    if token != env.Process.NodeServerToken {
        return errors.PermissionDenied()
    }

    return nil
}

func UpdateToken(req *pb.UpdateTokenReq, userId uint64) error {
    if err := Structure(req); err != nil {
        return errors.InvalidArgument(err.Error())
    }

    return nil
}

func Comments(req *pb.CommentsReq, userId uint64) error {
    if err := Structure(req); err != nil {
        return errors.InvalidArgument(err.Error())
    }

    // 验证album是否公开?

    return nil
}

func CreateComment(req *pb.CreateCommentReq, userId uint64) error {
    if err := Structure(req); err != nil {
        return errors.InvalidArgument(err.Error())
    }

    if userId == 0 {
        return errors.PermissionDenied()
    }

    return nil
}

func UpdateComment(req *pb.UpdateCommentReq, userId uint64) error {
    if err := Structure(req); err != nil {
        return errors.InvalidArgument(err.Error())
    }

    c, _ := model.Comment{}.ById(req.Id)
    if err := MustOwner(c.UserId, userId); err != nil {
        return errors.PermissionDenied(err.Error())
    }

    return nil
}

func CreateFingerprint(req *pb.CreateFingerprintReq, userId uint64) error {
    if err := Structure(req); err != nil {
        return errors.InvalidArgument(err.Error())
    }
    // 无需验证userId

    return nil
}

func UpdateFingerprint(req *pb.UpdateFingerprintReq, userId uint64) error {
    if err := Structure(req); err != nil {
        return errors.InvalidArgument(err.Error())
    }

    if err := MustOwner(req.UserId, userId); err != nil {
        return errors.PermissionDenied(err.Error())
    }
    return nil
}

func CreateJwt(req *pb.CreateJwtReq, userId uint64) error {
    if req.Id == 0 {
        return errors.InvalidArgument("id is required")
    }
    return nil
}

func CreateVerificationEmail(req *pb.CreateVerificationEmailReq, token string) error {
    if err := Structure(req); err != nil {
        return errors.InvalidArgument(err.Error())
    }

    allow, err := model.VerificationCode{}.IsAllowToken(token)
    if err != nil {
        return errors.Internal(err.Error())
    }

    if !allow {
        return errors.PermissionDenied()
    }

    return nil
}

func UpdatePassword(req *pb.UpdatePasswordReq) error {
    if err := Structure(req); err != nil {
        return errors.InvalidArgument(err.Error())
    }

    valid, err := model.VerificationCode{}.IsValidCode(req.Email, req.VerificationCode)
    if err != nil {
        return errors.Internal(err.Error())
    }

    if !valid {
        return errors.PermissionDenied()
    }

    return nil
}

func ImageVotes(req *pb.ImageVotesReq, userId uint64) error {
    if err := Structure(req); err != nil {
        return errors.InvalidArgument(err.Error())
    }
    if userId == 0 {
        return errors.PermissionDenied()
    }
    return nil
}

func ScoreLogs(req *pb.ScoreLogsReq, userId uint64) error {
    if err := Structure(req); err != nil {
        return errors.InvalidArgument(err.Error())
    }
    if userId == 0 {
        return errors.PermissionDenied()
    }
    return nil
}
