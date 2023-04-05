package item_repository

import (
	"github.com/abilsabili50/Hacktiv-Assignment-2.git/entity"
	"github.com/abilsabili50/Hacktiv-Assignment-2.git/pkg/errs"
)

type ItemRepository interface {
	FindItemsByItemCode([]string, int) ([]*entity.Items, errs.MessageErr)
}
