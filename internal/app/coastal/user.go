package coastal

import (
    "coastal/config/constant"
    "coastal/internal/app/coastal/response"
    "coastal/internal/app/coastal/validator"
    "coastal/internal/pkg/errors"
    "coastal/internal/pkg/model"
    "coastal/internal/pkg/pb"
    "coastal/pkg/hash"
    "github.com/dgrijalva/jwt-go"
    "github.com/jinzhu/gorm"
    "golang.org/x/net/context"
    "google.golang.org/grpc/codes"
)

func (s *Server) CreateUser(c context.Context, req *pb.CreateUserReq) (*pb.CreateUserRes, error) {
    empty := &pb.CreateUserRes{}
    token, err := GetTokenFromContext(c)
    if err != nil {
        return empty, err
    }

    if err := validator.CreateUser(req); err != nil {
        return empty, err
    }

    user, err := UserModel.Create(req.Name, req.Email, req.Password)
    if err != nil {
        return empty, errors.Internal(err.Error())
    }

    if req.ReferrerEmail != "" {
        err = UserModel.UpdateReferer(user.ID, req.ReferrerEmail)
        // FIXME create referer success message
        if err != nil {
            return empty, errors.Internal(err.Error())
        }
    }

    err = TokenModel.Bind(token, user.ID)
    if err != nil {
        return empty, errors.Internal(err.Error())
    }

    album, err := AlbumModel.SimpleCreate(req.Name+"喜欢的图", user.ID)
    if err != nil {
        return empty, errors.Internal(err.Error())
    }

    err = s.DBConnect.Model(&model.User{}).Where("id=?", user.ID).Update("album_id", album.ID).Error
    if err != nil {
        return empty, errors.Internal(err.Error())
    }
    user.AlbumId = album.ID

    return response.CreateUser(user), nil
}

func (s *Server) UpdateUser(c context.Context, req *pb.UpdateUserReq) (*pb.UpdateUserRes, error) {
    empty := &pb.UpdateUserRes{}
    auth, err := Auth(c, cachedUser)
    if err != nil {
        return empty, err
    }

    if err := validator.UpdateUser(req, auth.ID); err != nil {
        return empty, err
    }

    if err := s.DBConnect.Model(&UserModel).Where("id=?", req.Id).Updates(req).Error; err != nil {
        return empty, errors.Internal(err.Error())
    }

    user, err := UserModel.ById(req.Id)
    if err != nil {
        return empty, errors.Internal(err.Error())
    }

    return response.UpdateUser(user), nil
}

func (*Server) AuthUser(c context.Context, req *pb.AuthUserReq) (*pb.AuthUserRes, error) {
    empty := &pb.AuthUserRes{}
    token, err := GetTokenFromContext(c)
    if err != nil {
        return empty, err
    }

    if err := validator.AuthUser(req); err != nil {
        return empty, err
    }

    user, err := UserModel.ByEmail(req.Email)
    if !hash.Compare(req.Password, user.Password) {
        return empty, errors.New(codes.Unauthenticated, "{\"password\":\"password not match\"}")
    }

    err = TokenModel.Bind(token, user.ID)
    if err != nil {
        return empty, errors.Internal(err.Error())
    }

    return response.AuthUser(user), nil
}

func (*Server) ShowUser(c context.Context, req *pb.ShowUserReq) (*pb.ShowUserRes, error) {
    res := &pb.ShowUserRes{}
    _, err := GetTokenFromContext(c)
    if err != nil {
        return res, err
    }

    if err = validator.ShowUser(req, 0); err != nil {
        return res, err
    }

    user, err := UserModel.ById(req.Id)
    if err != nil {
        if err == gorm.ErrRecordNotFound{
            return res, errors.NotFound(err.Error())
        }
        return res, errors.Internal(err.Error())
    }

    return response.ShowUser(user), nil
}

func (*Server) CreateJwt(c context.Context, req *pb.CreateJwtReq) (*pb.CreateJwtRes, error) {
    res := &pb.CreateJwtRes{}

    _, err := GetTokenFromContext(c)
    if err != nil {
        return res, err
    }

    if err = validator.CreateJwt(req, 0); err != nil {
        return res, err
    }

    token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
        "id": req.Id,
    })

    tokenString, err := token.SignedString([]byte(constant.JwtAuthKey))

    if err != nil {
        return res, errors.Internal(err.Error())
    }
    res.Hash = tokenString
    return res, nil
}

func (s *Server) CreateVerificationEmail(c context.Context, req *pb.CreateVerificationEmailReq) (*pb.CreateVerificationEmailRes, error) {
    res := &pb.CreateVerificationEmailRes{}
    token, err := GetTokenFromContext(c)
    if err != nil {
        return res, err
    }

    if err := validator.CreateVerificationEmail(req, token); err != nil {
        return res, err
    }

    if err := sendVerificationEmail(req.Email, token); err != nil {
        return res, errors.Internal(err.Error())
    }

    return res, nil
}

func (s *Server) UpdatePassword(c context.Context, req *pb.UpdatePasswordReq) (*pb.UpdatePasswordRes, error) {
    res := &pb.UpdatePasswordRes{}
    _, err := GetTokenFromContext(c)
    if err != nil {
        return res, err
    }

    if err := validator.UpdatePassword(req); err != nil {
        return res, err
    }

    hashedPassword, err := hash.In(req.Password)
    if err != nil {
        return res, errors.Internal(err.Error())
    }

    err = s.DBConnect.Model(&model.User{}).Where("email=?", req.Email).Update("password", hashedPassword).Error
    if err != nil {
        return res, errors.Internal(err.Error())
    }

    return res, nil
}

func (s *Server) ImageVotes(c context.Context, req *pb.ImageVotesReq) (*pb.ImageVotesRes, error) {
    res := &pb.ImageVotesRes{}
    auth, err := Auth(c, cachedUser)
    if err != nil {
        return res, err
    }

    if err := validator.ImageVotes(req, auth.ID); err != nil {
        return res, err
    }

    imageVotes, err := model.ImageVote{}.ByUserId(auth.ID)
    if err != nil {
        return res, errors.Internal(err.Error())
    }

    return response.ImageVotes(imageVotes), nil
}

func (s *Server) ScoreLogs(c context.Context, req *pb.ScoreLogsReq) (*pb.ScoreLogsRes, error) {
    res := &pb.ScoreLogsRes{}
    auth, err := Auth(c, cachedUser)
    if err != nil {
        return res, err
    }
    if err := validator.ScoreLogs(req, auth.ID); err != nil {
        return res, err
    }

    order := "created_at DESC"
    if req.Order == pb.ScoreLogOrders_OLD {
        order = "created_at ASC"
    }

    var scoreLogs []model.ScoreLog
    err = model.DB.Preload("User",
    ).Preload("ScoreCategory",
    ).Order(order,
    ).Find(&scoreLogs).Error

    if err != nil && err != gorm.ErrRecordNotFound {
        return res, errors.Internal(err.Error())
    }

    return response.ScoreLogs(scoreLogs), nil
}
