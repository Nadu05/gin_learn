package routes

import (
	"ginLearn/internal/app"

	"github.com/gin-gonic/gin"
)

func RegisterUserRoutes(r *gin.Engine, app *app.App) {

	user := r.Group("/users")
	{
		user.GET("/:id", app.UserHandler.GetUserByID)
	}
}
