package routes

import (
	"assignment2/pkg/order"
	"net/http"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func InitHtttpRoute(g *gin.Engine, db *gorm.DB) {
	pingGroup := g.Group("ping")

	pingGroup.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "pong",
		})
	})

	// Order Routes
	orderGroup := g.Group("order")
	orderController := order.InitHttpOrderController(db)
	orderGroup.GET("", orderController.GetAllOrder)
	{
		orderGroup.POST("", orderController.CreateOrder)
		orderGroup.PATCH("/:id", orderController.UpdateOrder)
		orderGroup.DELETE("/:id", orderController.DeleteOrder)
	}
}
