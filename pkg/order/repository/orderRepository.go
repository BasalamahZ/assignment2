package repository

import (
	"assignment2/pkg/order/model"
	"fmt"

	"gorm.io/gorm"
)

type RepositoryInterfaceOrder interface {
	CreateOrder(order model.Order, item []model.Item) (model.Order, error)
	GetAllOrder() ([]model.Order, error)
	GetOrderById(id int) (model.Order, error)
	UpdateOrder(id int, order model.Order, item []model.Item) (model.Order, error)
	DeleteOrder(id int) error
}

type repositoryOrder struct {
	db *gorm.DB
}

func InitRepositoryOrder(db *gorm.DB) RepositoryInterfaceOrder {
	db.AutoMigrate(&model.Order{}, &model.Item{})
	// db.AutoMigrate(&model.Item{})
	return &repositoryOrder{
		db: db,
	}
}

func (r *repositoryOrder) CreateOrder(order model.Order, item []model.Item) (model.Order, error) {
	if err := r.db.Table("orders").Create(&order).Error; err != nil {
		return order, err
	}

	for _, eachItem := range item {
		if err := r.db.Table("items").Create(&model.Item{
			ItemCode:    eachItem.ItemCode,
			Description: eachItem.Description,
			Quantity:    eachItem.Quantity,
			OrderId:     int(order.OrderId),
		}).Error; err != nil {
			return order, err
		}
	}

	fmt.Println("item", item)
	return order, nil
}

func (r *repositoryOrder) GetAllOrder() ([]model.Order, error) {
	var orders []model.Order
	if err := r.db.Preload("Items").Find(&orders).Error; err != nil {
		return nil, err
	}

	return orders, nil
}

func (r *repositoryOrder) GetOrderById(id int) (model.Order, error) {
	var order model.Order
	if err := r.db.Preload("Items").Table("orders").Where("order_id = ?", id).First(&order).Error; err != nil {
		return order, err
	}

	return order, nil
}

func (r *repositoryOrder) UpdateOrder(id int, order model.Order, item []model.Item) (model.Order, error) {
	if err := r.db.Table("orders").Where("order_id = ?", id).Updates(&order).Error; err != nil {
		return order, err
	}

	for _, eachItem := range item {
		if err := r.db.Table("items").Where("order_id = ?", id).Updates(&model.Item{
			ItemCode:    eachItem.ItemCode,
			Description: eachItem.Description,
			Quantity:    eachItem.Quantity,
			OrderId:     int(order.OrderId),
		}).Error; err != nil {
			return order, err
		}
	}

	return order, nil
}

func (r *repositoryOrder) DeleteOrder(id int) error {
	if err := r.db.Table("orders").Where("order_id = ?", id).Delete(&model.Order{}).Error; err != nil {
		return err
	}

	return nil
}
