package rest

import (
	"kilhyun-kim/released-back/dblayer"
	"net/http"

	"github.com/gin-gonic/gin"
)

type HandlerInterface interface {
	GetAllTech(c *gin.Context)
}

type Handler struct {
	db dblayer.DBLayer
}

func NewHandler() (*Handler, error) {
	// Handler 객체에 대한 포인터 생성
	return new(Handler), nil
}

func (h *Handler) GetAllTech(c *gin.Context) {
	if h.db == nil {
		return
	}
	tech, err := h.db.GetAllTech()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
	}
	c.JSON(http.StatusOK, tech)
}
