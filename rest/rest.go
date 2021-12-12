package rest

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

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
	r.GET("/techs", h.GetTech)
	// 특정 기술 조회
	r.GET("/version/:id", h.FindVersion)
	// r.GET("/version/:id", func(c *gin.Context) {

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

func MyCustomMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Set("v", "123")

		c.Next()

		status := c.Writer.Status()
		fmt.Println(status)
	}
}

func MyCustomLogger() gin.HandlerFunc {
	return func(c *gin.Context) {
		fmt.Println("******************")
		c.Next()
		fmt.Println("******************")
	}
}
