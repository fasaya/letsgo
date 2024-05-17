package models

import "time"

type Variant struct {
	id         int
	name       string
	quantity   int
	product_id int
	created_at time.Time
	updated_at time.Time
}
