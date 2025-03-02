package post

import "goBlog/internal/domain/user"

type Post struct {
	Id       int64      `json:"id"`
	Title    string     `json:"title"`
	Content  string     `json:"content"`
	AuthorID int64      `json:"author_id"`
	Author   *user.User `json:"author" gorm:"foreignKey:AuthorID;references:id"`
}

func NewPost(title, content string, author *user.User) *Post {
	return &Post{Title: title, Content: content, Author: author, AuthorID: author.Id}
}
