package order

import (
	"assignment2/pkg/order/controller"
	"assignment2/pkg/order/repository"
	"assignment2/pkg/order/usecase"

	"gorm.io/gorm"
)

func InitHttpOrderController(db *gorm.DB) *controller.OrderHTTPController {
	repo := repository.InitRepositoryOrder(db)

	usecase := usecase.InitUsecaseOrder(repo)

	controller := controller.InitControllerOrder(usecase)

	return controller
}
