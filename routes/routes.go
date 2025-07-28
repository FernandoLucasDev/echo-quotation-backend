package routes

import (
	"api-echo/controllers"
	"api-echo/middlewares"

	"github.com/labstack/echo/v4"
)

func Init(e *echo.Echo) {

	e.POST("/request-code", controllers.CreateActivationCode)
	e.POST("/login", controllers.Login)
	e.POST("/create-user", controllers.CreateUser)
	e.GET("/news", controllers.GetNews)

	protected := e.Group("/users", middlewares.AuthMiddleware)

	protected.GET("", controllers.GetUsers)
	protected.GET("/:id", controllers.GetUserByID)
	protected.POST("", controllers.CreateUser)
	protected.PUT("/:id", controllers.UpdateUser)
	protected.DELETE("/:id", controllers.DeleteUser)
}
