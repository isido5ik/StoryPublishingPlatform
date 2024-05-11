package repository

import (
	"github.com/isido5ik/StoryPublishingPlatform/dtos"
	"github.com/jmoiron/sqlx"
)

type Repository interface {
	CreateUserAsClient(input dtos.SignUpInput) (int, error)
	GetUser(username, password string) (dtos.User, error)
	GetRoles(userId int) ([]string, error)
	GetRoleId(role string, userId int) (int, error)

	CreateStory(story dtos.AddPostInput, userId int) (int, error)
	GetStories(pagination dtos.PaginationParams) ([]dtos.Post, error)
	GetUsersStories(userId int) (string, []dtos.Post, error)
	GetStoriesByCategory(pagination dtos.PaginationParams, categoryId int) ([]dtos.Post, error)
	GetStory(postId int) (dtos.Post, error)
	DeleteStory(postId int) error
	UpdateStory(postId int, input dtos.UpdateStoryInput) error

	Like(userId, postId int) (int, error)
	CheckLike(userId, postId int) error
	RemoveLike(userId, postId int) error
	AddComment(userId, postId int, comment dtos.AddCommentInput) (int, error)
	ReplyToComment(userId, postId, parentId int, comment dtos.AddCommentInput) (int, error)
	CheckComment(userId, postId, commentId int) error
	UpdateComment(userId, postId, commentId int, newComment dtos.UpdateCommentInput) error
	DeleteComment(userId, postId, commentId int) error

	GetUsers() ([]dtos.User, error)
	GetUserById(userId int) (dtos.User, error)
	DeleteUser(userId int) error
}

type repository struct {
	db *sqlx.DB
}

func NewRepository(db *sqlx.DB) Repository {
	return &repository{
		db: db,
	}
}
