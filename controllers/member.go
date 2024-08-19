package member_controller

import (
	"database/sql"
	"fmt"
	member "gin-sample/models"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func GetMembers(c *gin.Context, db *sql.DB) {
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
}

func GetMemberById(c *gin.Context, db *sql.DB) {
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
}

func AddMember(c *gin.Context, db *sql.DB) {
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
}

func convertToInt(s string) int {
	i, _ := strconv.Atoi(s)
	return i
}
