package repo

import (
	"context"
	"errors"
	"fmt"

	"github.com/cripplemymind9/inventory-service/internal/domain/entity"
	"github.com/jackc/pgx/v5"
)

func (q *queries) GetItemByProductID(ctx context.Context, id int64) (entity.Item, error) {
	const query = `
		SELECT
			product_id,
			total_quantity,
			reserved_quantity,
			available_quantity
		FROM inventory
		WHERE product_id = $1
	`

	row := q.db.QueryRow(ctx, query, id)

	var out entity.Item

	err := row.Scan(
		&out.ProductID,
		&out.TotalQuantity,
		&out.ReservedQuantity,
		&out.AvailableQuantity,
	)
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return entity.Item{}, entity.ErrItemNotFound
		}
		return entity.Item{}, fmt.Errorf("get item by productID storage err: %w", err)
	}

	return out, nil
}

func (q *queries) ReserveItem(ctx context.Context, productID, quantity int64) error {
	const query = `
		UPDATE inventory
		SET
			reserved_quantity = reserved_quantity + $1,
			available_quantity = available_quantity - $1
		WHERE
			product_id = $2
			AND available_quantity >= $1
	`

	tag, err := q.db.Exec(ctx, query, quantity, productID)
	if err != nil {
		return fmt.Errorf("reserve item storage err: %w", err)
	}

	if tag.RowsAffected() == 0 {
		return entity.ErrNotEnoughStock
	}

	return nil
}

func (q *queries) CancelReservation(ctx context.Context, productID, quantity int64) error {
	const query = `
		UPDATE inventory
		SET 
			reserved_quantity = reserved_quantity - $1,
			available_quantity = available_quantity + $1
		WHERE 
			product_id = $2
			AND reserved_quantity >= $1
	`

	tag, err := q.db.Exec(ctx, query, quantity, productID)
	if err != nil {
		return fmt.Errorf("cancel reservation storage err: %w", err)
	}

	if tag.RowsAffected() == 0 {
		return entity.ErrNotEnoughReserved
	}

	return nil
}
