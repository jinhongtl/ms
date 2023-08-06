package user

import (
	"ms/app/models"
	"ms/internal"

	"github.com/gin-gonic/gin"
)

func Detail(c *gin.Context) {
	id := c.Param("id")
	var user models.User
	internal.Db.Where("id = ?", id).First(&user)
	internal.ReturnData(c, user)
}
