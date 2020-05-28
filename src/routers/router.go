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
		email := v1.Group("/email")
		{
			// 邮箱激活
			email.GET("/:userid/confirm-email/:uuid", apis.ConfirmEmail)
			// 重新发送激活邮件
			email.POST(":userid/send-confirm-email/", jwt.Auth(), apis.SendConfirmEmail)
		}
		users := v1.Group("/users")
		{
			{
				// 注册
				users.POST("/register", apis.AddUserAPI)
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
