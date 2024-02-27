
package routes

import (
	controllers "gorm_api/Controllers"
	databases "gorm_api/Databases"
	"log"

	"github.com/labstack/echo/v4"
)

func SetupCommentRoutes(e *echo.Echo) {

	db, err := databases.Connect()
	if err != nil {
		log.Println(err)
	}

	commentcontroller := controllers.CommentController{DB: db}

	e.GET("/comments", commentcontroller.GetComments)
	e.GET("/comments/:id", commentcontroller.GetComment)
	e.POST("/comments", commentcontroller.CreateComment)
	e.PUT("/comments/:id", commentcontroller.UpdateComment)
	e.DELETE("/comments/:id", commentcontroller.DeleteComment)
}
