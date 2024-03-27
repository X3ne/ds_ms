package server

import (
	"github.com/X3ne/ds_ms/gateway/config"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"net/http"
)

type Server struct {
	Echo   *echo.Echo
	Config *config.Config
}

func NewServer(cfg *config.Config) *Server {
	return &Server{
		Echo:   echo.New(),
		Config: cfg,
	}
}

func (s *Server) Start(port string) error {

	s.Echo.HideBanner = true

	s.Echo.Use(middleware.Recover())

	s.Echo.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"http://localhost:5173"},
		AllowHeaders: []string{echo.HeaderOrigin, echo.HeaderContentType, echo.HeaderAccept},
		AllowMethods: []string{http.MethodGet, http.MethodPut, http.MethodPost, http.MethodDelete},
	}))

	//s.Echo.Use(middleware.Logger())

	return s.Echo.Start(":" + port)
}
