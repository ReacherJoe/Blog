// user.go
package routes

import (
	controllers "gorm_api/Controllers"
	databases "gorm_api/Databases"
	"log"

	"github.com/labstack/echo/v4"
)

func SetupUserRoutes(e *echo.Echo) {

	db, err := databases.Connect()
	if err != nil {
		log.Println(err)
	}

	userController := controllers.UserController{DB: db}

	e.GET("/users", userController.GetUsers)
	e.GET("/users/:id", userController.GetUser)
	e.POST("/users", userController.CreateUser)
	e.PUT("/users/:id", userController.UpdateUser)
	e.DELETE("/users/:id", userController.DeleteUser)
}
