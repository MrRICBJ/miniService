package v1

import (
	"context"
	"github.com/gin-gonic/gin"
	"net/http"
)

func (h *handler) getItems(c *gin.Context) {
	StrIDs := c.QueryArray("id")

	IntIDs, err := valid(StrIDs)
	if err != nil {
		c.JSON(http.StatusBadRequest, nil)
		return
	}

	resItems, err := h.repo.GetById(context.Background(), IntIDs)
	if err != nil {
		c.JSON(http.StatusNotFound, nil)
		return
	}

	c.JSON(http.StatusOK, resItems)
}
