package routers

import (
	"ecode/middleware/jwt"
	apis "ecode/routers/apis/v1"

	"github.com/gin-gonic/gin"
)

// InitRouter -
func InitRouter() *gin.Engine {
	router := gin.Default()

	router.GET("/", apis.IndexAPI)

	v1 := router.Group("/v1")
	{
		// 登录
		v1.POST("/login", apis.Login)

		// 退出登录
		v1.DELETE("/logout", apis.Logout)
		users := v1.Group("/users")
		{
			// 注册用户
			users.POST("/", apis.AddUserAPI)

			users.GET("/", jwt.Auth(), apis.GetUsersAPI)

			users.GET("/:id", apis.GetUserAPI)

			users.GET("/:id/cards", apis.GetCardsByUserID)

			users.PUT("/:id", apis.ModUserAPI)
			// 注销用户
			users.DELETE("/:id", apis.DelUserAPI)
		}
		cards := v1.Group("/cards")
		{
			cards.POST("/", apis.AddCardAPI)

			cards.GET("/", apis.GetCardsAPI)

			cards.GET("/:id", apis.GetCardAPI)

			cards.PUT("/:id", apis.ModCardAPI)

			cards.DELETE("/:id", apis.DelCardAPI)
		}
	}

	return router
}
