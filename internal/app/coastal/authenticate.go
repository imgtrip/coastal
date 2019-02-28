package coastal

import (
    "coastal/config/constant"
    "coastal/internal/pkg/errors"
    "coastal/internal/pkg/model"
    "golang.org/x/net/context"
    "google.golang.org/grpc/codes"
    "google.golang.org/grpc/metadata"
)

func cachedUser(token string) (model.User, bool) {
    // 从cache获取user

    return model.User{}, false
}

func GetTokenFromContext(c context.Context) (string, error) {
    md, _ := metadata.FromIncomingContext(c)
    if tokens, ok := md[constant.TokenName]; ok {
        token := tokens[0]
        if len(token) == constant.TokenLen {
            return token, nil
        }
    }

    return "", errors.New(codes.PermissionDenied, "token not found")
}

func GetUserFromToken(token string) (model.User, error) {
    tokenData, err := model.Token{}.ByHash(token)
    if err != nil {
        return model.User{}, errors.Internal(err.Error())
    }

    if tokenData.UserId > 0 {
        user, err := model.User{}.ById(tokenData.UserId)
        if err != nil {
            return model.User{}, errors.Internal(err.Error())
        }
        return user, nil
    } else {
        return model.User{}, nil
    }
}

func GetUserFromContext(c context.Context) (model.User, error) {
    token, err := GetTokenFromContext(c)
    if err != nil {
        return model.User{}, err
    }
    return GetUserFromToken(token)
}

// 1. 从当前服务缓存中获取
// 2. 转交authentication服务
func Auth(c context.Context, cache func(token string) (model.User, bool)) (model.User, error) {
    // token, err := GetTokenFromContext(c)
    // if err != nil {
    // 	return model.User{}, err
    // }
    // if user, ok := cache(token); ok {
    // 	return user, nil
    // }

    return GetUserFromContext(c)
}
