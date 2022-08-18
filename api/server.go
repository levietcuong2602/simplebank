package api

import (
	"github.com/gin-gonic/gin"
	db "github.com/levietcuong2602/simplebank/db/sqlc"
)

// Server serves HTTP requests for banking service.
type Server struct {
	store  *db.Store
	router *gin.Engine
}

// NewServer creates a new HTTP server and setup routing
func NewServer(store *db.Store) *Server {
	server := &Server{store: store}
	router := gin.Default()

	router.POST("/accounts", server.CreateAccount)
	router.GET("/accounts", server.GetListAccounts)
	router.GET("/accounts/:id", server.GetAccount)

	server.router = router
	return server
}

// Start run http server on specific address
func (server *Server) Start(addr string) error {
	return server.router.Run(addr)
}

func errorResponse(err error) gin.H {
	return gin.H{"error": err.Error()}
}
