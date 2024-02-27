package controllers

import (
	models "gorm_api/Models"
	"net/http"
	"strconv"
	"time"

	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

type LikeController struct {
	DB *gorm.DB
}

func (u LikeController) GetLikes(c echo.Context) error {

	var Likes []models.Like

	result := u.DB.Find(&Likes)

	if result.Error != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{
			"error": "Internal Server Error",
		})
	}

	return c.JSON(http.StatusOK, Likes)
}

func (u LikeController) GetLike(c echo.Context) error {

	id, _ := strconv.ParseInt(c.Param("id"), 10, 64)

	var Likes models.Like

	result := u.DB.First(&Likes, id)

	if result.Error != nil {
		return c.JSON(http.StatusNotFound, map[string]string{
			"status": "Not found Like by this id",
		})
	}

	return c.JSON(http.StatusOK, Likes)
}

func (u LikeController) CreateLike(c echo.Context) error {

	var Likes models.Like

	if err := c.Bind(&Likes); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "Bad request"})
	}

	Likes.CreatedAt = time.Now()
	Likes.UpdatedAt = time.Now()

	result := u.DB.Create(&Likes)

	if result.Error != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{
			"error": "Internal Server Error",
		})
	}

	return c.JSON(http.StatusCreated, map[string]string{"status": "created"})
}

func (u LikeController) UpdateLike(c echo.Context) error {

	id, _ := strconv.ParseInt(c.Param("id"), 10, 64)

	var likes models.Like
	var existing_likes models.Like

	if err := c.Bind(&likes); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "Bad request"})
	}

	likes.UpdatedAt = time.Now()

	result := u.DB.First(&existing_likes, id)

	if result.Error != nil {
		return c.JSON(http.StatusNotFound, map[string]string{
			"status": "likes not found by this id",
		})
	}

	

	result = u.DB.Save(&existing_likes)

	if result.Error != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{
			"error": "Internal server error",
		})
	}

	return c.JSON(http.StatusOK, map[string]string{"status": "updated success"})
}

func (u LikeController) DeleteLike(c echo.Context) error {

	id, _ := strconv.ParseInt(c.Param("id"), 10, 64)
	var existing_Like models.Like
	result := u.DB.First(&existing_Like, id)
	if result.Error != nil {
		return c.JSON(http.StatusNotFound, map[string]string{
			"status": "Not found existing_Like by this id",
		})
	}

	u.DB.Delete(&existing_Like)

	return c.JSON(http.StatusOK, map[string]string{
		"status": "existing_Like deleted success",
	})
}
