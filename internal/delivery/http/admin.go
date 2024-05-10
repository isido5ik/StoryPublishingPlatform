package http

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

// @Summary Get Users
// @Security ApiKeyAuth
// @Tags users
// @Description get all users
// @ID get-all-users
// @Accept  json
// @Produce  json
// @Success 200 {object} getUsersResponse
// @Failure 400,404 {object} errorResponse
// @Failure 500 {object} errorResponse
// @Failure default {object} errorResponse
// @Router /api/admin/users [get]
func (h *Handler) getUsers(c *gin.Context) {
	users, err := h.useCases.GetUsers()
	if err != nil {
		logrus.WithError(err).Error("Failed to get users")
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}
	logrus.Info("Users retrieved successfully")
	c.JSON(http.StatusOK, getUsersResponse{
		Data: users,
	})
}

// @Summary Get User By Id
// @Security ApiKeyAuth
// @Tags users
// @Description get user by id
// @ID get-story-by-id
// @Param :user_id path int true "User ID"
// @Accept  json
// @Produce  json
// @Success 200 {object} map[string]interface{}
// @Failure 400,404 {object} errorResponse
// @Failure 500 {object} errorResponse
// @Failure default {object} errorResponse
// @Router /api/admin/users/{:user_id} [get]
func (h *Handler) getUserById(c *gin.Context) {
	userId, err := strconv.Atoi(c.Param("user_id"))
	if err != nil {
		logrus.WithError(err).Error("Failed to parse user ID parameter")
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}
	user, err := h.useCases.GetUserById(userId)
	if err != nil {
		logrus.WithError(err).Error("Failed to get user by ID")
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}
	logrus.WithField("user_id", userId).Info("User retrieved successfully")
	c.JSON(http.StatusOK, map[string]interface{}{
		"user": user,
	})
}

// @Summary Delete User By Id
// @Security ApiKeyAuth
// @Tags users
// @Description delete user by id
// @ID delete-story-by-id
// @Param :user_id path int true "User ID"
// @Accept  json
// @Produce  json
// @Success 200 {object} map[string]interface{}
// @Failure 400,404 {object} errorResponse
// @Failure 500 {object} errorResponse
// @Failure default {object} errorResponse
// @Router /api/admin/users/{:user_id} [delete]
func (h *Handler) deleteUser(c *gin.Context) {
	userId, err := strconv.Atoi(c.Param("user_id"))
	if err != nil {
		logrus.WithError(err).Error("Failed to parse user ID parameter")
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}
	if err := h.useCases.DeleteUser(userId); err != nil {
		logrus.WithError(err).Error("Failed to delete user")
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}
	logrus.WithField("user_id", userId).Info("User deleted successfully")
	c.JSON(http.StatusOK, map[string]interface{}{
		"message": "user deleted",
	})
}
