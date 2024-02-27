package controllers

import (
	models "gorm_api/Models"
	"net/http"
	"strconv"
	"time"

	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

type Post_CategorieController struct {
	DB *gorm.DB
}

func (u Post_CategorieController) GetPost_Categories(c echo.Context) error {

	var Post_Categories []models.Post_Categorie

	result := u.DB.Find(&Post_Categories)

	if result.Error != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{
			"error": "Internal Server Error",
		})
	}

	return c.JSON(http.StatusOK, Post_Categories)
}

func (u Post_CategorieController) GetPost_Categorie(c echo.Context) error {

	id, _ := strconv.ParseInt(c.Param("id"), 10, 64)

	var Post_Categories models.Post_Categorie

	result := u.DB.First(&Post_Categories, id)

	if result.Error != nil {
		return c.JSON(http.StatusNotFound, map[string]string{
			"status": "Not found Post_Categorie by this id",
		})
	}

	return c.JSON(http.StatusOK, Post_Categories)
}

func (u Post_CategorieController) CreatePost_Categorie(c echo.Context) error {

	var Post_Categories models.Post_Categorie

	if err := c.Bind(&Post_Categories); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "Bad request"})
	}

	Post_Categories.CreatedAt = time.Now()
	Post_Categories.UpdatedAt = time.Now()

	result := u.DB.Create(&Post_Categories)

	if result.Error != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{
			"error": "Internal Server Error",
		})
	}

	return c.JSON(http.StatusCreated, map[string]string{"status": "created"})
}

func (u Post_CategorieController) UpdatePost_Categorie(c echo.Context) error {

	id, _ := strconv.ParseInt(c.Param("id"), 10, 64)

	var Post_Categories models.Post_Categorie
	var existing_Post_Categories models.Post_Categorie

	if err := c.Bind(&Post_Categories); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "Bad request"})
	}

	Post_Categories.UpdatedAt = time.Now()

	result := u.DB.First(&existing_Post_Categories, id)

	if result.Error != nil {
		return c.JSON(http.StatusNotFound, map[string]string{
			"status": "Post_Categorie not found by this id",
		})
	}

	
	existing_Post_Categories.UpdatedAt = Post_Categories.UpdatedAt

	result = u.DB.Save(&existing_Post_Categories)

	if result.Error != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{
			"error": "Internal server error",
		})
	}

	return c.JSON(http.StatusOK, map[string]string{"status": "updated success"})
}

func (u Post_CategorieController) DeletePost_Categorie(c echo.Context) error {

	id, _ := strconv.ParseInt(c.Param("id"), 10, 64)
	var existing_Post_Categories models.Post_Categorie
	result := u.DB.First(&existing_Post_Categories, id)
	if result.Error != nil {
		return c.JSON(http.StatusNotFound, map[string]string{
			"status": "Not found existing_Post_Categorie by this id",
		})
	}

	u.DB.Delete(&existing_Post_Categories)

	return c.JSON(http.StatusOK, map[string]string{
		"status": "existing_Post_Categorie deleted success",
	})
}
