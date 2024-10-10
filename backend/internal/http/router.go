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
	e       *echo.Echo
	jwt     *util.JWT
	user    handler.UserHandler
	journal handler.JournalHandler
	forum   handler.ForumHandler
	topic   handler.TopicHandler
	comment handler.CommentHandler
}

func NewRouter(
	e *echo.Echo,
	jwt *util.JWT,
	user handler.UserHandler,
	journal handler.JournalHandler,
	forum handler.ForumHandler,
	topic handler.TopicHandler,
	comment handler.CommentHandler,
) *Router {
	return &Router{
		e:       e,
		jwt:     jwt,
		user:    user,
		journal: journal,
		forum:   forum,
		topic:   topic,
		comment: comment,
	}
}

func (r *Router) Run() http.Handler {
	r.e.Use(middleware.Logger())
	r.e.Use(middleware.Recover())
	r.e.Use(echojwt.WithConfig(echojwt.Config{
		SigningKey:  []byte(r.jwt.Secret),
		TokenLookup: "header:Authorization",
		Skipper: func(c echo.Context) bool {
			if c.Path() == "/api/v1/users/me" {
				return false
			}
			if c.Path() == "/api/v1/auth/login" || c.Path() == "/api/v1/auth/register" || c.Request().Method == "GET" {
				return true
			}
			return false
		},
	}))

	api := r.e.Group("/api/v1")
	auth := api.Group("/auth")
	users := api.Group("/users")
	journals := api.Group("/journals")
	forums := api.Group("/forums")
	comments := api.Group("/comments")
	topics := api.Group("/topics")

	api.GET("/docs", func(c echo.Context) error {
		htmlContent, err := scalar.ApiReferenceHTML(&scalar.Options{
			SpecURL: "./docs/swagger.json",
			CustomOptions: scalar.CustomOptions{
				PageTitle: "Zense API Docs",
			},
			DarkMode: true,
		})
		if err != nil {
			return err
		}

		return c.HTML(http.StatusOK, htmlContent)
	})

	auth.POST("/login", r.user.Login)
	auth.POST("/register", r.user.Register)

	users.GET("/me", r.user.FindMe)
	users.GET("", r.user.FindAll)
	users.GET("/:id", r.user.FindByID)
	users.PUT("/:id", r.user.Update)
	users.DELETE("/:id", r.user.Delete)

	journals.POST("", r.journal.Create)
	journals.GET("", r.journal.FindAll)
	journals.GET("/:id", r.journal.FindByID)
	journals.PUT("/:id", r.journal.Update)
	journals.DELETE("/:id", r.journal.Delete)

	forums.POST("", r.forum.Create)
	forums.GET("", r.forum.FindAll)
	forums.GET("/:id", r.forum.FindByID)
	forums.PUT("/:id", r.forum.Update)
	forums.DELETE("/:id", r.forum.Delete)

	comments.POST("", r.comment.Create)
	comments.GET("", r.comment.FindAll)
	comments.GET("/:id", r.comment.FindByID)
	comments.PUT("/:id", r.comment.Update)
	comments.DELETE("/:id", r.comment.Delete)

	topics.POST("", r.topic.Create)
	topics.GET("", r.topic.FindAll)
	topics.GET("/:id", r.topic.FindByID)
	topics.PUT("/:id", r.topic.Update)
	topics.DELETE("/:id", r.topic.Delete)

	return r.e
}
