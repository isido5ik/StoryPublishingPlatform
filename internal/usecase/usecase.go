package usecase

import (
	"github.com/isido5ik/StoryPublishingPlatform/dtos"
	"github.com/isido5ik/StoryPublishingPlatform/internal/repository"
)

//go:generate mockgen -source=usecase.go -destination=mocks/mock.go
type Usecase interface {
	CreateUserAsClient(input dtos.SignUpInput) (int, error)
	GenerateToken(username, password string) (string, []dtos.Roles, error)
	ParseToken(token string) (int, []dtos.Roles, error)

	CreateStory(story dtos.AddPostInput, userId int) (int, error)
	GetStories(pagination dtos.PaginationParams) ([]dtos.Post, error)
	GetUsersStories(userId int) (string, []dtos.Post, error)
	GetStoriesByCategory(pagination dtos.PaginationParams, categoryId int) ([]dtos.Post, error)
	GetStory(postId int) (dtos.Post, error)
	DeleteStory(postId, userId int, role string) error
	UpdateStory(postId, userId int, role string, input dtos.UpdateStoryInput) error

	Like(userId, postId int) (int, error)
	RemoveLike(userId, postId int) error
	AddComment(userId, postId int, comment dtos.AddCommentInput) (int, error)
	ReplyToComment(userId, postId, parentId int, comment dtos.AddCommentInput) (int, error)
	UpdateComment(userId, postId, commentId int, newComment dtos.UpdateCommentInput) error
	DeleteComment(userId, postId, commentId int) error

	GetUsers() ([]dtos.User, error)
	GetUserById(userId int) (dtos.User, error)
	DeleteUser(userId int) error
}

type usecase struct {
	repos repository.Repository
}

func NewUsecase(repos repository.Repository) Usecase {
	return &usecase{
		repos: repos,
	}
}
