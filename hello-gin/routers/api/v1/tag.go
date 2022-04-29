package v1

import (
	"github.com/gin-gonic/gin"
)

func GetTags(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "get tags",
	})
}

func AddTag(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "add tag",
	})
}

func EditTag(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "edit tag",
	})
}

func DeleteTag(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "delete tag",
	})
}
