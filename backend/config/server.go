package config

import (
	"net/http"

	"github.com/aternity/zense/internal/entity/domain"
	"github.com/aternity/zense/internal/handler"
	https "github.com/aternity/zense/internal/http"
	"github.com/aternity/zense/internal/repository"
	"github.com/aternity/zense/internal/service"
	"github.com/aternity/zense/internal/util"
	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

type Server struct {
	Host string
	Port string
	DB   *gorm.DB
	JWT  util.JWT
}

func NewServer(server Server) *Server {
	return &Server{
		Host: server.Host,
		Port: server.Port,
		DB:   server.DB,
		JWT:  server.JWT,
	}
}

func (s *Server) Run() error {
	e := echo.New()

	jwt := util.NewJWT(s.JWT.Secret)
	validator := validator.New(validator.WithRequiredStructEnabled())

	userRepository := repository.NewUserRepository(s.DB)
	userService := service.NewUserService(userRepository, jwt)
	userHandler := handler.NewUserHandler(userService, validator)

	journalRepository := repository.NewJournalRepository(s.DB)
	journalService := service.NewJournalService(journalRepository)
	journalHandler := handler.NewJournalHandler(journalService, validator)

	commentRepository := repository.NewCommentRepository(s.DB)
	commentService := service.NewCommentService(commentRepository)
	commentHandler := handler.NewCommentHandler(commentService, validator)

	topicRepository := repository.NewTopicRepository(s.DB)
	topicService := service.NewTopicService(topicRepository)
	topicHandler := handler.NewTopicHandler(topicService, validator)

	forumRepository := repository.NewForumRepository(s.DB)
	forumService := service.NewForumService(forumRepository, topicRepository)
	forumHandler := handler.NewForumHandler(forumService, validator)

	router := https.NewRouter(e, jwt, userHandler, journalHandler, forumHandler, topicHandler, commentHandler)

	s.DB.AutoMigrate(&domain.User{}, &domain.Journal{}, &domain.Forum{}, &domain.Topic{}, &domain.Comment{})

	return e.StartServer(&http.Server{
		Addr:    s.Host + ":" + s.Port,
		Handler: router.Run(),
	})
}
