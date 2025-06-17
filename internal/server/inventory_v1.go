package server

import (
	"context"
	"errors"

	"github.com/cripplemymind9/inventory-service/internal/domain/entity"
	"github.com/cripplemymind9/inventory-service/internal/domain/usecase"
	"github.com/cripplemymind9/inventory-service/pkg/api/v1"
)

func (s *Server) ReserveItem(
	ctx context.Context,
	req *api.ReserveItemRequest,
) (*api.ReserveItemResponse, error) {
	err := s.dependencies.reserveItemUseCase.ReserveItem(ctx, usecase.ReserveItemDTO{
		ProductID: req.GetProductId(),
		Quantity:  int64(req.GetQuantity()),
	})

	if err != nil {
		if errors.Is(err, entity.ErrNotEnoughStock) {
			return &api.ReserveItemResponse{
				Status: api.ResponseStatus_INSUFFICIENT_QUANTITY,
			}, err
		}
		return &api.ReserveItemResponse{
			Status: api.ResponseStatus_INTERNAL_ERROR,
		}, err
	}

	return &api.ReserveItemResponse{
		Status: api.ResponseStatus_SUCCESS,
	}, nil
}

func (s *Server) CompensateItem(
	ctx context.Context,
	req *api.CompensateItemRequest,
) (*api.CompensateItemResponse, error) {
	err := s.dependencies.cancelReservationItemUseCase.CancelReservation(ctx, usecase.CancelReservationItemDTO{
		ProductID: req.GetProductId(),
		Quantity:  int64(req.GetQuantity()),
	})

	if err != nil {
		if errors.Is(err, entity.ErrNotEnoughReserved) {
			return &api.CompensateItemResponse{
				Status: api.ResponseStatus_INSUFFICIENT_QUANTITY,
			}, err
		}
		return &api.CompensateItemResponse{
			Status: api.ResponseStatus_INTERNAL_ERROR,
		}, err
	}

	return &api.CompensateItemResponse{
		Status: api.ResponseStatus_SUCCESS,
	}, nil
}
