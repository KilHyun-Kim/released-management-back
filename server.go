package main

import (
	"database/sql"
	"fmt"
	"kilhyun-kim/released-back/crawler"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
)

func main() {

	go func() {
		crawler.Crawler()
	}()

	db, err := sql.Open("mysql", "root:pwd@tcp(127.0.0.1:3306)/gotest")
	if err != nil {
		log.Fatal(err)
	}

	defer db.Close()
	var name string
	err = db.QueryRow("SELECT title FROM topic ").Scan(&name)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(name)
	r := gin.Default()
	r.POST("/api/test", func(c *gin.Context) {
		c.JSON(http.StatusOK, name)
	})
	r.Run()
}
