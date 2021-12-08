package rest

import "github.com/gin-gonic/gin"

func RunAPI(address string) error {
	h, err := NewHandler()
	if err != nil {
		return err
	}
	return RunAPIWithHandler(address, h)
}

func RunAPIWithHandler(address string, h HandlerInterface) error {
	r := gin.Default()

	// 기술 전체 조회
	r.GET("/techs", h.GetAllTech)

	// 특정 기술 조회
	// r.GET("/tech/:id", func(c *gin.Context) {

	// })

	// commentGroup := r.Group("/comment")
	// {
	// 	commentGroup.POST("/comment/register")
	// 	commentGroup.POST("/comment/register")
	// 	commentGroup.POST("/comment/register")
	// }
	// r.POST("/comment/register", func(c *gin.Context) {

	// })

	// r.DELETE("/comment/delete", func(c *gin.Context) {

	// })
	// r.POST("/comment/update", func(c *gin.Context) {

	// })
	return r.Run(address)
}
