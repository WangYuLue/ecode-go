package routers

import (
	"ecode/config"
	"ecode/middlewares/jwt"
	apis "ecode/routers/apis/v1"

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

	router.GET("/", apis.IndexAPI)

	v1 := router.Group("/v1")
	{
		v1.GET("/ping", apis.IndexAPI)
		// 通过邮箱激活用户
		v1.GET("confirm-email/:userid/:uuid", apis.ConfirmEmail)
		// 通过邮箱设置新密码
		v1.POST("reset-password", apis.ResetPassword)
		email := v1.Group("/email")
		{
			// 重新发送激活邮件
			email.POST("send-confirm/:userid", jwt.Auth(), apis.SendConfirmEmail)
			// 想指定邮箱发送设置新密码的链接
			email.POST("send-reset-password/:email", apis.SendResetPasswordEmail)
		}
		users := v1.Group("/users")
		{
			{
				// 注册
				users.POST("/register", apis.RegisterAPI)
				// 登录
				users.POST("/login", apis.Login)
				// 更新 token
				users.POST("/token", apis.UpdateToken)
			}
			token := users.Group("/", jwt.Auth())
			{
				token.GET("/", apis.GetUsersAPI)

				token.GET("/:userid", apis.GetUserAPI)

				token.GET("/:userid/cards", apis.GetCardsByUserID)
				// 修改用户信息，账户名 密码 等
				token.PUT("/:userid", apis.ModUserAPI)
				// 注销用户
				token.DELETE("/:userid", apis.DelUserAPI)
			}

		}
		cards := v1.Group("/cards")
		{
			token := cards.Group("/")
			{
				token.POST("/", apis.AddCardAPI)

				token.GET("/", apis.GetCardsAPI)

				token.GET("/:cardid", apis.GetCardAPI)

				token.PUT("/:cardid", apis.ModCardAPI)

				token.DELETE("/:cardid", apis.DelCardAPI)
			}

		}
	}

	return router
}
