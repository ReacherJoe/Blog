package controllers

import (
	models "gorm_api/Models"
	"net/http"
	"strconv"
	"time"

	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

type CommentController struct {
	DB *gorm.DB
}

func (u CommentController) GetComments(c echo.Context) error {

	var Comments []models.Comment

	result := u.DB.Find(&Comments)

	if result.Error != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{
			"error": "Internal Server Error",
		})
	}

	return c.JSON(http.StatusOK, Comments)
}

func (u CommentController) GetComment(c echo.Context) error {

	id, _ := strconv.ParseInt(c.Param("id"), 10, 64)

	var Comments models.Comment

	result := u.DB.First(&Comments, id)

	if result.Error != nil {
		return c.JSON(http.StatusNotFound, map[string]string{
			"status": "Not found Comment by this id",
		})
	}

	return c.JSON(http.StatusOK, Comments)
}

func (u CommentController) CreateComment(c echo.Context) error {

	var Comments models.Comment

	if err := c.Bind(&Comments); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "Bad request"})
	}

	Comments.CreatedAt = time.Now()
	Comments.UpdatedAt = time.Now()

	result := u.DB.Create(&Comments)

	if result.Error != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{
			"error": "Internal Server Error",
		})
	}

	return c.JSON(http.StatusCreated, map[string]string{"status": "created"})
}

func (u CommentController) UpdateComment(c echo.Context) error {

	id, _ := strconv.ParseInt(c.Param("id"), 10, 64)

	var Comments models.Comment
	var existing_Comments models.Comment

	if err := c.Bind(&Comments); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "Bad request"})
	}

	Comments.UpdatedAt = time.Now()

	result := u.DB.First(&existing_Comments, id)

	if result.Error != nil {
		return c.JSON(http.StatusNotFound, map[string]string{
			"status": "Comment not found by this id",
		})
	}

	existing_Comments.Text = Comments.Text
	existing_Comments.UpdatedAt = Comments.UpdatedAt

	result = u.DB.Save(&existing_Comments)

	if result.Error != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{
			"error": "Internal server error",
		})
	}

	return c.JSON(http.StatusOK, map[string]string{"status": "updated success"})
}

func (u CommentController) DeleteComment(c echo.Context) error {

	id, _ := strconv.ParseInt(c.Param("id"), 10, 64)
	var existing_Comments models.Comment
	result := u.DB.First(&existing_Comments, id)
	if result.Error != nil {
		return c.JSON(http.StatusNotFound, map[string]string{
			"status": "Not found existing_Comment by this id",
		})
	}

	u.DB.Delete(&existing_Comments)

	return c.JSON(http.StatusOK, map[string]string{
		"status": "existing_Comment deleted success",
	})
}
