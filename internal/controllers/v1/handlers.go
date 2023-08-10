package v1

import (
	"github.com/gin-gonic/gin"
	"miniService/internal/repositories"
)

const (
	getItems = "/get-items"
)

type handler struct {
	repo repositories.Repo
}

func New(repo repositories.Repo) Handler {
	return &handler{repo: repo}
}

func (h *handler) Register() *gin.Engine {
	router := gin.Default()

	router.GET(getItems, h.getItems)

	return router
}
