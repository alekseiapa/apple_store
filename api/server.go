package api

import (
	db "github.com/alekseiapa/apple_store/db/sqlc"
	"github.com/gin-gonic/gin"
)

type Server struct {
	store  *db.Store
	router *gin.Engine
}

func NewServer(store *db.Store) *Server {
	server := &Server{store: store}
	router := gin.Default()

	router.POST("/api/users", server.createUser)
	router.GET("/api/users/:id", server.getUser)
	router.GET("'/api/users", server.listUser)
	router.PUT("/api/users/:id", server.updateUser)
	router.DELETE("/api/users/:id", server.deleteUser)

	server.router = router
	return server

	//
}

// Start runs the HTTP server on a specific address to start listening the api requests
func (server *Server) Start(address string) error {
	return server.router.Run(address)
}

func successDeleteResponse() gin.H {
	return gin.H{"success": "Deleted successfully"}
}

func notFoundResponse() gin.H {
	return gin.H{"error": "Not Found"}
}

func errorResponse(err error) gin.H {
	return gin.H{"error": err.Error()}
}
