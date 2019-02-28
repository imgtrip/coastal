package validator

import (
    "coastal/config/constant"
    "coastal/internal/pkg/errors"
    "coastal/internal/pkg/model"
    "coastal/internal/pkg/pb"
)

func Albums(req *pb.AlbumsReq) error {

    return nil
}

func Images(req *pb.ImagesReq) error {
    return nil
}

func UpdateImageVote(req *pb.UpdateImageVoteReq, userId uint64) error {
    if err := Structure(req); err != nil {
        return errors.InvalidArgument(err.Error())
    }

    if userId == 0 {
        return errors.PermissionDenied()
    }

    return nil
}

func UpdateImageName(req *pb.UpdateImageNameReq, userId uint64) error {
    if userId == 0 {
        return errors.PermissionDenied()
    }

    if err := Structure(req); err != nil {
        return errors.InvalidArgument(err.Error())
    }

    satisfied, err := model.User{}.IsScoreGreater(userId, constant.UpdateImageNameScoreRequired)
    if err != nil {
        return errors.Internal(err.Error())
    }

    if !satisfied {
        return errors.PermissionDenied()
    }

    return nil
}

func CreateImageTag(req *pb.CreateImageTagReq, userId uint64) error {
    if userId == 0 {
        return errors.PermissionDenied()
    }

    if err := Structure(req); err != nil {
        return errors.InvalidArgument(err.Error())
    }

    satisfied, err := model.User{}.IsScoreGreater(userId, constant.CreateImageTagScoreRequired)
    if err != nil {
        return errors.Internal(err.Error())
    }

    if !satisfied {
        return errors.PermissionDenied()
    }

    return nil
}

func DeleteImageTag(req *pb.DeleteImageTagReq, userId uint64) error {
    if userId == 0 {
        return errors.PermissionDenied()
    }

    if err := Structure(req); err != nil {
        return errors.InvalidArgument(err.Error())
    }

    satisfied, err := model.User{}.IsScoreGreater(userId, constant.DeleteImageTagScoreRequired)
    if err != nil {
        return errors.Internal(err.Error())
    }

    if !satisfied {
        return errors.PermissionDenied()
    }

    return nil
}

func CreateAlbum(req *pb.CreateAlbumReq, userId uint64) error {
    if err := Structure(req); err != nil {
        return errors.InvalidArgument(err.Error())
    }

    if err := MustOwner(req.UserId, userId); err != nil {
        return errors.PermissionDenied()
    }

    return nil
}

func UpdateAlbum(req *pb.UpdateAlbumReq, userId uint64) error {
    if err := Structure(req); err != nil {
        return errors.InvalidArgument(err.Error())
    }

    a, _ := model.Album{}.ById(req.Id)
    if err := MustOwner(a.UserId, userId); err != nil {
        return errors.PermissionDenied()
    }

    return nil
}

func DeleteAlbum(req *pb.DeleteAlbumReq, userId uint64) error {
    if err := Structure(req); err != nil {
        return errors.InvalidArgument(err.Error())
    }

    a, _ := model.Album{}.ById(req.Id)
    if err := MustOwner(a.UserId, userId); err != nil {
        return errors.PermissionDenied()
    }

    return nil
}

func UserAlbums(req *pb.UserAlbumsReq, userId uint64) error {
    if err := Structure(req); err != nil {
        return errors.InvalidArgument(err.Error())
    }
    return nil
}

func AlbumImages(req *pb.AlbumImagesReq, userId uint64) error {
    if err := Structure(req); err != nil {
        return errors.InvalidArgument(err.Error())
    }

    // album 是否公开检查

    return nil
}

func CreateAlbumImage(req *pb.CreateAlbumImageReq, userId uint64) error {
    if err := Structure(req); err != nil {
        return errors.InvalidArgument(err.Error())
    }

    album, err := model.Album{}.ById(req.AlbumId)
    if err != nil {
        return errors.Internal(err.Error())
    }

    if album.UserId != userId {
        return errors.PermissionDenied()
    }

    return nil
}

func DeleteAlbumImage(req *pb.DeleteAlbumImageReq, userId uint64) error {
    if err := Structure(req); err != nil {
        return errors.InvalidArgument(err.Error())
    }

    album, err := model.Album{}.ById(req.AlbumId)
    if err != nil {
        return errors.Internal(err.Error())
    }

    if album.UserId != userId {
        return errors.PermissionDenied()
    }

    return nil
}

func ShowAlbum(req *pb.ShowAlbumReq, userId uint64) error {
    if err := Structure(req); err != nil {
        return errors.InvalidArgument(err.Error())
    }

    album, err := model.Album{}.ById(req.Id)
    if err != nil {
        return errors.Internal(err.Error())
    }

    if !album.IsPublic && album.UserId != userId && userId != 0 {
        return errors.PermissionDenied()
    }

    return nil
}

func Posts(req *pb.PostsReq) error {
    if err := Structure(req); err != nil {
        return errors.InvalidArgument(err.Error())
    }

    return nil
}

func ShowPost(req *pb.ShowPostReq, userId uint64) error {
    if err := Structure(req); err != nil {
        return errors.InvalidArgument(err.Error())
    }

    // 验证访问权限

    return nil
}

func DownloadImage(req *pb.DownloadImageReq, userId uint64) error {
    if err := Structure(req); err != nil {
        return errors.InvalidArgument(err.Error())
    }

    if userId == 0 {
        return errors.PermissionDenied()
    }
    return nil
}

func UpdateImageTagVote(req *pb.UpdateImageTagVoteReq, userId uint64) error {
    if err := Structure(req); err != nil {
        return errors.InvalidArgument(err.Error())
    }

    if userId == 0 {
        return errors.PermissionDenied()
    }

    return nil
}
