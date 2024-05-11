package dtos

import (
	"errors"
	"strconv"

	"github.com/gin-gonic/gin"
)

func GetCategoryId(c *gin.Context) (int, error) {
	categoryIdStr := c.Query("category_id")
	categoryId, err := strconv.Atoi(categoryIdStr)
	if err != nil {
		return 0, errors.New("could not parse category_id")
	}

	if categoryId <= 0 {
		return categoryId, errors.New("invalid category_id")
	}
	return categoryId, nil
}
