package routes

import (
	"github.com/JGurus/template-initial-api/api/handlers"
	middleware "github.com/JGurus/template-initial-api/api/middlewares"
	"github.com/labstack/echo/v4"
)

//UserRoutes .
func UserRoutes(e *echo.Echo, s handlers.UserHandler) {
	handler := handlers.NewUser(s)
	users := e.Group("/v1/users")
	users.POST("", handler.Create)
	users.PUT("/:id", middleware.Auth(handler.Update))
	users.DELETE("/:id", middleware.Auth(handler.Delete))
	users.GET("", middleware.Auth(handler.GetAll))
	users.GET("/:id", middleware.Auth(handler.GetByID))
}

//LoginRoutes .
func LoginRoutes(e *echo.Echo, s handlers.UserHandler) {
	handler := handlers.NewLogin(s)
	e.POST("/v1/login", handler.LogIn)
}

//RegisterRoutes .
func RegisterRoutes(e *echo.Echo, s handlers.UserHandler) {
	handler := handlers.NewRegister(s)
	e.POST("/v1/register", handler.Signup)
}
