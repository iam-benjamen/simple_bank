package api

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"github.com/go-playground/validator/v10"
	db "github.com/iam-benjamen/simple_bank/db/sqlc"
	"github.com/iam-benjamen/simple_bank/token"
	"github.com/iam-benjamen/simple_bank/util"
)

// Server serves HTTP requests for our banking service
type Server struct {
	store      *db.Store
	router     *gin.Engine
	config     util.Config
	tokenMaker token.Maker
}

// NewServer creates a new HTTP server and set up routing
func NewServer(config util.Config, store *db.Store) (*Server, error) {
	tokenMaker, err := token.NewJWTMaker(config.TokenSymmetricKey)
	if err != nil {
		return nil, fmt.Errorf("cannot create token maker: %w", err)
	}
	server := &Server{
		config:     config,
		store:      store,
		tokenMaker: tokenMaker,
	}

	if v, ok := binding.Validator.Engine().(*validator.Validate); ok {
		v.RegisterValidation("currency", validCurrency)
	}

	server.setUpRouter()
	return server, nil
}

// NewServer creates a new HTTP server and set up routing
func (server *Server) setUpRouter() {
	router := gin.Default()

	//add routes to router
	router.POST("/users", server.createAccount)
	router.POST("/users/login", server.loginUser)
	router.POST("/accounts", server.createAccount)
	router.GET("/accounts/:id", server.getAccount)
	router.GET("/accounts", server.listAccounts)
	router.POST("/transfers", server.createTransfer)

	server.router = router
}

// start runs the HTTP server on a specific address
func (server *Server) Start(address string) error {
	return server.router.Run(address)
}

func errorResponse(err error) gin.H {
	return gin.H{"error": err.Error()}
}
