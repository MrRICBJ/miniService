package v1

import "github.com/gin-gonic/gin"

type Handler interface {
	Register() *gin.Engine
}
