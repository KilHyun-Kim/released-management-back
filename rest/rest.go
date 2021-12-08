package rest

import "github.com/gin-gonic/gin"

func RunAPI(address string) error {
	r := gin.Default()
	h, _ := NewHandler()

	// 기술 전체 조회
	r.GET("/techs", h.GetAllTech)

	// 특정 기술 조회
	// r.GET("/tech/:id", func(c *gin.Context) {

	// })

	// r.POST("/comment/register", func(c *gin.Context) {

	// })

	// r.DELETE("/comment/delete", func(c *gin.Context) {

	// })
	// r.POST("/comment/update", func(c *gin.Context) {

	// })

}
