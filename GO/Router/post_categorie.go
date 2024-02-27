// user.go
package routes

import (
	controllers "gorm_api/Controllers"
	databases "gorm_api/Databases"
	"log"

	"github.com/labstack/echo/v4"
)

func SetupPost_CategorieRoutes(e *echo.Echo) {

	db, err := databases.Connect()
	if err != nil {
		log.Println(err)
	}

	post_categorieController := controllers.Post_CategorieController{DB: db}

	e.GET("/post_categories", post_categorieController.GetPost_Categories)
	e.GET("/post_categories/:id", post_categorieController.GetPost_Categorie)
	e.POST("/post_categories", post_categorieController.CreatePost_Categorie)
	e.PUT("/post_categories/:id", post_categorieController.UpdatePost_Categorie)
	e.DELETE("/post_categories/:id", post_categorieController.DeletePost_Categorie)
}
