package controllers

import (
	"database/sql"
	member "gin-sample/models"
	"gin-sample/services"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func GetMembers(c *gin.Context, db *sql.DB) {
	memberService := services.NewMemberService(db)
	members, err := memberService.GetMembers()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "メンバーの取得に失敗しました",
		})
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
	memberService := services.NewMemberService(db)
	member, err := memberService.GetMemberById(index)

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
	tx, exists := c.Get("tx")
	var txPtr *sql.Tx
	if exists {
		txPtr = tx.(*sql.Tx)
	}

	memberService := services.NewMemberService(db)
	err := memberService.AddMemberWithTransaction(
		txPtr,
		c.PostForm("name"),
		convertToInt(c.PostForm("age")),
		member.Sex(c.PostForm("sex")),
	)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "メンバーが作成されました"})
}

func convertToInt(s string) int {
	i, _ := strconv.Atoi(s)
	return i
}
