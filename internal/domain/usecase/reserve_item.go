package usecase

import (
	"context"

	"github.com/cripplemymind9/inventory-service/internal/domain/contract"
	"github.com/cripplemymind9/inventory-service/internal/domain/entity"
)

type ReserveItemUseCase struct {
	transactor contract.RepoTransactor
}

func NewReserveItemUseCase(
	transactor contract.RepoTransactor,
) *ReserveItemUseCase {
	return &ReserveItemUseCase{
		transactor: transactor,
	}
}

func (ri *ReserveItemUseCase) ReserveItem(
	ctx context.Context,
	dto ReserveItemDTO,
) error {
	return ri.transactor.InTx(ctx, func(tx contract.TxRepo) error {
		item, err := tx.GetItemByProductID(ctx, dto.ProductID)
		if err != nil {
			return err
		}

		if item.AvailableQuantity < dto.Quantity {
			return entity.ErrNotEnoughStock
		}

		return tx.ReserveItem(ctx, dto.ProductID, dto.Quantity)
	})
}
