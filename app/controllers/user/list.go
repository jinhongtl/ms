package user

import (
	"ms/app/models"
	"ms/internal"

	"github.com/gin-gonic/gin"
)

func List(c *gin.Context) {
	var users []models.User
	internal.Db.Order("id Desc").Find(&users, nil)
	internal.ReturnData(c, users)
}
