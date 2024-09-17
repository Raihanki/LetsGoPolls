package repositories

import (
	"database/sql"

	"github.com/Raihanki/LetsGoPolls/internal/entities"
)

type VoteRepository struct {
	DB *sql.DB
}

func (repo *VoteRepository) CreateVote(request entities.CreateVoteRequest) (entities.Vote, error) {
	query := "INSERT INTO votes (user_id, poll_id, option_id) VALUES ($1, $2, $3) RETURNING *"
	vote := entities.Vote{}
	row := repo.DB.QueryRow(query, request.UserId, request.PollId, request.OptionId)
	err := row.Scan(&vote.Id, &vote.UserId, &vote.PollId, &vote.OptionId, &vote.CreatedAt)
	if err != nil {
		return vote, err
	}

	return vote, nil
}

func (repo *VoteRepository) UpdateVote(request entities.UpdateVoteRequest, voteId int) (entities.Vote, error) {
	query := "UPDATE votes SET option_id = $1 WHERE id = $2 AND user_id = $3 RETURNING *"
	vote := entities.Vote{}
	row := repo.DB.QueryRow(query, request.OptionId, voteId, request.UserId)
	err := row.Scan(&vote.Id, &vote.UserId, &vote.PollId, &vote.OptionId, &vote.CreatedAt)
	if err != nil {
		return vote, err
	}

	return vote, nil
}

func (repo *VoteRepository) GetVoteByUserIdAndPollId(userId int, pollId int) (entities.Vote, error) {
	query := "SELECT id, user_id, poll_id, option_id, created_at, updated_at FROM votes WHERE user_id = $1 AND poll_id = $2"
	row := repo.DB.QueryRow(query, userId, pollId)

	vote := entities.Vote{}
	err := row.Scan(&vote.Id, &vote.UserId, &vote.PollId, &vote.OptionId, &vote.CreatedAt, &vote.UpdatedAt)
	if err != nil {
		return vote, err
	}

	return vote, nil
}
