package model

type Comment struct {
    CommonFields
    UserId        uint64
    Content       string
    CommentHostId uint64
}

func (Comment) Paginate(commentHostId uint64, offset uint64, limit uint64) ([]Comment, error) {
    var comments []Comment
    err := Connect.Where("comment_host_id=?", commentHostId).Order("id DESC").Offset(offset).Limit(limit).Find(&comments).Error
    return comments, err
}

func (Comment) ById(id uint64) (Comment, error) {
    var comment Comment
    err := Connect.Where("id=?", id).First(&comment).Error
    return comment, err
}

func (Comment) Update(id uint64, updates Comment) (Comment, error) {
    err := Connect.Model(&Comment{}).Where("id=?", id).Updates(updates).Error
    return updates, err
}

func (Comment) SimpleCreate(CommentHostId uint64, userId uint64, content string) (Comment, error) {
    comment := Comment{
        CommentHostId: CommentHostId,
        UserId:        userId,
        Content:       content,
    }

    if err := Connect.Create(&comment).Error; err != nil {
        return Comment{}, err
    }

    return comment, nil
}
