package usecase

type ReserveItemDTO struct {
	ProductID int64
	Quantity  int64
}

type CancelReservationItemDTO struct {
	ProductID int64
	Quantity  int64
}
