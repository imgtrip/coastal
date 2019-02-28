package main

import (
	"coastal/internal/pkg/pb"
	"coastal/test/coastal/util"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestCreateFingerprint(t *testing.T) {
	fp, err := testServer.CreateFingerprint(touristCtx(), &pb.CreateFingerprintReq{Hash: ""})
	assert.NotNil(t, err, err)

	fp, err = testServer.CreateFingerprint(authCtx(), &pb.CreateFingerprintReq{Hash: ""})
	assert.NotNil(t, err, err)

	hash := util.Fingerprint()
	fp, err = testServer.CreateFingerprint(touristCtx(), &pb.CreateFingerprintReq{Hash: hash})
	assert.Nil(t, err, err)
	assert.Empty(t, fp.UserId)
	assert.Equal(t, hash, fp.Hash)

	hash = util.Fingerprint()
	ctx := authCtx()
	fp, err = testServer.CreateFingerprint(ctx, &pb.CreateFingerprintReq{Hash: hash})
	assert.Nil(t, err, err)
	assert.Equal(t, ctx.GetUserId(), fp.UserId)
	assert.Equal(t, hash, fp.Hash)

	hash = util.Fingerprint()[1:]
	fp, err = testServer.CreateFingerprint(authCtx(), &pb.CreateFingerprintReq{Hash: hash})
	assert.NotNil(t, err, err)
}

func TestUpdateFingerprint(t *testing.T) {
	_, err := testServer.UpdateFingerprint(touristCtx(), &pb.UpdateFingerprintReq{Hash: ""})
	assert.NotNil(t, err, err)

	_, err = testServer.UpdateFingerprint(authCtx(), &pb.UpdateFingerprintReq{Hash: ""})
	assert.NotNil(t, err, err)

	_, err = testServer.UpdateFingerprint(touristCtx(), &pb.UpdateFingerprintReq{Hash: util.Fingerprint()})
	assert.NotNil(t, err, err)

	_, err = testServer.UpdateFingerprint(touristCtx(), &pb.UpdateFingerprintReq{Hash: util.Fingerprint(), UserId: 1})
	assert.NotNil(t, err, err)

	ctx := authCtx()
	_, err = testServer.UpdateFingerprint(ctx, &pb.UpdateFingerprintReq{Hash: util.Fingerprint(), UserId: ctx.GetUserId() + 1})
	assert.NotNil(t, err, err)

	ctx = authCtx()
	_, err = testServer.UpdateFingerprint(ctx, &pb.UpdateFingerprintReq{Hash: util.Fingerprint(), UserId: ctx.GetUserId()})
	assert.Nil(t, err)
}
