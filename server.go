package main

import (
	"kilhyun-kim/released-back/rest"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

func main() {

	// go func() {
	// 	crawler.Crawler()
	// }()

	// db, err := sql.Open("mysql", "root:pwd@tcp(127.0.0.1:3306)/gotest")
	// if err != nil {
	// 	log.Fatal(err)
	// }

	// defer db.Close()
	// var name string
	// err = db.QueryRow("SELECT title FROM topic ").Scan(&name)
	// if err != nil {
	// 	log.Fatal(err)
	// }

	log.Println("Main log...")
	log.Fatal(rest.RunAPI("127.0.0.1:8000"))

}
