package api

import (
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"github.com/go-playground/validator/v10"
	db "github.com/levietcuong2602/simplebank/db/sqlc"
	"github.com/levietcuong2602/simplebank/validators"
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

	// register custom validators
	if v, ok := binding.Validator.Engine().(*validator.Validate); ok {
		v.RegisterValidation("currency", validators.ValidCurrency)
	}

	router.POST("/accounts", server.CreateAccount)
	router.GET("/accounts", server.GetListAccounts)
	router.GET("/accounts/:id", server.GetAccount)

	router.POST("/users", server.CreateUser)
	router.GET("/users/:username", server.GetUser)

	router.POST("/transfers", server.createTransfer)

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
