package contract

import (
	"context"

	"github.com/cripplemymind9/inventory-service/internal/domain/entity"
)

type RepoTransactor interface {
	InTx(ctx context.Context, f func(tx TxRepo) error) error
}

type TxRepo interface {
	GetItemByProductID(ctx context.Context, id int64) (entity.Item, error)
	ReserveItem(ctx context.Context, productID, quantity int64) error
	CancelReservation(ctx context.Context, productID, quantity int64) error
}
