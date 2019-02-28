package model

import (
    "time"
)

type Tag struct {
    ID        uint64 `gorm:"primary_key"`
    CreatedAt time.Time
    UpdatedAt time.Time
    Name      string
}

func (Tag) ById(id uint64) (Tag, error) {
    var tag Tag
    err := DB.Where("id=?", id).First(&tag).Error
    return tag, err
}
