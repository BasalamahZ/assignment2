package controller

import (
	dtoOrder "assignment2/pkg/order/dto"
	"assignment2/pkg/order/usecase"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type OrderHTTPController struct {
	usecaseOrder usecase.UsecaseInterfaceOrder
}

func InitControllerOrder(usecaseOrder usecase.UsecaseInterfaceOrder) *OrderHTTPController {
	return &OrderHTTPController{
		usecaseOrder: usecaseOrder,
	}
}

func (c *OrderHTTPController) GetAllOrder(ctx *gin.Context) {
	res, err := c.usecaseOrder.GetAllOrder()
	if err != nil {
		ctx.IndentedJSON(
			http.StatusInternalServerError, gin.H{
				"message": err.Error(),
				"status":  http.StatusInternalServerError,
			})
		return
	}
	ctx.IndentedJSON(http.StatusOK, gin.H{
		"message": "Success",
		"status":  "Success",
		"data":    res,
	})
}

func (c *OrderHTTPController) CreateOrder(ctx *gin.Context) {
	var input dtoOrder.OrderRequest
	if err := ctx.ShouldBindJSON(&input); err != nil {
		ctx.IndentedJSON(
			http.StatusBadRequest, gin.H{
				"status":  "Bad Request",
				"message": err.Error(),
			})
		return
	}

	order, err := c.usecaseOrder.CreateOrder(input)
	if err != nil {
		ctx.IndentedJSON(
			http.StatusInternalServerError, gin.H{
				"message": err.Error(),
				"status":  http.StatusInternalServerError,
			})
		return
	}

	ctx.IndentedJSON(http.StatusCreated, gin.H{
		"message": "Success",
		"status":  "Success",
		"data":    order,
	})
}

func (c *OrderHTTPController) UpdateOrder(ctx *gin.Context) {
	var input dtoOrder.OrderRequest
	idString := ctx.Param("id")
	if err := ctx.ShouldBindJSON(&input); err != nil {
		ctx.IndentedJSON(
			http.StatusBadRequest, gin.H{
				"status":  "Bad Request",
				"message": "title cannot be null",
			})
		return
	}

	id, err := strconv.Atoi(idString)
	if err != nil {
		ctx.IndentedJSON(http.StatusBadRequest, gin.H{
			"message": "Invalid Parsing Order Id",
			"status":  http.StatusBadRequest,
		})
		return
	}

	res, err := c.usecaseOrder.UpdateOrder(id, input)
	if err != nil {
		ctx.IndentedJSON(
			http.StatusInternalServerError, gin.H{
				"message": err.Error(),
				"status":  http.StatusInternalServerError,
			})
		return
	}

	ctx.IndentedJSON(http.StatusOK, gin.H{
		"message": "Success",
		"status":  "Success",
		"data":    res,
	})
}

func (c *OrderHTTPController) DeleteOrder(ctx *gin.Context) {
	idString := ctx.Param("id")

	id, err := strconv.Atoi(idString)
	if err != nil {
		ctx.IndentedJSON(http.StatusBadRequest, gin.H{
			"message": "Invalid Parsing Order Id",
			"status":  http.StatusBadRequest,
		})
		return
	}

	err = c.usecaseOrder.DeleteOrder(id)
	if err != nil {
		ctx.IndentedJSON(
			http.StatusBadRequest, gin.H{
				"message": err.Error(),
				"status":  http.StatusBadRequest,
			})
		return
	}

	ctx.IndentedJSON(http.StatusOK, gin.H{
		"message": "Success",
		"status":  "Success",
		"data":    gin.H{},
	})
}
