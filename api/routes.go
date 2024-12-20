package api

import (
	"github.com/gin-gonic/gin"
	"override/internal/services"
	"override/internal/utils"
)

func SetupRoutes(r *gin.Engine) {
	user := new(services.UserService)

	r.GET("/users", func(c *gin.Context) {
		users, err := user.GetAllUsers()
		if err != nil {
			c.JSON(200, utils.Error(500, err.Error()))
			return
		}
		c.JSON(200, utils.Success(users))
	})

	// 更多路由...
}
