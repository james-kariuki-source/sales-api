package models

import (
	"time"
)

type Cashier struct {
	Id        uint      `json:"id"`
	Name      string    `json:"name"`
	Passcode  string    `json:"passcode"`
	CreateAt  time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}
