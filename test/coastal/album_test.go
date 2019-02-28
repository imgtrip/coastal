package main

import (
    "coastal/internal/pkg/model"
    "coastal/internal/pkg/pb"
    "coastal/test/coastal/util"
    "github.com/icrowley/fake"
    "github.com/stretchr/testify/assert"
    "testing"
)

func TestAlbums(t *testing.T) {
    _, err := testServer.Albums(touristCtx(), &pb.AlbumsReq{})
    assert.Nil(t, err)
}

func TestCreateAlbum(t *testing.T) {
    title := fake.Title()

    req := &pb.CreateAlbumReq{Title: title}
    _, err := testServer.CreateAlbum(authCtx(), req)
    assert.NotNil(t, err)

    ctx := authCtx()
    req = &pb.CreateAlbumReq{UserId: ctx.GetUserId()}
    _, err = testServer.CreateAlbum(ctx, req)
    assert.NotNil(t, err)

    ctx = authCtx()
    req = &pb.CreateAlbumReq{Title: title, UserId: ctx.GetUserId()}
    _, err = testServer.CreateAlbum(touristCtx(), req)
    assert.NotNil(t, err)

    ctx = authCtx()
    req = &pb.CreateAlbumReq{Title: title, UserId: ctx.GetUserId()}
    _, err = testServer.CreateAlbum(ctx, req)
    assert.Nil(t, err)
}

func TestShowAlbum(t *testing.T) {
    album := &model.Album{}
    err := dbConnect.First(album).Error
    assert.Nil(t, err, err)

    _, err = testServer.ShowAlbum(authCtx(), &pb.ShowAlbumReq{})
    assert.NotNil(t, err, err)

    res, err := testServer.ShowAlbum(authCtx(), &pb.ShowAlbumReq{Id: album.ID})
    assert.Nil(t, err, err)
    assert.Equal(t, album.ID, res.Album.Id)
}

func TestUpdateAlbum(t *testing.T) {
    ctx := authCtx()
    album, err := model.Album{}.SimpleCreate(fake.Title(), ctx.GetUserId())
    assert.Nil(t, err)

    _, err = testServer.UpdateAlbum(ctx, &pb.UpdateAlbumReq{Id: album.ID, Title: fake.Title()})
    assert.Nil(t, err)

    _, err = testServer.UpdateAlbum(ctx, &pb.UpdateAlbumReq{Id: album.ID, Description: fake.Title()})
    assert.Nil(t, err)
}

func TestDeleteAlbum(t *testing.T) {
    ctx := authCtx()
    album, err := model.Album{}.SimpleCreate(fake.Title(), ctx.GetUserId())
    assert.Nil(t, err, err)

    _, err = testServer.DeleteAlbum(ctx, &pb.DeleteAlbumReq{Id: album.ID})
    assert.Nil(t, err)
}

func TestUserAlbums(t *testing.T) {
    ctx := authCtx()
    _, err := model.Album{}.SimpleCreate(fake.Title(), ctx.GetUserId())
    assert.Nil(t, err)

    _, err = testServer.UserAlbums(ctx, &pb.UserAlbumsReq{UserId: ctx.GetUserId() + 1})
    assert.Nil(t, err)

    ctx = authCtx()
    _, err = testServer.UserAlbums(touristCtx(), &pb.UserAlbumsReq{UserId: ctx.GetUserId()})
    assert.Nil(t, err)

    ctx = authCtx()
    _, err = testServer.UserAlbums(ctx, &pb.UserAlbumsReq{UserId: ctx.GetUserId()})
    assert.Nil(t, err)
}

func TestAlbumImages(t *testing.T) {
    ctx := authCtx()
    albumId := ctx.GetUser().AlbumId
    err := dbConnect.Create(&model.AlbumImage{AlbumId: albumId, ImageId: util.ImageId()}).Error
    assert.Nil(t, err)

    _, err = testServer.AlbumImages(touristCtx(), &pb.AlbumImagesReq{AlbumId: 0})
    assert.NotNil(t, err)

    _, err = testServer.AlbumImages(touristCtx(), &pb.AlbumImagesReq{AlbumId: albumId})
    assert.Nil(t, err)
}

func TestCreateAlbumImage(t *testing.T) {
    imageId := util.ImageId()

    ctx := authCtx()
    _, err := testServer.CreateAlbumImage(touristCtx(), &pb.CreateAlbumImageReq{AlbumId: ctx.GetUser().AlbumId, ImageId: imageId})
    assert.NotNil(t, err)

    _, err = testServer.CreateAlbumImage(authCtx(), &pb.CreateAlbumImageReq{})
    assert.NotNil(t, err)

    ctx = authCtx()
    _, err = testServer.CreateAlbumImage(ctx, &pb.CreateAlbumImageReq{AlbumId: ctx.GetUser().AlbumId})
    assert.NotNil(t, err)

    ctx = authCtx()
    _, err = testServer.CreateAlbumImage(ctx, &pb.CreateAlbumImageReq{AlbumId: ctx.GetUser().AlbumId, ImageId: imageId})
    assert.Nil(t, err)
}

func TestDeleteAlbumImage(t *testing.T) {
    ctx := authCtx()
    album, err := model.Album{}.SimpleCreate(fake.Title(), ctx.GetUserId())
    assert.Nil(t, err)
    imageId := util.ImageId()
    _, err = model.AlbumImage{}.Create(album.ID, imageId)
    assert.Nil(t, err)

    _, err = testServer.DeleteAlbumImage(touristCtx(), &pb.DeleteAlbumImageReq{AlbumId: album.ID, ImageId: imageId})
    assert.NotNil(t, err)

    _, err = testServer.DeleteAlbumImage(authCtx(), &pb.DeleteAlbumImageReq{})
    assert.NotNil(t, err)

    ctx = authCtx()
    album, err = model.Album{}.SimpleCreate(fake.Title(), ctx.GetUserId())
    assert.Nil(t, err)
    _, err = testServer.DeleteAlbumImage(ctx, &pb.DeleteAlbumImageReq{AlbumId: album.ID})
    assert.NotNil(t, err)

    album, err = model.Album{}.SimpleCreate(fake.Title(), ctx.GetUserId())
    assert.Nil(t, err)
    imageId = util.ImageId()
    _, err = model.AlbumImage{}.Create(album.ID, imageId)
    assert.Nil(t, err)
    _, err = testServer.DeleteAlbumImage(ctx, &pb.DeleteAlbumImageReq{AlbumId: album.ID, ImageId: imageId})
    assert.Nil(t, err)
}
