package entities

import "time"

type Vote struct {
	Id        int       `json:"id"`
	UserId    int       `json:"user_id"`
	PollId    int       `json:"poll_id"`
	OptionId  int       `json:"option_id"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

type CreateVoteRequest struct {
	UserId   int `json:"user_id"`
	PollId   int `json:"poll_id"`
	OptionId int `json:"option_id"`
}

type UpdateVoteRequest struct {
	OptionId int `json:"option_id"`
	UserId   int `json:"user_id"`
}
