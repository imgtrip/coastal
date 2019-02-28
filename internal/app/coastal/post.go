package coastal

import (
	"coastal/internal/app/coastal/response"
	"coastal/internal/app/coastal/validator"
	"coastal/internal/pkg/errors"
	"coastal/internal/pkg/model"
	"coastal/internal/pkg/pb"
	"golang.org/x/net/context"
)

func (*Server) Posts(c context.Context, req *pb.PostsReq) (*pb.PostsRes, error) {
	res := &pb.PostsRes{}
	_, err := GetTokenFromContext(c)
	if err != nil {
		return res, err
	}

	if err := validator.Posts(req); err != nil {
		return res, errors.Internal(err.Error())
	}

	posts, err := PostModel.Paginate(req.Start, req.Limit)
	if err != nil {
		return res, errors.Internal(err.Error())
	}

	return response.Posts(posts), nil
}

func (*Server) ShowPost(c context.Context, req *pb.ShowPostReq) (*pb.ShowPostRes, error) {
	res := &pb.ShowPostRes{}
	auth, err := Auth(c, cachedUser)
	if err != nil {
		return res, err
	}

	if err := validator.ShowPost(req, auth.ID); err != nil {
		return res, errors.Internal(err.Error())
	}

	err = PostModel.Increment(req.Id, model.PostViews)
	if err != nil {
		return res, errors.Internal(err.Error())
	}

	post, err := PostModel.ById(req.Id)
	if err != nil {
		return res, errors.Internal(err.Error())
	}

	return response.ShowPost(post), nil
}
