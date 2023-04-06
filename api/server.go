package api

import (
	"github.com/gin-gonic/gin"
	db "github.com/techschool/db/sqlc"
)

type Server struct {
	store  *db.Store
	router *gin.Engine
}

func NewServer(store *db.Store) *Server {
	server := &Server{store: store}
	router := gin.Default()

	router.GET("/",server.root)
	router.GET("/account/:id", server.getAccount)
	router.POST("/account/create", server.createAccount)
	router.GET("/accounts", server.listAccount)
	router.PUT("/account/update",server.updateAccount)
	router.DELETE("account/delete/:id",server.deleteAccount)
	router.PUT("account/add",server.addAccountBalance)

	server.router = router
	return server
}

// start runs server on {address}
func (server *Server) Start(address string) error {
	return server.router.Run(address)
}

func errorResponse(err error) gin.H {
	return gin.H{"error": err.Error()}
}