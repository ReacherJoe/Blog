// user.go
package routes

import (
	controllers "gorm_api/Controllers"
	databases "gorm_api/Databases"
	"log"

	"github.com/labstack/echo/v4"
)

func SetupCategorieRoutes(e *echo.Echo) {

	db, err := databases.Connect()
	if err != nil {
		log.Println(err)
	}

	categorieController := controllers.CategorieController{DB: db}

	e.GET("/categories", categorieController.GetCategories)
	e.GET("/categories/:id", categorieController.GetCategorie)
	e.POST("/categories", categorieController.CreateCategorie)
	e.PUT("/categories/:id", categorieController.UpdateCategorie)
	e.DELETE("/categories/:id", categorieController.DeleteCategorie)
}
