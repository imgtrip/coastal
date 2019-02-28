package model

import (
    "coastal/config/constant"
    "time"
)

type VerificationCode struct {
    CommonFields
    Email string
    Code  string
    Token string
}

func (VerificationCode) TableName() string {
    return "verification_codes"
}

func (v VerificationCode) IsAllowToken(token string) (bool, error) {
    count := 0

    createdAt := v.GenerateExpireAt(constant.VerificationCodeTokenExpireSeconds)
    // 该token创建时间>2分钟之前,即在2分钟之内有数据则属于非法请求
    if e := Connect.Model(&VerificationCode{}).Where("token=? AND created_at>?", token, createdAt).Count(&count).Error; e != nil {
        return false, e
    }

    return count == 0, nil
}

func (VerificationCode) GenerateExpireAt(seconds int) time.Time {
    return time.Now().Add(time.Duration(-seconds) * time.Second)
}

func (VerificationCode) SimpleCreate(email string, code string, token string) (*VerificationCode, error) {
    data := &VerificationCode{Email: email, Code: code, Token: token}
    err := Connect.Create(data).Error
    return data, err
}

func (v VerificationCode) IsValidCode(email string, code string) (bool, error) {
    count := 0
    createdAt := v.GenerateExpireAt(constant.VerificationCodeExpireSeconds)
    // 该code创建时间在config时间分钟之内,则有效
    err := Connect.Model(&VerificationCode{}).Where("email=? AND code=? AND created_at>?", email, code, createdAt).Count(&count).Error
    return count > 0, err
}
