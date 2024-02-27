package controllers

import (
	models "gorm_api/Models"
	"net/http"
	"strconv"
	"time"

	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

type ReplyController struct {
	DB *gorm.DB
}

func (u ReplyController) GetReplys(c echo.Context) error {

	var Replys []models.Reply

	result := u.DB.Find(&Replys)

	if result.Error != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{
			"error": "Internal Server Error",
		})
	}

	return c.JSON(http.StatusOK, Replys)
}

func (u ReplyController) GetReply(c echo.Context) error {

	id, _ := strconv.ParseInt(c.Param("id"), 10, 64)

	var Replys models.Reply

	result := u.DB.First(&Replys, id)

	if result.Error != nil {
		return c.JSON(http.StatusNotFound, map[string]string{
			"status": "Not found Reply by this id",
		})
	}

	return c.JSON(http.StatusOK, Replys)
}

func (u ReplyController) CreateReply(c echo.Context) error {

	var Replys models.Reply

	if err := c.Bind(&Replys); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "Bad request"})
	}

	Replys.CreatedAt = time.Now()
	Replys.UpdatedAt = time.Now()

	result := u.DB.Create(&Replys)

	if result.Error != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{
			"error": "Internal Server Error",
		})
	}

	return c.JSON(http.StatusCreated, map[string]string{"status": "created"})
}

func (u ReplyController) UpdateReply(c echo.Context) error {

	id, _ := strconv.ParseInt(c.Param("id"), 10, 64)

	var Replys models.Reply
	var existing_Replys models.Reply

	if err := c.Bind(&Replys); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "Bad request"})
	}

	Replys.UpdatedAt = time.Now()

	result := u.DB.First(&existing_Replys, id)

	if result.Error != nil {
		return c.JSON(http.StatusNotFound, map[string]string{
			"status": "Reply not found by this id",
		})
	}

	existing_Replys.Text = Replys.Text
	existing_Replys.UpdatedAt = Replys.UpdatedAt

	result = u.DB.Save(&existing_Replys)

	if result.Error != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{
			"error": "Internal server error",
		})
	}

	return c.JSON(http.StatusOK, map[string]string{"status": "updated success"})
}

func (u ReplyController) DeleteReply(c echo.Context) error {

	id, _ := strconv.ParseInt(c.Param("id"), 10, 64)
	var existing_Replys models.Reply
	result := u.DB.First(&existing_Replys, id)
	if result.Error != nil {
		return c.JSON(http.StatusNotFound, map[string]string{
			"status": "Not found existing_Reply by this id",
		})
	}

	u.DB.Delete(&existing_Replys)

	return c.JSON(http.StatusOK, map[string]string{
		"status": "existing_Reply deleted success",
	})
}
