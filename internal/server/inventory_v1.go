package server

import (
	"context"

	"github.com/cripplemymind9/inventory-service/internal/domain/entity"
	"github.com/cripplemymind9/inventory-service/internal/domain/usecase"
	"github.com/cripplemymind9/inventory-service/pkg/api/v1"
)

func (s *Server) ReserveItem(
	ctx context.Context,
	req *api.ReserveItemRequest,
) (*api.ReserveItemResponse, error) {
	err := s.dependencies.reserveItemUseCase.ReserveItem(ctx, usecase.ReserveItemDTO{
		ProductID: req.ProductId,
		Quantity:  int64(req.Quantity),
	})

	if err != nil {
		switch err {
		case entity.ErrNotEnoughStock:
			return &api.ReserveItemResponse{
				Status: api.ResponseStatus_INSUFFICIENT_QUANTITY,
			}, nil
		default:
			return &api.ReserveItemResponse{
				Status: api.ResponseStatus_INTERNAL_ERROR,
			}, err
		}
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
		ProductID: req.ProductId,
		Quantity:  int64(req.Quantity),
	})

	if err != nil {
		switch err {
		case entity.ErrNotEnoughReserved:
			return &api.CompensateItemResponse{
				Status: api.ResponseStatus_INSUFFICIENT_QUANTITY,
			}, nil
		default:
			return &api.CompensateItemResponse{
				Status: api.ResponseStatus_INTERNAL_ERROR,
			}, err
		}
	}

	return &api.CompensateItemResponse{
		Status: api.ResponseStatus_SUCCESS,
	}, nil
}
