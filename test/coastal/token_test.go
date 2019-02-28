package main

import (
	"coastal/config/constant"
	"coastal/internal/pkg/pb"
	"coastal/test/coastal/util"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestAuthToken(t *testing.T) {
	_, err := testServer.AuthToken(touristCtx(), &pb.AuthTokenReq{})
	assert.NotNil(t, err)

	_, err = testServer.AuthToken(touristCtx(), &pb.AuthTokenReq{Token: util.Token()[1:]})
	assert.NotNil(t, err)

	ctx := authCtx()
	_, err = testServer.AuthToken(touristCtx(), &pb.AuthTokenReq{Token: ctx.GetToken()})
	assert.Nil(t, err)
}

func TestCreateToken(t *testing.T) {
	_, err := testServer.CreateToken(touristCtx(), &pb.CreateTokenReq{})
	assert.NotNil(t, err)

	_, err = testServer.CreateToken(authCtx(), &pb.CreateTokenReq{})
	assert.NotNil(t, err)

	token, err := testServer.CreateToken(&util.CertifiedServerContext{}, &pb.CreateTokenReq{})
	assert.Nil(t, err)
	assert.Equal(t, constant.TokenLen, len(token.Token))
	assert.Empty(t, token.UserId)

	userId := util.UserId()
	token, err = testServer.CreateToken(&util.CertifiedServerContext{}, &pb.CreateTokenReq{UserId: userId})
	assert.Nil(t, err)
	assert.Equal(t, constant.TokenLen, len(token.Token))
	assert.Equal(t, userId, token.UserId)
}

//func TestUpdateToken(t *testing.T) {
//	_, err := testServer.UpdateToken(touristCtx(), &pb.UpdateTokenReq{})
//	assert.NotNil(t, err)
//
//	_, err = testServer.UpdateToken(authCtx(), &pb.UpdateTokenReq{})
//	assert.NotNil(t, err)
//
//	_, err = testServer.UpdateToken(authCtx(), &pb.UpdateTokenReq{})
//	assert.Nil(t, err)
//}
