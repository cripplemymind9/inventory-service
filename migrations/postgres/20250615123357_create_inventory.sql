-- +goose Up
-- +goose StatementBegin
CREATE TABLE inventory (
    product_id INTEGER PRIMARY KEY,
    total_quantity INTEGER NOT NULL DEFAULT 0,
    reserved_quantity INTEGER NOT NULL DEFAULT 0,
    available_quantity INTEGER NOT NULL DEFAULT 0,
    CONSTRAINT check_quantities CHECK (
        total_quantity >= 0 AND
        reserved_quantity >= 0 AND
        available_quantity >= 0 AND
        total_quantity >= reserved_quantity AND
        available_quantity = total_quantity - reserved_quantity
    )
);

CREATE INDEX idx_inventory_product_id ON inventory(product_id);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS inventory;
-- +goose StatementEnd
