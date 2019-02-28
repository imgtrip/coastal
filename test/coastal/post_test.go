package main

import (
	"coastal/internal/pkg/model"
	"coastal/internal/pkg/pb"
	"github.com/icrowley/fake"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestPosts(t *testing.T) {
	_, err := testServer.Posts(touristCtx(), &pb.PostsReq{})
	assert.Nil(t, err)

	ctx := authCtx()
	post := &model.Post{Title: fake.Title(), UserId: ctx.GetUserId(), Body: fake.EmailSubject(), CommentHostId: ctx.GetAlbum().CommentHostId}
	err = dbConnect.Create(post).Error
	assert.Nil(t, err)

	res, err := testServer.Posts(ctx, &pb.PostsReq{Start: 0, Limit: 10})
	assert.Nil(t, err)
	assert.NotEmpty(t, res.Items)
}

func TestShowPost(t *testing.T) {
	_, err := testServer.ShowPost(touristCtx(), &pb.ShowPostReq{})
	assert.NotNil(t, err)

	ctx := authCtx()
	post := &model.Post{Title: fake.Title(), UserId: ctx.GetUserId(), Body: fake.EmailSubject(), CommentHostId: ctx.GetAlbum().CommentHostId}
	err = dbConnect.Create(post).Error
	assert.Nil(t, err)

	res, err := testServer.ShowPost(touristCtx(), &pb.ShowPostReq{Id: post.ID})
	assert.Nil(t, err)
	assert.Equal(t, res.Post.Views, uint64(1))

	res, err = testServer.ShowPost(touristCtx(), &pb.ShowPostReq{Id: post.ID})
	assert.Nil(t, err)
	assert.Equal(t, res.Post.Views, uint64(2))
}
