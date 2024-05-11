package http

import (
	"github.com/gin-gonic/gin"
	"github.com/isido5ik/StoryPublishingPlatform/internal/usecase"
	"github.com/sirupsen/logrus"

	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"

	_ "github.com/isido5ik/StoryPublishingPlatform/docs"
)

const (
	authorizationHeader = "Authorization"
	userCtx             = "userId"
	adminCtx            = "ADMIN"
	clientCtx           = "CLIENT"
	rolesCtx            = "roles"
)

type Handler struct {
	useCases usecase.Usecase
}

func NewHandler(useCase usecase.Usecase) *Handler {
	return &Handler{useCases: useCase}
}

func (h *Handler) InitRoutes() *gin.Engine {
	logrus.Info("Initializing routes...")

	router := gin.New()

	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	api := router.Group("/api")
	{
		auth := api.Group("/auth")
		{
			auth.POST("/sign-up", h.signUp)
			auth.POST("/sign-in", h.signIn)
		}

		stories := api.Group("/stories")
		{
			userIdentityMiddleware := h.UserIdentity()

			stories.GET("/", h.getStories)
			stories.GET("/filter", h.getStoriesByCategory)

			stories.POST("/", userIdentityMiddleware, h.createStory)
			stories.GET("/my", userIdentityMiddleware, h.getUsersStories)
			stories.GET("/:story_id", userIdentityMiddleware, h.getStory)
			stories.PUT("/:story_id", userIdentityMiddleware, h.updateStory)
			stories.DELETE("/:story_id", userIdentityMiddleware, h.deleteStory)

			like := stories.Group("/:story_id/like", userIdentityMiddleware)
			{
				like.PUT("/", h.like)
				like.DELETE("/", h.removeLike)
			}
			comment := stories.Group("/:story_id/comment", userIdentityMiddleware)
			{
				comment.POST("/", h.addComment)
				comment.POST("/:comment_id", h.replyToComment)
				comment.PUT("/:comment_id", h.updateComment)
				comment.DELETE("/:comment_id", h.deleteComment)
			}
		}

		admin := api.Group("/admin")
		{
			admin.Use(h.UserIdentity())
			admin.Use(h.CheckRole(adminCtx))

			admin.GET("/users", h.getUsers)
			admin.GET("/users/:user_id", h.getUserById)
			admin.DELETE("/users/:user_id", h.deleteUser)
		}
	}

	return router
}
