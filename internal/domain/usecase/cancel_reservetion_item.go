package usecase

import (
	"context"

	"github.com/cripplemymind9/inventory-service/internal/domain/contract"
	"github.com/cripplemymind9/inventory-service/internal/domain/entity"
)

type CanelReservationItemUseCase struct {
	transactor contract.RepoTransactor
}

func NewCancelReservationItemUseCase(
	transactor contract.RepoTransactor,
) *CanelReservationItemUseCase {
	return &CanelReservationItemUseCase{
		transactor: transactor,
	}
}

func (cri *CanelReservationItemUseCase) CancelReservation(
	ctx context.Context,
	dto CancelReservationItemDTO,
) error {
	return cri.transactor.InTx(ctx, func(tx contract.TxRepo) error {
		item, err := tx.GetItemByProductID(ctx, dto.ProductID)
		if err != nil {
			return err
		}

		if item.ReservedQuantity < dto.Quantity {
			return entity.ErrNotEnoughReserved
		}

		return tx.CancelReservation(ctx, dto.ProductID, dto.Quantity)
	})
}
