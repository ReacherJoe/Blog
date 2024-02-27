package controllers

import (
	models "gorm_api/Models"
	"net/http"
	"strconv"
	"time"

	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

type PostController struct {
	DB *gorm.DB
}

func (u PostController) GetPosts(c echo.Context) error {

	var posts []models.Post

	result := u.DB.Find(&posts)

	if result.Error != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{
			"error": "Internal Server Error",
		})
	}

	return c.JSON(http.StatusOK, posts)
}

func (u PostController) GetPost(c echo.Context) error {

	id, _ := strconv.ParseInt(c.Param("id"), 10, 64)

	var posts models.Post

	result := u.DB.First(&posts, id)

	if result.Error != nil {
		return c.JSON(http.StatusNotFound, map[string]string{
			"status": "Not found post by this id",
		})
	}

	return c.JSON(http.StatusOK, posts)
}

func (u PostController) CreatePost(c echo.Context) error {

	var post models.Post

	if err := c.Bind(&post); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "Bad request"})
	}

	post.CreatedAt = time.Now()
	post.UpdatedAt = time.Now()

	result := u.DB.Create(&post)

	if result.Error != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{
			"error": "Internal Server Error",
		})
	}

	return c.JSON(http.StatusCreated, map[string]string{"status": "created"})
}

func (u PostController) UpdatePost(c echo.Context) error {

	id, _ := strconv.ParseInt(c.Param("id"), 10, 64)

	var post models.Post
	var existing_post models.Post

	if err := c.Bind(&post); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "Bad request"})
	}

	post.UpdatedAt = time.Now()

	result := u.DB.First(&existing_post, id)

	if result.Error != nil {
		return c.JSON(http.StatusNotFound, map[string]string{
			"status": "post not found by this id",
		})
	}

	existing_post.Title = post.Title
	existing_post.Body = post.Body
	existing_post.Photo = post.Photo
	existing_post.UpdatedAt = post.UpdatedAt

	result = u.DB.Save(&existing_post)

	if result.Error != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{
			"error": "Internal server error",
		})
	}

	return c.JSON(http.StatusOK, map[string]string{"status": "updated success"})
}

func (u PostController) DeletePost(c echo.Context) error {

	id, _ := strconv.ParseInt(c.Param("id"), 10, 64)
	var existing_post models.Post
	result := u.DB.First(&existing_post, id)
	if result.Error != nil {
		return c.JSON(http.StatusNotFound, map[string]string{
			"status": "Not found existing_post by this id",
		})
	}

	u.DB.Delete(&existing_post)

	return c.JSON(http.StatusOK, map[string]string{
		"status": "existing_post deleted success",
	})
}
