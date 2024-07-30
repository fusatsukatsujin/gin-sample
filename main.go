package main

import (
	"gin-sample/member"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	r.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "こんにちは、Gin!",
		})
	})

	members := initializeMembers()

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

func initializeMembers() []member.Member {
	return []member.Member{
		*member.NewMember("山田", 20, member.Male),
		*member.NewMember("田中", 21, member.Female),
	}
}
