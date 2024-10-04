package config

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

type Server struct {
	Host string
	Port string
	DB   *gorm.DB
}

func NewServer(server Server) error {
	e := echo.New()

	return e.StartServer(&http.Server{
		Addr: server.Host + ":" + server.Port,
	})
}
