package server

import "github.com/cripplemymind9/inventory-service/internal/domain/usecase"

type Dependencies struct {
	reserveItemUseCase           *usecase.ReserveItemUseCase
	cancelReservationItemUseCase *usecase.CanelReservationItemUseCase
}

func NewDependencies(
	reserveItemUseCase *usecase.ReserveItemUseCase,
	cancelReservationItemUseCase *usecase.CanelReservationItemUseCase,
) *Dependencies {
	return &Dependencies{
		reserveItemUseCase:           reserveItemUseCase,
		cancelReservationItemUseCase: cancelReservationItemUseCase,
	}
}
