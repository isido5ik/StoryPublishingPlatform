package http

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

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
