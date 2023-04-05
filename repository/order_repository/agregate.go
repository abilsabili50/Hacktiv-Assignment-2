package order_repository

import "github.com/abilsabili50/Hacktiv-Assignment-2.git/entity"

type OrderItem struct {
	Order entity.Orders
	Items []entity.Items
}
