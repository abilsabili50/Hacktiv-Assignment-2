package order_repository

import (
	"github.com/abilsabili50/Hacktiv-Assignment-2.git/entity"
	"github.com/abilsabili50/Hacktiv-Assignment-2.git/pkg/errs"
)

type OrderRepository interface {
	CreateOrder(orderPayload entity.Orders, itemsPayload []entity.Items) (*entity.Orders, errs.MessageErr)
	GetAllOrder() ([]entity.Orders, errs.MessageErr)
	GetOrderById(int) (*entity.Orders, errs.MessageErr)
	UpdateOrder(entity.Orders, []entity.Items) (*OrderItem, errs.MessageErr)
	DeleteOrder(int) errs.MessageErr
}
