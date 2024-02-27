package controllers

import (
	models "gorm_api/Models"
	"net/http"
	"strconv"
	"time"

	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

type CategorieController struct {
	DB *gorm.DB
}

func (u CategorieController) GetCategories(c echo.Context) error {

	var Categories []models.Categorie

	result := u.DB.Find(&Categories)

	if result.Error != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{
			"error": "Internal Server Error",
		})
	}

	return c.JSON(http.StatusOK, Categories)
}

func (u CategorieController) GetCategorie(c echo.Context) error {

	id, _ := strconv.ParseInt(c.Param("id"), 10, 64)

	var Categories models.Categorie

	result := u.DB.First(&Categories, id)

	if result.Error != nil {
		return c.JSON(http.StatusNotFound, map[string]string{
			"status": "Not found Categorie by this id",
		})
	}

	return c.JSON(http.StatusOK, Categories)
}

func (u CategorieController) CreateCategorie(c echo.Context) error {

	var Categories models.Categorie

	if err := c.Bind(&Categories); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "Bad request"})
	}

	Categories.CreatedAt = time.Now()
	Categories.UpdatedAt = time.Now()

	result := u.DB.Create(&Categories)

	if result.Error != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{
			"error": "Internal Server Error",
		})
	}

	return c.JSON(http.StatusCreated, map[string]string{"status": "created"})
}

func (u CategorieController) UpdateCategorie(c echo.Context) error {

	id, _ := strconv.ParseInt(c.Param("id"), 10, 64)

	var Categories models.Categorie
	var existing_Categories models.Categorie

	if err := c.Bind(&Categories); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "Bad request"})
	}

	Categories.UpdatedAt = time.Now()

	result := u.DB.First(&existing_Categories, id)

	if result.Error != nil {
		return c.JSON(http.StatusNotFound, map[string]string{
			"status": "Categorie not found by this id",
		})
	}

	existing_Categories.Name = Categories.Name
	existing_Categories.UpdatedAt = Categories.UpdatedAt

	result = u.DB.Save(&existing_Categories)

	if result.Error != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{
			"error": "Internal server error",
		})
	}

	return c.JSON(http.StatusOK, map[string]string{"status": "updated success"})
}

func (u CategorieController) DeleteCategorie(c echo.Context) error {

	id, _ := strconv.ParseInt(c.Param("id"), 10, 64)
	var existing_Categories models.Categorie
	result := u.DB.First(&existing_Categories, id)
	if result.Error != nil {
		return c.JSON(http.StatusNotFound, map[string]string{
			"status": "Not found existing_Categorie by this id",
		})
	}

	u.DB.Delete(&existing_Categories)

	return c.JSON(http.StatusOK, map[string]string{
		"status": "existing_Categorie deleted success",
	})
}
