package model

import (
    "coastal/config/constant"
    "coastal/pkg/magic"
    "errors"
    "fmt"
    "github.com/jinzhu/gorm"
)

const (
    AlbumViews    = "views"
    AlbumAmounts  = "amounts"
    AlbumComments = "comments"
)

type Album struct {
    CommonFields
    Title         string
    UserId        uint64
    IsPublic      bool
    Comments      uint64
    Views         uint64
    Amounts       uint64
    Cover         string
    CommentHostId uint64
    Score         uint64
    Description   string
}

func (Album) Paginate(offset uint64, limit uint64, order string) ([]Album, error) {
    var albums []Album
    e := Connect.Where(&Album{IsPublic: true},
    ).Where("amounts>?", 0,
    ).Offset(offset,
    ).Order("score DESC",
    ).Limit(limit,
    ).Find(&albums,
    ).Error
    return albums, e
}

func (Album) ById(id uint64) (Album, error) {
    var album Album
    e := Connect.First(&album, id).Error
    return album, e
}

func (Album) ByCommentHostId(id uint64) (Album, error) {
    var album Album
    e := Connect.Where("comment_host_id=?", id).First(&album).Error
    return album, e
}

func (Album) SimpleCreate(title string, userId uint64) (*Album, error) {
    commentHost, err := CommentHost{}.SimpleCreate()
    if err != nil {
        return &Album{}, err
    }

    album := &Album{
        Title:         title,
        UserId:        userId,
        IsPublic:      true,
        CommentHostId: commentHost.ID,
        Cover:         constant.DefaultAlbumCover,
    }

    e := Connect.Create(album).Error
    return album, e
}

func (Album) Delete(id uint64) error {
    return Connect.Where("id=?", id).Delete(Album{}).Error
}

func (Album) ByUserId(userId uint64) ([]Album, error) {
    var albums []Album
    e := Connect.Where("user_id=?", userId).Order("id DESC").Find(&albums).Error
    return albums, e
}

func (a Album) Increment(id uint64, field string) error {
    if err := a.InDecrementFieldCheck(field); err != nil {
        return err
    }

    return Connect.Model(&Album{}).Where("id=?", id).Update(field, gorm.Expr(field+" + ?", 1)).Error
}

func (a Album) Decrement(id uint64, field string) error {
    if err := a.InDecrementFieldCheck(field); err != nil {
        return err
    }
    return Connect.Model(&Album{}).Where("id=?", id).Update(field, gorm.Expr(field+" - ?", 1)).Error
}

func (Album) InDecrementFieldCheck(field string) error {
    if magic.Arr.HasString(field, []string{AlbumAmounts, AlbumViews, AlbumComments}) {
        return nil
    }
    return errors.New(fmt.Sprintf("illegal increment or decrement field : \"%v\"", field))
}

func (a Album) IsOfficialAlbum(albumId uint64) (bool, error) {
    album, err := a.ById(albumId)
    if err != nil {
        return false, err
    }

    return User{}.IsOfficialUser(album.UserId)
}
