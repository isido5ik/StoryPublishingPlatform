package dtos

type SignUpInput struct {
	Username string `json:"username" db:"username" binding:"required"`
	Password string `json:"password" db:"password" binding:"required"`
	Email    string `json:"email" db:"email" binding:"required"`
}

type SignInInput struct {
	Username string `json:"username" db:"username" binding:"required"`
	Password string `json:"password" db:"password" binding:"required"`
}

type AddPostInput struct {
	Title   string `json:"title" binding:"required"`
	Content string `json:"content" binding:"required"`
}
type UpdateStoryInput struct {
	Content string `json:"content" db:"content"`
}
type UpdateCommentInput struct {
	CommentText string `json:"comment_text" binding:"required"`
}

type AddCommentInput struct {
	CommentText string `json:"comment_text" db:"comment_text" binding:"required"`
}
