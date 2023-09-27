package repository

import "github.com/saquibhasan3/go-reservation/internal/models"

type DatabaseRepo interface {
	AllUsers() bool
	InsertBooking(res models.Booking) error
}
