package api

import (
	"fmt"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"github.com/go-playground/validator/v10"
	db "github.com/levietcuong2602/simplebank/db/sqlc"
	"github.com/levietcuong2602/simplebank/token"
	"github.com/levietcuong2602/simplebank/validators"
)

// Server serves HTTP requests for banking service.
type Server struct {
	store      *db.Store
	router     *gin.Engine
	tokenMaker token.Maker
}

// NewServer creates a new HTTP server and setup routing
func NewServer(store *db.Store) (*Server, error) {
	SYMMETRIC_SECRET_KEY := os.Getenv("SYMMETRIC_SECRET_KEY")
	fmt.Println("new server: ", SYMMETRIC_SECRET_KEY, len(SYMMETRIC_SECRET_KEY))
	tokenMaker, err := token.NewPasetoMaker(SYMMETRIC_SECRET_KEY)
	if err != nil {
		return nil, fmt.Errorf("Cannot Create Token Maker: %w", err)
	}
	server := &Server{store: store, tokenMaker: tokenMaker}
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
	router.POST("/login", server.LoginUser)

	router.POST("/transfers", server.createTransfer)

	server.router = router
	return server, nil
}

// Start run http server on specific address
func (server *Server) Start(addr string) error {
	return server.router.Run(addr)
}

func errorResponse(err error) gin.H {
	return gin.H{"error": err.Error()}
}
