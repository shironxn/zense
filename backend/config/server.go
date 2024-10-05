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
	}
}

func (s *Server) Run() error {
	e := echo.New()

	jwt := util.NewJWT(s.JWT.Secret)
	validator := validator.New(validator.WithRequiredStructEnabled())

	userRepository := repository.NewUserRepository(s.DB)
	userService := service.NewUserService(userRepository, *jwt)
	userHandler := handler.NewUserHandler(userService, validator)

	router := https.NewRouter(e, userHandler)

	s.DB.AutoMigrate(&domain.User{}, &domain.Journal{}, &domain.Forum{}, &domain.Comment{})

	return e.StartServer(&http.Server{
		Addr:    s.Host + ":" + s.Port,
		Handler: router.Run(),
	})
}
