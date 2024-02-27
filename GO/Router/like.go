// user.go
package routes

import (
	controllers "gorm_api/Controllers"
	databases "gorm_api/Databases"
	"log"

	"github.com/labstack/echo/v4"
)

func SetupLikeRoutes(e *echo.Echo) {

	db, err := databases.Connect()
	if err != nil {
		log.Println(err)
	}

	likecontroller := controllers.LikeController{DB: db}

	e.GET("/likes", likecontroller.GetLikes)
	e.GET("/likes/:id", likecontroller.GetLike)
	e.POST("/likes", likecontroller.CreateLike)
	e.PUT("/likes/:id", likecontroller.UpdateLike)
	e.DELETE("/likes/:id", likecontroller.DeleteLike)
}
