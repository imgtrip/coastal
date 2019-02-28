package util

import (
	"coastal/config/constant"
	"coastal/pkg/magic"
	"coastal/internal/pkg/model"
	"log"
)

func ImageId() uint64 {
	return uint64(magic.Num.RandInt(1, 9999))
}

func Fingerprint() string {
	return magic.Str.Random(constant.FingerprintLen)
}

func Token() string {
	return magic.Str.Random(constant.TokenLen)
}

func UserId() uint64 {
	return uint64(magic.Num.RandInt(1, 9999))
}

func VerificationCode() string {
	return magic.Num.IntToString(magic.Num.RandInt(100000, 999999))
}

func ValidVerificationCode(email string, token string) string {
	code := VerificationCode()
	_, err := model.VerificationCode{}.SimpleCreate(email, code, token)
	if err != nil {
		log.Fatal(err)
	}
	return code
}
