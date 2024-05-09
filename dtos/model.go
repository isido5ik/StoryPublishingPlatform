package dtos

import "time"

type Roles struct {
	RoleId   int    `json:"RoleId"`
	RoleName string `json:"RoleName"`
}

type User struct {
	UserID    int       `json:"user_id" db:"user_id"`
	Username  string    `json:"username" db:"username" `
	Email     string    `json:"email" db:"email" `
	Password  string    `json:"password" db:"password_hash" `
	CreatedAt time.Time `json:"created_at" db:"created_at"`
}

type Post struct {
	PostID    int       `json:"post_id" db:"post_id"`
	Author    string    `json:"author" db:"username"`
	UserID    int       `json:"user_id" db:"user_id"`
	Title     string    `json:"title" db:"title"`
	Content   string    `json:"content" db:"content"`
	Comments  int       `json:"comments" db:"comments"`
	Likes     int       `json:"likes" db:"likes"`
	CreatedAt time.Time `json:"created_at" db:"created_at"`
}

type Comment struct {
	CommentID   int       `json:"comment_id" db:"comment_id"`
	UserID      int       `json:"user_id" db:"user_id"`
	PostID      int       `json:"post_id" db:"post_id"`
	ParentID    *int      `json:"parent_id,omitempty" db:"parent_id"`
	CommentText string    `json:"comment_text" db:"comment_text" `
	CreatedAt   time.Time `json:"created_at" db:"created_at"`
}

type Like struct {
	LikeID    int       `json:"like_id" db:"like_id"`
	UserID    int       `json:"user_id" db:"user_id"`
	PostID    int       `json:"post_id" db:"post_id"`
	CreatedAt time.Time `json:"created_at" db:"created_at"`
}

// type Story struct {
// 	StoryId   int       `json:"story_id" db:"post_id"`
// 	UserId    int       `json:"user_id" db:"user_id"`
// 	Author    string    `json:"author" db:"username"`
// 	Content   string    `json:"content" db:"content"`
// 	Comments  int       `json:"comments" db:"comments"`
// 	Likes     int       `json:"likes" db:"likes"`
// 	CreatedAt time.Time `json:"created_at" db:"created_at"`
// }

// type MyStory struct {
// 	StoryId   int       `json:"story_id" db:"post_id"`
// 	Content   string    `json:"content" db:"content"`
// 	Comments  int       `json:"comments" db:"comments"`
// 	Likes     int       `json:"likes" db:"likes"`
// 	CreatedAt time.Time `json:"created_at" db:"created_at"`
// }
