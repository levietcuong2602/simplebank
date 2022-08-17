package api

import (
	"github.com/gin-gonic/gin"
)

type Server struct {
	store  *db.store
	router *gin.Engine
}
