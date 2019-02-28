package model

import "github.com/jinzhu/gorm"

type Post struct {
	CommonFields
	Title         string
	UserId        uint64
	Comments      uint64
	Views         uint64
	Body          string
	Cover         string
	CommentHostId uint64
}

const (
	PostViews = "views"
)

func (Post) Paginate(offset uint64, limit uint64) ([]Post, error) {
	var posts []Post
	e := Connect.Offset(offset).Limit(limit).Order("updated_at desc").Find(&posts).Error
	return posts, e
}

func (Post) ById(id uint64) (Post, error) {
	var post = Post{}
	err := Connect.First(&post, id).Error
	return post, err
}

func (p Post) Increment(id uint64, field string) error {
	return Connect.Model(&Post{}).Where("id=?", id).Update(field, gorm.Expr(field+" + ?", 1)).Error
}

