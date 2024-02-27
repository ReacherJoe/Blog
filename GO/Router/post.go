// user.go
package routes

import (
	controllers "gorm_api/Controllers"
	databases "gorm_api/Databases"
	"log"

	"github.com/labstack/echo/v4"
)

func SetupPostRoutes(e *echo.Echo) {

	db, err := databases.Connect()
	if err != nil {
		log.Println(err)
	}

	postcontroller := controllers.PostController{DB: db}

	e.GET("/posts", postcontroller.GetPosts)
	e.GET("/posts/:id", postcontroller.GetPost)
	e.POST("/posts", postcontroller.CreatePost)
	e.PUT("/posts/:id", postcontroller.UpdatePost)
	e.DELETE("/posts/:id", postcontroller.DeletePost)
}
