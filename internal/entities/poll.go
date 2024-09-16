package entities

import "time"

type Poll struct {
	Id             int       `json:"id"`
	UserId         int       `json:"user_id"`
	Title          string    `json:"title"`
	Description    string    `json:"description"`
	EndDate        time.Time `json:"end_date"`
	IsActive       bool      `json:"is_active"`
	WinnerOptionId int       `json:"winner_option_id"`
	CreatedAt      string    `json:"created_at"`
	UpdatedAt      string    `json:"updated_at"`
}

type CreatePollRequest struct {
	Title       string    `json:"title"`
	Description string    `json:"description"`
	EndDate     time.Time `json:"end_date"`
	IsActive    bool      `json:"is_active"`
}

type UpdatePollRequest struct {
	Title       string    `json:"title"`
	Description string    `json:"description"`
	EndDate     time.Time `json:"end_date"`
	IsActive    bool      `json:"is_active"`
}
