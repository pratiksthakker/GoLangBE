package handler

import (
	"github.com/gin-gonic/gin"
)

func GetUserDetails(c *gin.Context) {

	c.BindJSON("userID")

}
