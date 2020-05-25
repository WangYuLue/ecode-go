package routers

import (
	"github.com/gin-gonic/gin"

	apis "ecode/routers/apis/v1"
)

// InitRouter -
func InitRouter() *gin.Engine {
	router := gin.Default()

	router.GET("/", apis.IndexAPI)

	v1 := router.Group("/v1")
	{
		users := v1.Group("/users")
		{
			users.POST("/", apis.AddUserAPI)

			users.GET("/", apis.GetUsersAPI)

			users.GET("/:id", apis.GetUserAPI)

			users.GET("/:id/cards", apis.GetCardsByUserID)

			users.PUT("/:id", apis.ModUserAPI)

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
