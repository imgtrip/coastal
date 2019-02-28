package coastal

import (
    "coastal/internal/app/coastal/response"
    "coastal/internal/app/coastal/validator"
    "coastal/internal/pkg/errors"
    "coastal/internal/pkg/model"
    "coastal/internal/pkg/pb"
    "golang.org/x/net/context"
)

const OrderBy = "id desc"

func (*Server) Albums(c context.Context, req *pb.AlbumsReq) (*pb.AlbumsRes, error) {
    res := &pb.AlbumsRes{}
    if err := validator.Albums(req); err != nil {
        return res, err
    }

    albums, err := AlbumModel.Paginate(req.Start, req.Limit, OrderBy)
    if err != nil {
        return res, err
    }

    total := 0
    if req.Total {
        err = model.DB.Model(model.Album{}).Where(
            "is_public=?", true,
        ).Where("score>?", 0,
        ).Where("amounts>?", 0).Count(&total,
        ).Error
        if err != nil {
            return res, errors.Internal(err.Error())
        }
    }

    return response.Albums(albums, total), nil
}

func (*Server) CreateAlbum(c context.Context, req *pb.CreateAlbumReq) (*pb.CreateAlbumRes, error) {
    res := &pb.CreateAlbumRes{}
    auth, err := Auth(c, cachedUser)
    if err != nil {
        return res, err
    }

    if err := validator.CreateAlbum(req, auth.ID); err != nil {
        return res, err
    }

    album, err := AlbumModel.SimpleCreate(req.Title, req.UserId)
    if err != nil {
        return res, err
    }

    res.Id = album.ID
    return res, nil
}

func (*Server) ShowAlbum(c context.Context, req *pb.ShowAlbumReq) (*pb.ShowAlbumRes, error) {
    res := &pb.ShowAlbumRes{}
    user, err := Auth(c, cachedUser)
    if err != nil {
        return res, err
    }

    if err := validator.ShowAlbum(req, user.ID); err != nil {
        return res, err
    }

    album, err := AlbumModel.ById(req.Id)
    if err != nil {
        return res, err
    }

    err = AlbumModel.Increment(req.Id, model.AlbumViews)
    if err != nil {
        return res, err
    }

    return response.ShowAlbum(album), nil
}

func (s *Server) UpdateAlbum(c context.Context, req *pb.UpdateAlbumReq) (*pb.UpdateAlbumRes, error) {
    auth, err := Auth(c, cachedUser)
    empty := &pb.UpdateAlbumRes{}
    if err != nil {
        return empty, err
    }

    if err := validator.UpdateAlbum(req, auth.ID); err != nil {
        return empty, err
    }

    if err := s.DBConnect.Model(AlbumModel).Where("id=?", req.Id).Updates(req).Error; err != nil {
        return empty, err
    }

    return empty, nil
}

func (*Server) CreateAlbumImage(c context.Context, req *pb.CreateAlbumImageReq) (*pb.CreateAlbumImageRes, error) {
    empty := &pb.CreateAlbumImageRes{}
    auth, err := Auth(c, cachedUser)
    if err != nil {
        return empty, err
    }

    if err := validator.CreateAlbumImage(req, auth.ID); err != nil {
        return empty, err
    }

    if _, err = AlbumImageModel.Create(req.AlbumId, req.ImageId); err != nil {
        return empty, err
    }

    return empty, nil
}

func (*Server) DeleteAlbumImage(c context.Context, req *pb.DeleteAlbumImageReq) (*pb.DeleteAlbumImageRes, error) {
    empty := &pb.DeleteAlbumImageRes{}
    auth, err := Auth(c, cachedUser)

    if err := validator.DeleteAlbumImage(req, auth.ID); err != nil {
        return empty, err
    }

    if err = AlbumImageModel.Delete(req.AlbumId, req.ImageId); err != nil {
        return empty, err
    }

    return empty, nil
}

func (*Server) DeleteAlbum(c context.Context, req *pb.DeleteAlbumReq) (*pb.DeleteAlbumRes, error) {
    auth, err := Auth(c, cachedUser)
    empty := &pb.DeleteAlbumRes{}
    if err != nil {
        return empty, err
    }

    if err := validator.DeleteAlbum(req, auth.ID); err != nil {
        return empty, err
    }

    if err := AlbumModel.Delete(req.Id); err != nil {
        return empty, err
    }

    return empty, nil
}

func (*Server) UserAlbums(c context.Context, req *pb.UserAlbumsReq) (*pb.UserAlbumsRes, error) {
    auth, err := Auth(c, cachedUser)
    empty := &pb.UserAlbumsRes{}
    if err != nil {
        return empty, err
    }

    if err := validator.UserAlbums(req, auth.ID); err != nil {
        return empty, err
    }
    albums, err := AlbumModel.ByUserId(req.UserId)
    if err != nil {
        return empty, err
    }

    return response.UserAlbums(albums), nil
}

func (*Server) AlbumImages(c context.Context, req *pb.AlbumImagesReq) (*pb.AlbumImagesRes, error) {
    auth, err := Auth(c, cachedUser)
    res := &pb.AlbumImagesRes{}
    if err != nil {
        return res, err
    }

    if err := validator.AlbumImages(req, auth.ID); err != nil {
        return res, err
    }

    images, err := AlbumImageModel.Paginate(req.AlbumId, req.Start, req.Limit, OrderBy)
    if err != nil {
        return res, errors.Internal(err.Error())
    }

    imagesWithTags, err := imageWithRelations(auth, images)
    if err != nil {
        return res, errors.Internal(err.Error())
    }

    total := 0
    if req.Total {
        err = model.DB.Model(model.AlbumImage{}).Where(
            "album_id=?", req.AlbumId,
        ).Count(&total,
        ).Error
        if err != nil {
            return res, errors.Internal(err.Error())
        }
    }

    return response.AlbumImages(imagesWithTags, total), nil
}
