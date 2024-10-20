package api

import (
	"github.com/RokayaEG/golang-library-service/service/genre"
	"github.com/gin-gonic/gin"
	"github.com/jmoiron/sqlx"
)

type APIServer struct {
	port string
	db   *sqlx.DB
}

func NewAPIServer(port string, db *sqlx.DB) *APIServer {
	return &APIServer{
		port: port,
		db:   db,
	}
}

func (s *APIServer) Run() {
	router := gin.Default()
	subrouter := router.Group("/api/v1")

	genreRouter := subrouter.Group("/genre")

	genreStore := genre.NewStore(s.db)
	genreHandler := genre.NewHandler(genreStore)
	genreHandler.RegisterRoutes(genreRouter)

	router.Run(s.port)
}
