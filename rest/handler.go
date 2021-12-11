package rest

import (
	"fmt"
	"kilhyun-kim/released-back/dblayer"
	"net/http"

	"github.com/gin-gonic/gin"
)

type HandlerInterface interface {
	GetTech(c *gin.Context)
}

type Handler struct {
	db dblayer.DBLayer
}

func NewHandler() (*Handler, error) {
	// Handler 객체에 대한 포인터 생성
	db, err := dblayer.NewORM("mysql", "root:111111@tcp(127.0.0.1:3306)/test_schema")
	fmt.Println(db)
	if err != nil {
		return nil, err
	}
	return &Handler{
		db: db,
	}, nil
}

func (h *Handler) GetTech(c *gin.Context) {
	if h.db == nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "server database error"})
		return
	}
	tech, err := h.db.GetAllTech()

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	fmt.Printf("Found %d products\n", len(tech))
	c.JSON(http.StatusOK, tech)
}
