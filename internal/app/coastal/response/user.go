package response

import (
    "coastal/config/constant"
    "coastal/internal/pkg/errors"
    "coastal/internal/pkg/model"
    "coastal/internal/pkg/pb"
)

func user(user model.User) *pb.User {
    return &pb.User{
        Id:   user.ID,
        Name: user.Name,
        AlbumId:   user.AlbumId,
        Avatar:    user.Avatar,
        CreatedAt: user.CreatedAt.Format(constant.DateTimeFormat),
        Score:     user.Score,
        DailyFreeDownloadNumber: user.DailyFreeDownloadNumber,
    }
}

func comment(c model.Comment, u model.User) *pb.Comment {
    return &pb.Comment{
        Id:            c.ID,
        CommentHostId: c.CommentHostId,
        CreatedAt:     c.CreatedAt.Format(constant.DateTimeFormat),
        Content:       c.Content,
        User:          user(u),
    }
}

func CreateUser(u model.User) *pb.CreateUserRes {
    return &pb.CreateUserRes{User: user(u)}
}
func AuthUser(u model.User) *pb.AuthUserRes {
    return &pb.AuthUserRes{User: user(u)}
}

func UpdateUser(u model.User) *pb.UpdateUserRes {
    return &pb.UpdateUserRes{User: user(u)}
}

func ShowUser(u model.User) *pb.ShowUserRes {
    return &pb.ShowUserRes{User: user(u)}
}

func CreateToken(token model.Token) *pb.CreateTokenRes {
    return &pb.CreateTokenRes{Token: token.Hash, UserId: token.UserId}
}

func AuthToken(u model.User) *pb.AuthTokenRes {
    return &pb.AuthTokenRes{User: user(u)}
}

func Comments(comments []model.Comment) (*pb.CommentsRes, error) {
    var s []*pb.Comment

    for _, c := range comments {
        u, err := model.User{}.ById(c.UserId)
        if err != nil {
            u = model.User{}
        }
        s = append(s, comment(c, u))
    }

    return &pb.CommentsRes{Items: s}, nil
}

func CreateComment(c model.Comment) (*pb.CreateCommentRes, error) {
    empty := &pb.CreateCommentRes{}
    u, err := model.User{}.ById(c.UserId)
    if err != nil {
        return empty, errors.Internal(err.Error())
    }

    return &pb.CreateCommentRes{Comment: comment(c, u)}, nil
}

func CreateCommentHost(host *model.CommentHost) (*pb.CreateCommentHostRes, error) {
    return &pb.CreateCommentHostRes{Id: host.ID}, nil
}

func CreateFingerprint(fp model.Fingerprint) *pb.CreateFingerprintRes {
    return &pb.CreateFingerprintRes{Hash: fp.Hash, UserId: fp.UserId}
}

func ImageVoteItem(imageVote model.ImageVote) *pb.ImageVote {
    return &pb.ImageVote{
        ImageId: imageVote.ImageId,
        Vote:    int64(imageVote.Vote),
    }
}

func ImageVotes(imageVotes []model.ImageVote) *pb.ImageVotesRes {
    var items []*pb.ImageVote

    for _, v := range imageVotes {
        items = append(items, ImageVoteItem(v))
    }

    return &pb.ImageVotesRes{Items: items}
}

func DownloadImage(status bool) *pb.DownloadImageRes {
    return &pb.DownloadImageRes{Status: status}
}

func ScoreLogs(scoreLogs []model.ScoreLog) *pb.ScoreLogsRes {
    var items []*pb.ScoreLog
    for _, scoreLog := range scoreLogs {
        items = append(items, &pb.ScoreLog{
            User:         user(scoreLog.User),
            Score:        scoreLog.Score,
            Description:  scoreLog.Description,
            CategoryName: scoreLog.ScoreCategory.Name,
            Symbol:       scoreLog.ScoreCategory.Symbol,
            CreatedAt:    scoreLog.CreatedAt.Format(constant.DateTimeFormat),
        })
    }
    return &pb.ScoreLogsRes{Items: items}
}
