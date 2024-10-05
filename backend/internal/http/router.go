package http

import (
	"net/http"

	"github.com/aternity/zense/internal/handler"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

type Router struct {
	e    *echo.Echo
	user handler.UserHandler
}

func NewRouter(e *echo.Echo, userHandler handler.UserHandler) *Router {
	return &Router{
		e:    e,
		user: userHandler,
	}
}

func (r *Router) Run() http.Handler {
	r.e.Use(middleware.Logger())
	r.e.Use(middleware.Recover())

	api := r.e.Group("/api/v1")
	auth := api.Group("/auth")
	users := api.Group("/users")

	auth.POST("/login", r.user.Login)
	auth.POST("/register", r.user.Register)

	users.GET("", r.user.FindAll)
	users.GET("/:id", r.user.FindByID)
	users.PUT("/:id", r.user.Update)
	users.DELETE("/:id", r.user.Delete)

	return r.e
}
