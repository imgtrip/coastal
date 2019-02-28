package main

import (
    "coastal/internal/pkg/model"
    "coastal/internal/pkg/pb"
    "coastal/pkg/magic"
    "coastal/test/coastal/util"
    "github.com/icrowley/fake"
    "github.com/stretchr/testify/assert"
    "testing"
)

func TestCreateZoomLog(t *testing.T) {
    fp := util.Fingerprint()
    imageId := util.ImageId()

    _, err := testServer.CreateZoomLog(touristCtx(), &pb.CreateZoomLogReq{})
    assert.NotNil(t, err)

    _, err = testServer.CreateZoomLog(touristCtx(), &pb.CreateZoomLogReq{Fingerprint: fp})
    assert.NotNil(t, err)

    _, err = testServer.CreateZoomLog(touristCtx(), &pb.CreateZoomLogReq{ImageId: imageId})
    assert.NotNil(t, err)

    _, err = testServer.CreateZoomLog(touristCtx(), &pb.CreateZoomLogReq{Fingerprint: fp, ImageId: imageId})
    assert.Nil(t, err)
}

func TestCreateDownloadLog(t *testing.T) {
    fp := util.Fingerprint()
    imageId := util.ImageId()

    _, err := testServer.CreateDownloadLog(touristCtx(), &pb.CreateDownloadLogReq{})
    assert.NotNil(t, err)

    _, err = testServer.CreateDownloadLog(touristCtx(), &pb.CreateDownloadLogReq{Fingerprint: fp})
    assert.NotNil(t, err)

    _, err = testServer.CreateDownloadLog(touristCtx(), &pb.CreateDownloadLogReq{ImageId: imageId})
    assert.NotNil(t, err)

    _, err = testServer.CreateDownloadLog(touristCtx(), &pb.CreateDownloadLogReq{Fingerprint: fp, ImageId: imageId})
    assert.Nil(t, err)
}

func TestCreateErrorLog(t *testing.T) {

    code := magic.Num.RandInt(1, 10)
    message := fake.Title()
    _, err := tServer.CreateErrorLog(touristCtx(), &pb.CreateErrorLogReq{
        Code:        uint64(code),
        Message:     message,
        Url:         fake.Words(),
        Payload:     fake.Words(),
        Environment: fake.Title(),
        Header:      fake.Words(),
        Cookie:      fake.Words(),
    })
    assert.Nil(t, err)

    count := 0
    err = db.Model(model.ErrorLog{}).Where("code=?", code).Where("message=?", message).Count(&count).Error
    assert.Nil(t, err)
    assert.Equal(t, 1, count)
}
