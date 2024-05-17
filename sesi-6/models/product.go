package models

import "time"

const (
	Table = "products"
)

type Product struct {
	ID        int       `db:"id" json:"id"`
	Name      string    `db:"name" json:"name"`
	CreatedAt time.Time `db:"created_at" json:"created_at"`
	UpdatedAt time.Time `db:"update_at" json:"update_at"`

	Variants []Variant `db:"variants" json:"variants"`
}
