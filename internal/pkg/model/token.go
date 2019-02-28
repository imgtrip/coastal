package model

import (
	"time"
	"errors"
	"fmt"
	"coastal/pkg/magic"
	"coastal/config/constant"
)

type Token struct {
	CommonFields
	Hash     string
	UserId   uint64
	ExpireAt time.Time
}

func (Token) Create(userId uint64) (Token, error) {
	token := Token{Hash: magic.Str.Random(constant.TokenLen), UserId: userId, ExpireAt: GenerateExpireAt()}

	if err := Connect.Create(&token).Error; err != nil {
		return token, err
	}
	return token, nil
}

func (Token) Bind(hash string, userId uint64) error {
	expireAt := GenerateExpireAt()

	// WARNING when update with struct, GORM will only update those fields that with non blank value
	// For below Update, nothing will be updated as "", 0, false are blank values of their types
	return Connect.Model(&Token{}).Where("hash=?", hash).Updates(Token{UserId: userId, ExpireAt: expireAt}).Error
}

func (Token) ByHash(hash string, args ...bool) (Token, error) {
	withExpired := false
	if len(args) > 0 {
		withExpired = args[0]
	}

	token := Token{}
	var err error
	if withExpired {
		err = Connect.Where("hash=?", hash).First(&token).Error
		if err != nil {
			err = errors.New(fmt.Sprintf("not found token \"%v\"", hash))
		}
	} else {
		err = Connect.Where("hash=? and expire_at>?", hash, time.Now()).First(&token).Error
		if err != nil {
			err = errors.New(fmt.Sprintf("not found token \"%v\" or token expired ,\"%v\" ", hash, err))
		}
	}

	return token, err
}

func (Token) Refresh(hash string) error {
	return Connect.Model(&Token{}).Where("hash=?", hash).Updates(Token{ExpireAt: GenerateExpireAt()}).Error
}

func GenerateExpireAt() time.Time {
	return time.Now().Add(time.Second * constant.TokenExpireSeconds)
}
