package routes

import (
	"github.com/labstack/echo/v4"
)

func Setup(e *echo.Echo) {

	SetupUserRoutes(e)
	SetupPostRoutes(e)
	SetupLikeRoutes(e)
	SetupCommentRoutes(e)
	SetupReplyRoutes(e)
	SetupCategorieRoutes(e)
	SetupPost_CategorieRouter(e)
}