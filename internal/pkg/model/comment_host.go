package model

type CommentHost struct {
	CommonFields
}

func (CommentHost) TableName() string {
	return "comment_hosts"
}

func (CommentHost) SimpleCreate() (*CommentHost, error) {
	host := &CommentHost{}
	err := Connect.Create(host).Error
	return host, err
}
