package main

import (
	"database/sql"
	"gin-sample/pkg/setting"
	routers "gin-sample/routers"
	"log"

	_ "github.com/lib/pq"
)

func main() {

	setting.Setup()

	// TODO:接続文字列
	connStr := "host=db port=5432 user=myuser password=mypassword dbname=mydb sslmode=disable"

	// データベースに接続
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	// 接続テスト
	err = db.Ping()
	if err != nil {
		log.Fatal(err)
	}

	log.Println("データベースに正常に接続されました")

	router := routers.InitRouter(db)
	router.Run(":8080")
}
