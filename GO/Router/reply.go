// user.go
package routes

import (
	controllers "gorm_api/Controllers"
	databases "gorm_api/Databases"
	"log"

	"github.com/labstack/echo/v4"
)

func SetupReplyRoutes(e *echo.Echo) {

	db, err := databases.Connect()
	if err != nil {
		log.Println(err)
	}

	replyController := controllers.ReplyController{DB: db}

	e.GET("/replys", replyController.GetReplys)
	e.GET("/replys/:id", replyController.GetReply)
	e.POST("/replys", replyController.CreateReply)
	e.PUT("/replys/:id", replyController.UpdateReply)
	e.DELETE("/replys/:id", replyController.DeleteReply)
}
