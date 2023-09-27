package dbrepo

import (
	"context"
	"time"

	"github.com/saquibhasan3/go-reservation/internal/models"
)

func (m *postgresDBRepo) AllUsers() bool {
	return true
}

func (m *postgresDBRepo) InsertBooking(res models.Booking) error {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()
	stmt := `INSERT INTO bookings (first_name, last_name, email, phone, start_date, end_date, room_id, adults, children, created_at, updated_at) VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9)`
	_, err := m.DB.ExecContext(ctx, stmt,
		res.FirstName,
		res.LastName,
		res.Email,
		res.Phone,
		res.StartDate,
		res.EndDate,
		res.RoomID,
		res.Adults,
		res.Children,
		time.Now(),
		time.Now(),
	)
	if err != nil {
		return err
	}
	return nil
}
