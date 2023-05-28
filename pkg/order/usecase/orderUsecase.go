package usecase

import (
	"assignment2/pkg/order/dto"
	"assignment2/pkg/order/model"
	"assignment2/pkg/order/repository"
	"fmt"
)

type UsecaseInterfaceOrder interface {
	GetAllOrder() (*[]model.Order, error)
	CreateOrder(input dto.OrderRequest) (model.Order, error)
	UpdateOrder(id int, input dto.OrderRequest) (*model.Order, error)
	DeleteOrder(id int) error
}

type usecaseOrder struct {
	repositoryOrder repository.RepositoryInterfaceOrder
}

func InitUsecaseOrder(repositoryOrder repository.RepositoryInterfaceOrder) UsecaseInterfaceOrder {
	return &usecaseOrder{
		repositoryOrder: repositoryOrder,
	}
}

// GetAllOrder implements UsecaseInterfaceOrder
func (u *usecaseOrder) GetAllOrder() (*[]model.Order, error) {
	order, err := u.repositoryOrder.GetAllOrder()
	if err != nil {
		return nil, err
	}

	return &order, nil
}

// CreateOrder implements UsecaseInterfaceOrder
func (u *usecaseOrder) CreateOrder(input dto.OrderRequest) (model.Order, error) {
	order := model.Order{
		OrderedAt:    input.OrderedAt,
		CostumerName: input.CostumerName,
	}

	itemsPayload := []model.Item{}

	for _, eachItem := range input.Item {
		item := model.Item{
			ItemCode:    eachItem.ItemCode,
			Quantity:    eachItem.Quantity,
			Description: eachItem.Description,
		}

		itemsPayload = append(itemsPayload, item)
	}
	newOrder, err := u.repositoryOrder.CreateOrder(order, itemsPayload)
	if err != nil {
		return newOrder, err
	}

	
	itemsResponse := []model.Item{}

	for _, eachItem := range itemsPayload {
		newItem := model.Item{
			ItemId:      eachItem.ItemId,
			OrderId:     int(newOrder.OrderId),
			ItemCode:    eachItem.ItemCode,
			Quantity:    eachItem.Quantity,
			Description: eachItem.Description,
		}
		itemsResponse = append(itemsResponse, newItem)
	}
	fmt.Println("item", itemsResponse)

	result := model.Order{
		OrderId:      newOrder.OrderId,
		CostumerName: newOrder.CostumerName,
		OrderedAt:    newOrder.OrderedAt,
		Items:        itemsPayload,
	}

	return result, nil

}

// UpdateOrder implements UsecaseInterfaceOrder
func (u *usecaseOrder) UpdateOrder(id int, input dto.OrderRequest) (*model.Order, error) {
	order := model.Order{
		OrderedAt:    input.OrderedAt,
		CostumerName: input.CostumerName,
	}

	itemsPayload := []model.Item{}

	for _, eachItem := range input.Item {
		item := model.Item{
			ItemCode:    eachItem.ItemCode,
			Quantity:    eachItem.Quantity,
			Description: eachItem.Description,
		}

		itemsPayload = append(itemsPayload, item)
	}

	newOrder, err := u.repositoryOrder.UpdateOrder(id, order, itemsPayload)
	if err != nil {
		return nil, err
	}

	itemsResponse := []model.Item{}

	for _, eachItem := range itemsPayload {
		newItem := model.Item{
			ItemId:      eachItem.ItemId,
			OrderId:     int(newOrder.OrderId),
			ItemCode:    eachItem.ItemCode,
			Quantity:    eachItem.Quantity,
			Description: eachItem.Description,
		}
		itemsResponse = append(itemsResponse, newItem)
	}

	result := model.Order{
		OrderId:      newOrder.OrderId,
		CostumerName: newOrder.CostumerName,
		OrderedAt:    newOrder.OrderedAt,
		Items:        itemsResponse,
	}

	return &result, nil
}

// DeleteOrder implements UsecaseInterfaceOrder
func (u *usecaseOrder) DeleteOrder(id int) error {
	return u.repositoryOrder.DeleteOrder(id)
}
