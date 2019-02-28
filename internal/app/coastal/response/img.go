package response

import (
    "coastal/config/constant"
    "coastal/internal/app/coastal/types"
    "coastal/internal/pkg/model"
    "coastal/internal/pkg/pb"
)

func albumItems(albums []model.Album) []*pb.Album {
    var items []*pb.Album
    for _, a := range albums {
        items = append(items, album(a))
    }
    return items
}

func imageTags(tagsWithRelation []types.TagWithRelation) []*pb.ImageTag {
    var imageTags []*pb.ImageTag

    for _, tag := range tagsWithRelation {
        imageTags = append(imageTags, &pb.ImageTag{
            Id:      tag.TagID,
            Name:    tag.TagName,
            IsVoted: tag.Voted,
            // FIXME vote_up 0
            VoteUp:     0,
            ImageTagId: tag.ImageTagId,
        })
    }
    return imageTags
}

func image(img types.ImageWithFavorite, tagsWithRelation []types.TagWithRelation) *pb.Image {
    var tags []*pb.ImageTag
    tags = imageTags(tagsWithRelation)

    return &pb.Image{
        Id:        img.ID,
        Name:      img.Name,
        Src:       img.Src,
        Favorited: img.Favorite,
        VoteUp:    img.VoteUp,
        VoteDown:  img.VoteDown,
        Tags:      tags,
    }
}

func imageItems(ImageWithRelations []types.ImageWithRelation) []*pb.Image {
    var items []*pb.Image
    for _, item := range ImageWithRelations {
        items = append(items, image(item.Image, item.Tags))
    }
    return items
}

func postsItems(posts []model.Post) []*pb.Post {
    var items []*pb.Post
    for _, p := range posts {
        items = append(items, post(p, true))
    }
    return items
}

func album(album model.Album) *pb.Album {
    return &pb.Album{
        Id:            album.ID,
        Title:         album.Title,
        UserId:        album.UserId,
        CreatedAt:     album.CreatedAt.Format(constant.DateFormat),
        Comments:      album.Comments,
        Views:         album.Views,
        Cover:         album.Cover,
        Amounts:       album.Amounts,
        CommentHostId: album.CommentHostId,
        Description:   album.Description,
    }
}

func post(p model.Post, shorter bool) *pb.Post {
    body := p.Body

    // if shorter && len(p.Body) > constant.PostDescreptionLen {
    // 	body = p.Body[:constant.PostDescreptionLen]
    // }

    return &pb.Post{
        Id:            p.ID,
        Title:         p.Title,
        Body:          body,
        UserId:        p.UserId,
        CreatedAt:     p.CreatedAt.Format(constant.DateFormat),
        Comments:      p.Comments,
        Views:         p.Views,
        Cover:         p.Cover,
        CommentHostId: p.CommentHostId,
    }
}

func Albums(albums []model.Album, total int) *pb.AlbumsRes {
    return &pb.AlbumsRes{Items: albumItems(albums), Total: uint64(total)}
}

func UserAlbums(albums []model.Album) *pb.UserAlbumsRes {
    return &pb.UserAlbumsRes{Items: albumItems(albums)}
}

func AlbumImages(imageWithTag []types.ImageWithRelation, total int) *pb.AlbumImagesRes {
    return &pb.AlbumImagesRes{Items: imageItems(imageWithTag), Total: uint64(total)}
}

func Images(imagesWithTags []types.ImageWithRelation, total int) *pb.ImagesRes {
    return &pb.ImagesRes{Items: imageItems(imagesWithTags), Total: uint64(total)}
}

func ShowAlbum(a model.Album) *pb.ShowAlbumRes {
    return &pb.ShowAlbumRes{Album: album(a)}
}

func Posts(posts []model.Post) *pb.PostsRes {
    return &pb.PostsRes{Items: postsItems(posts)}
}

func ShowPost(p model.Post) *pb.ShowPostRes {
    return &pb.ShowPostRes{Post: post(p, false)}
}
