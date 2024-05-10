package dtos

import (
	"errors"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

type PaginationParams struct {
	Page     int `form:"page"`     // Page number
	PageSize int `form:"pageSize"` // Number of items per page
}

// ValidatePage validates the page parameters and returns their values.
func ValidatePage(c *gin.Context) (int, int, error) {
	pageStr := c.DefaultQuery("page", "1")         // Get the 'page' query parameter, defaulting to 1 if not present
	pageSizeStr := c.DefaultQuery("pageSize", "3") // Get the 'pageSize' query parameter, defaulting to 3 if not present

	page, err := strconv.Atoi(pageStr) // Convert 'page' string to integer
	if err != nil {
		logrus.WithError(err).Error("Failed to parse page parameter")
		return -1, -1, errors.New("invalid value of query parameter 'page' (it must be an integer)")
	}

	pageSize, err := strconv.Atoi(pageSizeStr) // Convert 'pageSize' string to integer
	if err != nil {
		logrus.WithError(err).Error("Failed to parse pageSize parameter")
		return -1, -1, errors.New("invalid value of query parameter 'pageSize' (it must be an integer)")
	}

	if pageSize <= 0 || page <= 0 {
		logrus.Error("Invalid value of query parameters 'page' or 'pageSize', negative number or zero")
		return -1, -1, errors.New("invalid value of query parameter 'page' or 'pageSize', negative number or zero")
	}

	logrus.WithFields(logrus.Fields{"page": page, "pageSize": pageSize}).Info("Valid page parameters")
	return page, pageSize, nil
}
