package coastal

import (
    "coastal/internal/app/coastal/response"
    "coastal/internal/app/coastal/validator"
    "coastal/internal/pkg/errors"
    "coastal/internal/pkg/pb"
    "golang.org/x/net/context"
)

func (*Server) AuthToken(c context.Context, req *pb.AuthTokenReq) (*pb.AuthTokenRes, error) {
    // 检查token时效性,返回token包含的值信息
    // res := EmptyAuthTokenRes
    res := &pb.AuthTokenRes{}
    if err := validator.AuthToken(req); err != nil {
        return res, err
    }

    t, err := TokenModel.ByHash(req.Token)
    if err != nil {
        return res, errors.Internal(err.Error())
    }

    if err = TokenModel.Refresh(t.Hash); err != nil {
        return res, errors.Internal(err.Error())
    }

    if t.UserId == 0 {
        return res, nil
    }

    user, err := UserModel.ById(t.UserId)
    if err != nil {
        return res, errors.Internal(err.Error())
    }

    return response.AuthToken(user), nil
}

func (*Server) CreateToken(c context.Context, req *pb.CreateTokenReq) (*pb.CreateTokenRes, error) {
    empty := &pb.CreateTokenRes{}
    t, err := GetTokenFromContext(c)
    if err != nil {
        return empty, err
    }

    if err := validator.CreateToken(t); err != nil {
        return empty, errors.Internal(err.Error())
    }
    token, err := TokenModel.Create(req.UserId)
    if err != nil {
        return empty, errors.Internal(err.Error())
    }

    return response.CreateToken(token), nil
}

func (*Server) UpdateToken(c context.Context, req *pb.UpdateTokenReq) (*pb.UpdateTokenRes, error) {
    res := &pb.UpdateTokenRes{}
    auth, err := Auth(c, cachedUser)
    if err != nil {
        return res, err
    }

    if err = validator.UpdateToken(req, auth.ID); err != nil {
        return res, errors.Internal(err.Error())
    }

    token, err := GetTokenFromContext(c)
    if err != nil {
        return res, err
    }

    if err := TokenModel.Refresh(token); err != nil {
        return res, errors.Internal(err.Error())
    }

    // 前端更没有传 user_id
    // err = model.DB.Model(model.Token{}).Where("hash=?", token,
    // ).Update("user_id", req.UserId,
    // ).Error
    // if err != nil {
    //     return res, errors.Internal(err.Error())
    // }

    return res, nil
}
