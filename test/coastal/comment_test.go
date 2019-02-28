package main

import (
    "coastal/internal/pkg/pb"
    "github.com/icrowley/fake"
    "github.com/stretchr/testify/assert"
    "testing"
)

func TestComments(t *testing.T) {
    _, err := testServer.Comments(touristCtx(), &pb.CommentsReq{})
    assert.NotNil(t, err)
}

func TestCreateComment(t *testing.T) {
    ctx := authCtx()
    _, err := testServer.CreateComment(touristCtx(), &pb.CreateCommentReq{CommentHostId: ctx.GetAlbum().CommentHostId, Content: fake.EmailSubject()})
    assert.NotNil(t, err)

    ctx = authCtx()
    _, err = testServer.CreateComment(ctx, &pb.CreateCommentReq{CommentHostId: ctx.GetAlbum().CommentHostId})
    assert.NotNil(t, err)

    ctx = authCtx()
    _, err = testServer.CreateComment(ctx, &pb.CreateCommentReq{Content: fake.EmailSubject()})
    assert.NotNil(t, err)

    // 测试增加views
    // originAlbum, err := model.Album{}.ById(ctx.GetAlbum().ID)
    // assert.Nil(t, err)

    ctx = authCtx()
    _, err = testServer.CreateComment(ctx, &pb.CreateCommentReq{CommentHostId: ctx.GetAlbum().CommentHostId, Content: fake.EmailSubject()})
    assert.Nil(t, err)

    // album, err := model.Album{}.ById(ctx.GetAlbum().ID)
    // assert.Nil(t, err)
    // assert.Equal(t, originAlbum.Comments+1, album.Comments)
}
