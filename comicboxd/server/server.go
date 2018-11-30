package server

import (
	"context"
	"fmt"
	"net/http"
	"time"

	"github.com/zwzn/comicbox/comicboxd/j"
	"github.com/spf13/viper"

	_ "github.com/golang-migrate/migrate/source/file"
	"github.com/gorilla/mux"
	_ "github.com/mattn/go-sqlite3"
)

// Server is the http server for the service. It runs the web interface and the API
type Server struct {
	srv *http.Server

	Router *mux.Router
	// DB     *sqlx.DB
}

func New() *Server {
	s := &Server{}

	s.Router = mux.NewRouter()

	s.srv = &http.Server{
		Addr:    fmt.Sprintf(":%d", viper.GetInt("port")),
		Handler: s.Router,

		WriteTimeout: 30 * time.Second,
		ReadTimeout:  30 * time.Second,
	}

	return s
}

func (s *Server) Start() error {
	j.Infof("Starting server at http://localhost:%d", viper.GetInt("port"))
	return s.srv.ListenAndServe()
}

func (s *Server) Stop() error {
	// err := s.DB.Close()
	// if err != nil {
	// 	return err
	// }
	return s.srv.Shutdown(context.TODO())
}
