package routes

import (
	"github.com/labstack/echo/v4"
	"api-echo/controllers"
)

func Init(e *echo.Echo) {
	e.GET("/users", controllers.GetUsers)
	e.GET("/users/:id", controllers.GetUserByID)
	e.POST("/users", controllers.CreateUser)
	e.PUT("/users/:id", controllers.UpdateUser)
	e.DELETE("/users/:id", controllers.DeleteUser)
}
