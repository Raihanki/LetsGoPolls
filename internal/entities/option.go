package entities

type Option struct {
	Id        int    `json:"id"`
	PollId    int    `json:"poll_id"`
	Body      string `json:"body"`
	VoteCount int    `json:"vote_count"`
	CreatedAt string `json:"created_at"`
	UpdatedAt string `json:"updated_at"`
}

type CreateOptionRequest struct {
	Body string `json:"body"`
}

type UpdateOptionRequest struct {
	Body string `json:"body"`
}
