package router

import (
	"github.com/gin-gonic/gin"

	"yuan-api/user-web/api"
)

func InitUserRouter(Router *gin.RouterGroup) {
	UserRouter := Router.Group("user")
	{
		UserRouter.GET("list", api.GetUserList)

	}
}
