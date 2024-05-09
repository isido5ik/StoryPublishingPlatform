package http

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/isido5ik/StoryPublishingPlatform/dtos"
)

// @Summary Create Story
// @Security ApiKeyAuth
// @Tags stories
// @Description create story
// @ID create-story
// @Accept  json
// @Produce  json
// @Param input body dtos.AddPostInput true "recipe info"
// @Success 200 {integer} integer 1
// @Failure 400,404 {object} errorResponse
// @Failure 500 {object} errorResponse
// @Failure default {object} errorResponse
// @Router /api/stories [post]
func (h *Handler) createStory(c *gin.Context) {
	var story dtos.AddPostInput
	if err := c.BindJSON(&story); err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	userId, err := getUserId(c)
	if err != nil {
		newErrorResponse(c, http.StatusForbidden, err.Error())
		return
	}

	postId, err := h.useCases.CreateStory(story, userId)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}
	c.JSON(http.StatusOK, map[string]interface{}{
		"postId": postId,
	})
}

// @Summary Get All Stories
// @Tags stories
// @Description get all stories
// @ID get-all-stories
// @Accept  json
// @Produce  json
// @Success 200 {object} getStoriesResponse
// @Failure 400,404 {object} errorResponse
// @Failure 500 {object} errorResponse
// @Failure default {object} errorResponse
// @Router /api/stories [get]
func (h *Handler) getStories(c *gin.Context) {

	var pagination dtos.PaginationParams
	var err error
	pagination.Page, pagination.PageSize, err = dtos.ValidatePage(c)
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	stories, err := h.useCases.GetStories(pagination)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}
	c.JSON(http.StatusOK, getStoriesResponse{
		Data: stories,
	})
}

// @Summary Get My Story
// @Security ApiKeyAuth
// @Tags stories
// @Description get users stories
// @ID get-my-story
// @Accept  json
// @Produce  json
// @Success 200 {object} getMyStoriesResponse
// @Failure 400,404 {object} errorResponse
// @Failure 500 {object} errorResponse
// @Failure default {object} errorResponse
// @Router /api/stories/my [get]
func (h *Handler) getUsersStories(c *gin.Context) {
	userId, err := getUserId(c)
	if err != nil {
		newErrorResponse(c, http.StatusForbidden, err.Error())
		return
	}
	username, stories, err := h.useCases.GetUsersStories(userId)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}
	c.JSON(http.StatusOK, getMyStoriesResponse{
		Username:  username,
		MyStories: stories,
	})
}

// @Summary Get Story By Id
// @Security ApiKeyAuth
// @Tags stories
// @Description get story by id
// @ID get-story-by-id
// @Param :story_id path int true "Story ID"
// @Accept  json
// @Produce  json
// @Success 200 {object} dtos.Post
// @Failure 400,404 {object} errorResponse
// @Failure 500 {object} errorResponse
// @Failure default {object} errorResponse
// @Router /api/stories/{:story_id} [get]
func (h *Handler) getStory(c *gin.Context) {

	postId, err := strconv.Atoi(c.Param("story_id"))
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	story, err := h.useCases.GetStory(postId)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}
	c.JSON(http.StatusOK, map[string]interface{}{
		"story:": story,
	})
}

// @Summary Update Story By Id
// @Security ApiKeyAuth
// @Tags stories
// @Description update story by id
// @ID update-story-by-id
// @Param :story_id path int true "Story ID"
// @Accept  json
// @Produce  json
// @Param input body dtos.UpdateStoryInput true "account info"
// @Success 200 {object} map[string]interface{}
// @Failure 400,404 {object} errorResponse
// @Failure 500 {object} errorResponse
// @Failure default {object} errorResponse
// @Router /api/stories/{:story_id} [put]
func (h *Handler) updateStory(c *gin.Context) {
	var input dtos.UpdateStoryInput
	if err := c.BindJSON(&input); err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	postId, err := strconv.Atoi(c.Param("story_id"))
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}
	userId, err := getUserId(c)
	if err != nil {
		newErrorResponse(c, http.StatusForbidden, err.Error())
		return
	}
	role, err := getRole(c)
	if err != nil {
		newErrorResponse(c, http.StatusForbidden, err.Error())
		return
	}
	err = h.useCases.UpdateStory(postId, userId, role, input)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}
	c.JSON(http.StatusOK, map[string]interface{}{
		"message": "updated",
	})
}

// @Summary Delete Story By Id
// @Security ApiKeyAuth
// @Tags stories
// @Description delete story by id
// @ID delete-story-by-id
// @Param :story_id path int true "Post ID"
// @Accept  json
// @Produce  json
// @Success 200 {object} map[string]interface{}
// @Failure 400,404 {object} errorResponse
// @Failure 500 {object} errorResponse
// @Failure default {object} errorResponse
// @Router /api/stories/{:story_id} [delete]
func (h *Handler) deleteStory(c *gin.Context) {
	postId, err := strconv.Atoi(c.Param("story_id"))
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}
	userId, err := getUserId(c)
	if err != nil {
		newErrorResponse(c, http.StatusForbidden, err.Error())
		return
	}

	role, err := getRole(c)
	if err != nil {
		newErrorResponse(c, http.StatusForbidden, err.Error())
		return
	}

	err = h.useCases.DeleteStory(postId, userId, role)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, map[string]interface{}{
		"message": "deleted",
	})

}
