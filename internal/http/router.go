package http

import (
	"net/http"

	"github.com/MarceloPetrucio/go-scalar-api-reference"
	"github.com/aternity/zense/internal/handler"
	"github.com/aternity/zense/internal/util"
	"github.com/labstack/echo-jwt/v4"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

type Router struct {
	e        *echo.Echo
	jwt      *util.JWT
	handlers Handlers
}

type Handlers struct {
	User    handler.UserHandler
	Journal handler.JournalHandler
	Topic   handler.TopicHandler
	Comment handler.CommentHandler
	Forum   handler.ForumHandler
	Vent    handler.VentHandler
}

func NewRouter(
	e *echo.Echo,
	jwt *util.JWT,
	handlers Handlers,
) *Router {
	return &Router{
		e:        e,
		jwt:      jwt,
		handlers: handlers,
	}
}

func (r *Router) Run() http.Handler {
	r.e.Use(middleware.Logger())
	r.e.Use(middleware.Recover())
	r.setupCORS()
	r.setupJWT()

	r.e.GET("/", func(c echo.Context) error {
		return c.JSON(http.StatusOK, "Welcome To Zense")
	})
	api := r.e.Group("/api/v1")
	r.setupRoutes(api)

	return r.e
}

func (r *Router) setupCORS() {
	r.e.Use(middleware.CORS())
}

func (r *Router) setupJWT() {
	r.e.Use(echojwt.WithConfig(echojwt.Config{
		SigningKey:  []byte(r.jwt.Secret),
		TokenLookup: "header:Authorization",
		Skipper: func(c echo.Context) bool {
			switch c.Path() {
			case "/api/v1/auth/login", "/api/v1/auth/register":
				return true
			case "/api/v1/users/me":
				return false
			default:
				return c.Request().Method == http.MethodGet
			}
		},
	}))
}

func (r *Router) setupRoutes(api *echo.Group) {
	auth := api.Group("/auth")
	users := api.Group("/users")
	journals := api.Group("/journals")
	forums := api.Group("/forums")
	comments := api.Group("/comments")
	topics := api.Group("/topics")
	vents := api.Group("/vents")

	api.GET("/docs", func(c echo.Context) error {
		htmlContent, err := scalar.ApiReferenceHTML(&scalar.Options{
			SpecURL: "./docs/swagger.json",
			CustomOptions: scalar.CustomOptions{
				PageTitle: "Zense API Docs",
			},
			DarkMode: true,
		})
		if err != nil {
			return echo.NewHTTPError(http.StatusInternalServerError, "Failed to load API docs: "+err.Error())
		}
		return c.HTML(http.StatusOK, htmlContent)
	})

	auth.POST("/login", r.handlers.User.Login)
	auth.POST("/register", r.handlers.User.Register)

	users.GET("/me", r.handlers.User.FindMe)
	users.GET("", r.handlers.User.FindAll)
	users.GET("/:id", r.handlers.User.FindByID)
	users.PUT("/:id", r.handlers.User.Update)
	users.DELETE("/:id", r.handlers.User.Delete)

	journals.POST("", r.handlers.Journal.Create)
	journals.GET("", r.handlers.Journal.FindAll)
	journals.GET("/:id", r.handlers.Journal.FindByID)
	journals.PUT("/:id", r.handlers.Journal.Update)
	journals.DELETE("/:id", r.handlers.Journal.Delete)

	comments.POST("", r.handlers.Comment.Create)
	comments.GET("", r.handlers.Comment.FindAll)
	comments.GET("/:id", r.handlers.Comment.FindByID)
	comments.PUT("/:id", r.handlers.Comment.Update)
	comments.DELETE("/:id", r.handlers.Comment.Delete)

	topics.POST("", r.handlers.Topic.Create)
	topics.GET("", r.handlers.Topic.FindAll)
	topics.GET("/:id", r.handlers.Topic.FindByID)
	topics.PUT("/:id", r.handlers.Topic.Update)
	topics.DELETE("/:id", r.handlers.Topic.Delete)

	forums.POST("", r.handlers.Forum.Create)
	forums.GET("", r.handlers.Forum.FindAll)
	forums.GET("/:id", r.handlers.Forum.FindByID)
	forums.PUT("/:id", r.handlers.Forum.Update)
	forums.DELETE("/:id", r.handlers.Forum.Delete)
	forums.DELETE("/:id/topic", r.handlers.Forum.RemoveTopic)

	vents.POST("", r.handlers.Vent.Chat)
	vents.DELETE("", r.handlers.Vent.Clear)
}
