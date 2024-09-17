package entities

import "time"

type Poll struct {
	Id             int       `json:"id"`
	UserId         int       `json:"user_id"`
	Title          string    `json:"title"`
	Description    string    `json:"description"`
	EndDate        time.Time `json:"end_date"`
	IsActive       bool      `json:"is_active"`
	WinnerOptionId *int      `json:"winner_option_id"`
	CreatedAt      string    `json:"created_at"`
	UpdatedAt      string    `json:"updated_at"`
}

type customTime struct {
	time.Time
}

type CreatePollRequest struct {
	Title       string     `json:"title"`
	Description string     `json:"description"`
	EndDate     customTime `json:"end_date"`
	IsActive    bool       `json:"is_active"`
}

type UpdatePollRequest struct {
	Title       string     `json:"title"`
	Description string     `json:"description"`
	EndDate     customTime `json:"end_date"`
	IsActive    bool       `json:"is_active"`
}

func (c *customTime) UnmarshalJSON(b []byte) error {
	t, err := time.Parse(`"2006-01-02 15:04:05"`, string(b))
	if err != nil {
		return err
	}
	c.Time = t
	return nil
}
