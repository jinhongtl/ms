package user

import (
	"ms/app/models"
	"ms/internal"

	"github.com/gin-gonic/gin"
)

func Add(c *gin.Context) {

	var user models.User

	err := c.Bind(&user)

	if err != nil {
		internal.BadRequest(c)
	}

	err = internal.Db.Create(&user).Error

	if err != nil {
		internal.Logger.Error("添加用户出错：" + err.Error())
		internal.InternalError(c)
		return
	}

	if user.ID <= 0 {
		internal.Logger.Error("无法添加新用户：" + err.Error())
		internal.InternalError(c)
		return
	}
	internal.SimpleSuccess(c)
}
