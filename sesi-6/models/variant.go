package models

import "time"

type Variant struct {
	ID        int       `db:"id"`
	Name      string    `db:"name"`
	Quantity  int       `db:"quantity"`
	ProductID int       `db:"product_id"`
	CreatedAt time.Time `db:"created_at"`
	UpdatedAt time.Time `db:"updated_at"`
}
