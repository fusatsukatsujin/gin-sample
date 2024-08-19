package routers

import (
	"database/sql"
	"fmt"
	member "gin-sample/models"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func InitRouter(db *sql.DB) *gin.Engine {
	r := gin.Default()

	r.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "こんにちは、Gin!",
		})
	})

	r.POST("/members", func(c *gin.Context) {
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

		members, err := member.GetMembers(db)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"error": "メンバーの取得に失敗しました",
			})
			fmt.Println(err)
			return
		}

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
		member, err := member.GetMemberById(db, index)

		if err != nil {
			c.JSON(http.StatusNotFound, gin.H{
				"error": "Member not found",
			})
			return
		}
		c.JSON(http.StatusOK, gin.H{
			"member": member,
		})
	})
	return r
}

func convertToInt(s string) int {
	i, _ := strconv.Atoi(s)
	return i
}
