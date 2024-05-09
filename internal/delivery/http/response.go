package http

import (
	"github.com/gin-gonic/gin"
	"github.com/isido5ik/StoryPublishingPlatform/dtos"
	"github.com/sirupsen/logrus"
)

type errorResponse struct {
	Message string `json:"message"`
}

type getStoriesResponse struct {
	Data []dtos.Post `json:"data"`
}

type getUsersResponse struct {
	Data []dtos.User `json:"data"`
}
type getMyStoriesResponse struct {
	Username  string      `json:"username"`
	MyStories []dtos.Post `json:"stories"`
}

func newErrorResponse(c *gin.Context, statusCode int, message string) {
	logrus.Error(message)
	c.AbortWithStatusJSON(statusCode, errorResponse{
		Message: message,
	})
}
