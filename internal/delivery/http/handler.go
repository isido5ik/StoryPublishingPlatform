// package http

// import (
// 	"github.com/gin-gonic/gin"
// 	"github.com/isido5ik/StoryPublishingPlatform/internal/usecase"
// )

// type Handler struct {
// 	useCases usecase.Usecase
// }

// func NewHandler(useCase usecase.Usecase) *Handler {
// 	return &Handler{useCases: useCase}
// }

// func (h *Handler) InitRoutes() *gin.Engine {
// 	router := gin.New()

// 	auth := router.Group("/auth")
// 	{
// 		auth.POST("/sign-up", h.signUp)
// 		auth.POST("/sign-in", h.signIn)
// 	}
// 	story := router.Group("/story")
// 	{

// 		userIdentityMiddleware := h.UserIdentity()
// 		story.Use(userIdentityMiddleware)

// 		adminMiddleware := h.CheckRole(adminCtx)
// 		clientMiddleware := h.CheckRole(clientCtx)

// 		client := story.Group("/client")
// 		{
// 			client.POST("/", clientMiddleware, h.createPost)
// 			client.GET("/", clientMiddleware, h.getPost)
// 			client.DELETE("/:id", clientMiddleware, h.deletePost)
// 		}

// 		admin := story.Group("/admin")
// 		{
// 			admin.DELETE("/:id", adminMiddleware, h.deletePost)
// 		}

// 	}

// 	//other handlers
// 	return router
// }

package http

import (
	"github.com/gin-gonic/gin"
	"github.com/isido5ik/StoryPublishingPlatform/internal/usecase"
)

type Handler struct {
	useCases usecase.Usecase
}

func NewHandler(useCase usecase.Usecase) *Handler {
	return &Handler{useCases: useCase}
}

func (h *Handler) InitRoutes() *gin.Engine {
	router := gin.New()

	api := router.Group("/api")
	{
		auth := api.Group("/auth")
		{
			auth.POST("/sign-up", h.signUp)
			auth.POST("/sign-in", h.signIn)
		}

		stories := api.Group("/stories")
		{
			// Middleware
			userIdentityMiddleware := h.UserIdentity()
			stories.Use(userIdentityMiddleware)

			// Client routes
			stories.GET("/", h.getStories)
			stories.GET("/:story_id", h.getStory)
			stories.POST("/", h.createStory)
			stories.PUT("/:story_id", h.updateStory)
			stories.DELETE("/:story_id", h.deleteStory)

			// Admin routes
		}
		admin := api.Group("/admin")
		{
			admin.Use(h.UserIdentity())
			admin.Use(h.CheckRole(adminCtx))

			admin.GET("/users", h.getUsers)
			admin.GET("/users/:user_id", h.getUser)
			admin.PUT("/users/:user_id", h.updateUser)
			admin.DELETE("/users/:user_id", h.deleteUser)
		}
	}

	// Define other handlers here if needed

	return router
}
