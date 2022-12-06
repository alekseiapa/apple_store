package api

import (
	"fmt"

	db "github.com/alekseiapa/apple_store/db/sqlc"
	"github.com/alekseiapa/apple_store/token"
	"github.com/alekseiapa/apple_store/util"
	"github.com/gin-gonic/gin"
)

type Server struct {
	config     util.Config
	store      db.Store
	tokenMaker token.Maker
	router     *gin.Engine
}

func NewServer(config util.Config, store db.Store) (*Server, error) {
	tokenMaker, err := token.NewPasetoMaker(config.TokenSymmetricKey)
	if err != nil {
		return nil, fmt.Errorf("error creating token maker: %v", err)
	}
	server := &Server{
		config:     config,
		store:      store,
		tokenMaker: tokenMaker,
	}
	server.setupRouter()
	return server, nil
}

func (server *Server) setupRouter() {
	router := gin.Default()

	// TODO: add protected routes
	// authRoutes := router.Group("/api").Use(authMiddleware(server.tokenMaker))

	router.POST("/api/users", server.createUser)
	router.POST("/api/users/login", server.loginUser)
	router.GET("/api/users/:id", server.getUser)
	router.GET("/api/users", server.listUser)
	router.PUT("/api/users/:id", server.updateUser)
	router.DELETE("/api/users/:id", server.deleteUser)

	router.POST("/api/products", server.createProduct)
	router.GET("/api/products/:id", server.getProduct)
	router.GET("/api/products", server.listProduct)
	router.PUT("/api/products/:id", server.updateProduct)
	router.DELETE("/api/products/:id", server.deleteProduct)

	router.GET("/api/orders/:id", server.getOrder)
	router.POST("/api/orders", server.createOrder)
	router.DELETE("/api/orders/:id", server.deleteOrder)

	// TODO: The following routes should be implemented
	// router.GET("/api/orders/:id", server.getProduct)
	// router.GET("/api/orders", server.listProduct)
	// router.PUT("/api/orders/:id", server.updateProduct)

	server.router = router
}

// Start runs the HTTP server on a specific address to start listening the api requests
func (server *Server) Start(address string) error {
	return server.router.Run(address)
}

func successDeleteResponse() gin.H {
	return gin.H{"success": "Deleted successfully"}
}

func notFoundResponse(obj string) gin.H {
	return gin.H{"error": fmt.Sprintf("%s not found", obj)}
}

func errorResponse(err error) gin.H {
	return gin.H{"error": err.Error()}
}

func userNameExistsResponse() gin.H {
	return gin.H{"error": "Username already exists"}
}
