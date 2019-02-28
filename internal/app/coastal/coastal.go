package coastal

import (
    "coastal/internal/pkg/email"
    "coastal/internal/pkg/model"
    "coastal/internal/runtime"
    "github.com/jinzhu/gorm"
)

var (
    AlbumModel        = model.Album{}
    AlbumImageModel   = model.AlbumImage{}
    ImageSessionModel = model.ImageSession{}
    ImageModel        = model.Image{}
    PostModel         = model.Post{}
    CommentModel      = model.Comment{}
    CommentHostModel  = model.CommentHost{}
    ZoomLogModel      = model.ZoomLog{}
    DownloadLogModel  = model.DownloadLog{}
    TokenModel        = model.Token{}
    UserModel         = model.User{}
)

type Server struct {
    DBConnect *gorm.DB
    Email     *email.Client
}

var server *Server

func New() *Server {
    server = &Server{
        DBConnect: runtime.DBConnect(),
        Email:     runtime.Email(),
    }
    return server
}
