package main

import (
	"time"
)

type Transaction struct {
	ID string `db:"id" json:"id"`
	Price float64 `db:"price" json:"price"`
	Status string `db:"status" json:"status"`
	CreateTime time.Time `db:"create_time" json:"create_time,omitempty"`
}

type Event struct {
	ID string `db:"id" json:"id"`
	EventType string `db:"event_type" json:"event_type"`
	Payload []byte `db:"payload" json:"payload"`
	CreateTime time.Time `db:"create_time" json:"create_time,omitempty"`
}
