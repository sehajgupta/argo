package server

import (
	"github.com/gin-gonic/gin"
	"github.com/jmoiron/sqlx"
)

// server struct
type Server struct {
	Router *gin.Engine
	DB     *sqlx.DB
}

// NewServer creates a new server
func NewServer(db *sqlx.DB) *Server {
	return &Server{
		DB:     db,
		Router: gin.Default(),
	}
}
