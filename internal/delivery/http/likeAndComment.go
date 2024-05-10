package http

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/isido5ik/StoryPublishingPlatform/dtos"
	"github.com/sirupsen/logrus"
)

// @Summary Like Story
// @Security ApiKeyAuth
// @Tags Like and Comment
// @Description like story
// @ID like-story
// @Accept json
// @Produce json
// @Param :story_id path int true "Story ID"
// @Success 200 {integer} integer 1
// @Failure 400,401 {object} errorResponse
// @Failure 500 {object} errorResponse
// @Failure default {object} errorResponse
// @Router /api/stories/{:story_id}/like [put]
func (h *Handler) like(c *gin.Context) {
	userId, err := getUserId(c)
	if err != nil {
		logrus.WithError(err).Error("Failed to get user ID")
		newErrorResponse(c, http.StatusUnauthorized, err.Error())
		return
	}

	postId, err := strconv.Atoi(c.Param("story_id"))
	if err != nil {
		logrus.WithError(err).Error("Failed to parse story ID parameter")
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}
	likeId, err := h.useCases.Like(userId, postId)
	if err != nil {
		logrus.WithError(err).Error("Failed to like post")
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, map[string]interface{}{
		"likeId": likeId,
	})
}

// @Summary Remove Like
// @Security ApiKeyAuth
// @Tags Like and Comment
// @Description remove like
// @ID remove-like
// @Accept json
// @Produce json
// @Param :story_id path int true "Story ID"
// @Success 200 {object} map[string]interface{}
// @Failure 400,401 {object} errorResponse
// @Failure 500 {object} errorResponse
// @Failure default {object} errorResponse
// @Router /api/stories/{:story_id}/like [delete]
func (h *Handler) removeLike(c *gin.Context) {
	postId, err := strconv.Atoi(c.Param("story_id"))
	if err != nil {
		logrus.WithError(err).Error("Failed to parse story ID parameter")
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	userId, err := getUserId(c)
	if err != nil {
		logrus.WithError(err).Error("Failed to get user ID")
		newErrorResponse(c, http.StatusUnauthorized, err.Error())
		return
	}

	err = h.useCases.RemoveLike(userId, postId)
	if err != nil {
		logrus.WithError(err).Error("Failed to remove like from post")
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}
	c.JSON(http.StatusOK, map[string]interface{}{
		"message": "like removed",
	})
}

// @Summary Add Comment
// @Security ApiKeyAuth
// @Tags Like and Comment
// @Description add comment
// @ID add-comment
// @Accept json
// @Produce json
// @Param :story_id path int true "Story ID"
// @Param input body dtos.AddCommentInput true "comment info"
// @Success 200 {integer} integer 1
// @Failure 400,401 {object} errorResponse
// @Failure 500 {object} errorResponse
// @Failure default {object} errorResponse
// @Router /api/stories/{:story_id}/comment [post]
func (h *Handler) addComment(c *gin.Context) {
	userId, err := getUserId(c)
	if err != nil {
		logrus.WithError(err).Error("Failed to get user ID")
		newErrorResponse(c, http.StatusUnauthorized, err.Error())
		return
	}

	postId, err := strconv.Atoi(c.Param("story_id"))
	if err != nil {
		logrus.WithError(err).Error("Failed to parse story ID parameter")
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	var comment dtos.AddCommentInput
	if err := c.BindJSON(&comment); err != nil {
		logrus.WithError(err).Error("Failed to bind JSON")
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	commentId, err := h.useCases.AddComment(userId, postId, comment)
	if err != nil {
		logrus.WithError(err).Error("Failed to add comment")
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, map[string]interface{}{
		"commentId": commentId,
	})
}

// @Summary Reply to Comment
// @Security ApiKeyAuth
// @Tags Like and Comment
// @Description reply to comment
// @ID reply-to-comment
// @Accept json
// @Produce json
// @Param :story_id path int true "Story ID"
// @Param :comment_id path int true "Comment ID"
// @Param input body dtos.AddCommentInput true "comment info"
// @Success 200 {integer} integer 1
// @Failure 400,401 {object} errorResponse
// @Failure 500 {object} errorResponse
// @Failure default {object} errorResponse
// @Router /api/stories/{:story_id}/comment/{:comment_id} [post]
func (h *Handler) replyToComment(c *gin.Context) {
	userId, err := getUserId(c)
	if err != nil {
		logrus.WithError(err).Error("Failed to get user ID")
		newErrorResponse(c, http.StatusUnauthorized, err.Error())
		return
	}
	postId, err := strconv.Atoi(c.Param("story_id"))
	if err != nil {
		logrus.WithError(err).Error("Failed to parse story ID parameter")
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}
	parentCommentID, err := strconv.Atoi(c.Param("comment_id"))
	if err != nil {
		logrus.WithError(err).Error("Failed to parse comment ID parameter")
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	var comment dtos.AddCommentInput
	if err := c.BindJSON(&comment); err != nil {
		logrus.WithError(err).Error("Failed to bind JSON")
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	commentId, err := h.useCases.ReplyToComment(userId, postId, parentCommentID, comment)
	if err != nil {
		logrus.WithError(err).Error("Failed to reply to comment")
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}
	c.JSON(http.StatusOK, map[string]interface{}{
		"commentId": commentId,
	})
}

// @Summary Update Comment
// @Security ApiKeyAuth
// @Tags Like and Comment
// @Description update comment
// @ID update-comment
// @Accept json
// @Produce json
// @Param :story_id path int true "Story ID"
// @Param :comment_id path int true "Comment ID"
// @Param input body dtos.UpdateCommentInput true "comment info"
// @Success 200 {object} map[string]interface{}
// @Failure 400,401 {object} errorResponse
// @Failure 500 {object} errorResponse
// @Failure default {object} errorResponse
// @Router /api/stories/{:story_id}/comment/{:comment_id} [put]
func (h *Handler) updateComment(c *gin.Context) {

	commentID, err := strconv.Atoi(c.Param("comment_id"))
	if err != nil {
		logrus.WithError(err).Error("Failed to parse comment ID parameter")
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	var newComment dtos.UpdateCommentInput
	if err := c.BindJSON(&newComment); err != nil {
		logrus.WithError(err).Error("Failed to bind JSON")
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}
	postID, err := strconv.Atoi(c.Param("story_id"))
	if err != nil {
		logrus.WithError(err).Error("Failed to parse story ID parameter")
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}
	userID, err := getUserId(c)
	if err != nil {
		logrus.WithError(err).Error("Failed to get user ID")
		newErrorResponse(c, http.StatusUnauthorized, err.Error())
		return
	}

	err = h.useCases.UpdateComment(userID, postID, commentID, newComment)
	if err != nil {
		logrus.WithError(err).Error("Failed to update comment")
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}
	c.JSON(http.StatusOK, map[string]interface{}{
		"message": "comment edited",
	})
}

// @Summary Delete Comment
// @Security ApiKeyAuth
// @Tags Like and Comment
// @Description delete comment
// @ID delete-comment
// @Accept json
// @Produce json
// @Param :story_id path int true "Story ID"
// @Param :comment_id path int true "Comment ID"
// @Success 200 {object} map[string]interface{}
// @Failure 400,401 {object} errorResponse
// @Failure 500 {object} errorResponse
// @Failure default {object} errorResponse
// @Router /api/stories/{:story_id}/comment/{:comment_id} [delete]
func (h *Handler) deleteComment(c *gin.Context) {
	commentID, err := strconv.Atoi(c.Param("comment_id"))
	if err != nil {
		logrus.WithError(err).Error("Failed to parse comment ID parameter")
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}
	userID, err := getUserId(c)

	if err != nil {
		logrus.WithError(err).Error("Failed to get user ID")
		newErrorResponse(c, http.StatusUnauthorized, err.Error())
		return
	}
	postID, err := strconv.Atoi(c.Param("story_id"))
	if err != nil {
		logrus.WithError(err).Error("Failed to parse story ID parameter")
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}
	err = h.useCases.DeleteComment(userID, postID, commentID)
	if err != nil {
		logrus.WithError(err).Error("Failed to delete comment")
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}
	c.JSON(http.StatusOK, map[string]interface{}{
		"message": "comment deleted",
	})
}
