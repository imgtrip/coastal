package types

import (
    "coastal/internal/pkg/model"
)

type TagWithRelation struct {
    TagID      uint64
    TagName    string
    ImageTagId uint64
    Voted      bool
}

type ImageWithRelation struct {
    Image ImageWithFavorite
    Tags  []TagWithRelation
}

type ImageWithFavorite struct {
    model.Image
    Favorite bool
}

type ReviewWithImageRelation struct {
    model.Review
    ImageWithRelation ImageWithRelation
    Content string
}
