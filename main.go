package main

import (
	"database/sql"
	"fmt"
	member "gin-sample/models"
	"log"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"

	_ "github.com/lib/pq"
)

func main() {

	// 接続文字列
	connStr := "host=localhost port=5432 user=myuser password=mypassword dbname=mydb sslmode=disable"

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

	fmt.Println("データベースに正常に接続されました")

	r := gin.Default()

	r.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "こんにちは、Gin!",
		})
	})

	members := initializeMembers()

	r.POST("/members", func(c *gin.Context) {

		fmt.Println(c.PostForm("name"))
		fmt.Println(c.PostForm("age"))
		fmt.Println(c.PostForm("sex"))

		member := member.Member{
			Name: c.PostForm("name"),
			Age:  convertToInt(c.PostForm("age")),
			Sex:  member.Sex(c.PostForm("sex")),
		}

		err := member.AddMember(db)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"error": "メンバーの作成に失敗しました",
			})
			return
		}

		c.JSON(http.StatusOK, gin.H{
			"message": "メンバーが作成されました",
		})
	})

	r.GET("/members", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"members": members,
		})
	})

	r.GET("/members/:id", func(c *gin.Context) {
		id := c.Param("id")
		index, err := strconv.Atoi(id)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": "Invalid ID",
			})
			return
		}
		if index < 0 || index >= len(members) {
			c.JSON(http.StatusNotFound, gin.H{
				"error": "Member not found",
			})
			return
		}
		c.JSON(http.StatusOK, gin.H{
			"member": members[index],
		})
	})

	r.Run(":8080")
}

func convertToInt(s string) int {
	i, _ := strconv.Atoi(s)
	return i
}

func initializeMembers() []member.Member {
	return []member.Member{
		*member.NewMember("山田", 20, member.Male),
		*member.NewMember("田中", 21, member.Female),
	}
}
