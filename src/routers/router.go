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
		users := v1.Group("/users")
		{
			{
				// 注册
				users.POST("/register", apis.AddUserAPI)
				// 登录
				users.POST("/login", apis.Login)
				// 更新 token
				users.POST("/token", apis.UpdateToken)
				// 邮箱激活
				users.GET("/:id/email-confirm/:uuid", apis.EmailConfirm)
			}
			token := users.Group("/", jwt.Auth())
			{
				token.GET("/", apis.GetUsersAPI)

				token.GET("/:id", apis.GetUserAPI)

				token.GET("/:id/cards", apis.GetCardsByUserID)

				token.PUT("/:id", apis.ModUserAPI)
				// 注销用户
				token.DELETE("/:id", apis.DelUserAPI)
			}

		}
		cards := v1.Group("/cards")
		{
			token := cards.Group("/")
			{
				token.POST("/", apis.AddCardAPI)

				token.GET("/", apis.GetCardsAPI)

				token.GET("/:id", apis.GetCardAPI)

				token.PUT("/:id", apis.ModCardAPI)

				token.DELETE("/:id", apis.DelCardAPI)
			}

		}
	}

	return router
}
