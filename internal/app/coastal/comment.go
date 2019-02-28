package coastal

import (
	"coastal/internal/app/coastal/response"
	"coastal/internal/app/coastal/validator"
	"coastal/internal/pkg/model"
	"coastal/internal/pkg/pb"
	"golang.org/x/net/context"
)

func (*Server) Comments(c context.Context, req *pb.CommentsReq) (*pb.CommentsRes, error) {
	res := &pb.CommentsRes{}
	auth, err := Auth(c, cachedUser)
	if err != nil {
		return res, err
	}

	if err := validator.Comments(req, auth.ID); err != nil {
		return res, err
	}
	comments, err := CommentModel.Paginate(req.CommentHostId, req.Start, req.Limit)
	if err != nil {
		return res, err
	}

	return response.Comments(comments)
}

func (*Server) CreateComment(c context.Context, req *pb.CreateCommentReq) (*pb.CreateCommentRes, error) {
	res := &pb.CreateCommentRes{}
	auth, err := Auth(c, cachedUser)
	if err != nil {
		return res, err
	}

	if err := validator.CreateComment(req, auth.ID); err != nil {
		return res, err
	}
	comment, err := model.Comment{}.SimpleCreate(req.CommentHostId, auth.ID, req.Content)
	if err != nil {
		return res, nil
	}

	return response.CreateComment(comment)
}

func (*Server) UpdateComment(c context.Context, req *pb.UpdateCommentReq) (*pb.UpdateCommentRes, error) {
	empty := &pb.UpdateCommentRes{}
	auth, err := Auth(c, cachedUser)
	if err != nil {
		return empty, err
	}

	if err := validator.UpdateComment(req, auth.ID); err != nil {
		return empty, err
	}

	_, err = CommentModel.Update(req.Id, model.Comment{Content: req.Content})
	if err != nil {
		return empty, err
	}

	return &pb.UpdateCommentRes{}, nil
}

func (*Server) CreateCommentHost(c context.Context, req *pb.CreateCommentHostReq) (*pb.CreateCommentHostRes, error) {
	res := &pb.CreateCommentHostRes{}

	//token, err := GetTokenFromContext(c)
	//if err != nil {
	//	return res, err
	//}

	//if err := validator.CreateCommentHost(token); err != nil {
	//	return res, err
	//}

	host, err := CommentHostModel.SimpleCreate()
	if err != nil {
		return res, err
	}

	return response.CreateCommentHost(host)
}
