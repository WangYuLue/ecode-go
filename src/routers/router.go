package routers

import (
	"ecode/config"
	"ecode/middlewares/jwt"
	APIs "ecode/routers/apis/v1"
	CardAPIs "ecode/routers/apis/v1/card"
	UserAPIs "ecode/routers/apis/v1/user"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

// InitRouter -
func InitRouter() *gin.Engine {
	router := gin.Default()
	corsConfig := cors.DefaultConfig()
	corsConfig.AllowOrigins = []string{config.FrontendURL}
	corsConfig.AllowHeaders = []string{"Authorization", "Content-Type"}

	router.Use(cors.New(corsConfig))
	// router.Use(cors.Default())

	router.GET("/", APIs.Ping)

	v1 := router.Group("/v1")
	{
		v1.GET("/ping", APIs.Ping)
		// 通过邮箱激活用户
		v1.GET("confirm-email/:userid/:uuid", APIs.ConfirmEmail)
		// 通过邮箱设置新密码
		v1.POST("reset-password", APIs.ResetPassword)
		email := v1.Group("/email")
		{
			// 重新发送激活邮件
			email.POST("send-confirm/:userid", jwt.Auth(), APIs.SendConfirmEmail)
			// 想指定邮箱发送设置新密码的链接
			email.POST("send-reset-password/:email", APIs.SendResetPasswordEmail)
		}

		// 注册
		v1.POST("/register", UserAPIs.Register)
		// 登录
		v1.POST("/login", APIs.Login)
		// 更新 token
		v1.POST("/token", APIs.UpdateToken)
		users := v1.Group("/users")
		{
			token := users.Group("/", jwt.Auth())
			{
				token.GET("/", UserAPIs.GetUsers)

				token.GET("/:userid", UserAPIs.GetUser)

				// 修改用户信息，账户名 密码 等
				token.PUT("/:userid", UserAPIs.ModUser)
				// 注销用户
				token.DELETE("/:userid", UserAPIs.DelUser)

				cards := token.Group("/:userid/cards")
				{
					cards.GET("/", UserAPIs.GetCards)

					cards.POST("/", UserAPIs.AddCard)

					cards.GET("/:cardid", UserAPIs.GetCard)

					cards.PUT("/:cardid", UserAPIs.ModCard)

					cards.DELETE("/:cardid", UserAPIs.DelCard)
				}
			}
		}
		cards := v1.Group("/cards")
		{
			cards.GET("/", CardAPIs.GetCards)

			cards.GET("/:cardid", CardAPIs.GetCard)
		}
	}

	return router
}
