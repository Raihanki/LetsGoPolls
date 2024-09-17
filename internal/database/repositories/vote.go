package repositories

import (
	"database/sql"

	"github.com/Raihanki/LetsGoPolls/internal/entities"
)

type VoteRepository struct {
	DB *sql.DB
}

func (repo *VoteRepository) CreateVote(request entities.CreateVoteRequest) (entities.Vote, error) {
	tx, err := repo.DB.Begin()
	if err != nil {
		return entities.Vote{}, err
	}
	defer tx.Rollback()

	// lock options
	queryLockOptions := "SELECT id FROM options WHERE id = $1 FOR UPDATE"
	_, err = tx.Exec(queryLockOptions, request.OptionId)
	if err != nil {
		return entities.Vote{}, err
	}

	// check poll
	queryCheckPoll := "SELECT id FROM polls WHERE id = $1"
	rowPoll := tx.QueryRow(queryCheckPoll, request.PollId)
	var pollId int
	err = rowPoll.Scan(&pollId)
	if err != nil {
		return entities.Vote{}, err
	}

	query := "INSERT INTO votes (user_id, poll_id, option_id) VALUES ($1, $2, $3) RETURNING *"
	vote := entities.Vote{}
	row := tx.QueryRow(query, request.UserId, request.PollId, request.OptionId)
	err = row.Scan(&vote.Id, &vote.UserId, &vote.PollId, &vote.OptionId, &vote.CreatedAt, &vote.UpdatedAt)
	if err != nil {
		return vote, err
	}

	queryUpdateOptions := "UPDATE options SET vote_count = vote_count + 1 WHERE id = $1"
	_, err = tx.Exec(queryUpdateOptions, request.OptionId)
	if err != nil {
		return vote, err
	}

	err = tx.Commit()
	if err != nil {
		return vote, err
	}

	return vote, nil
}

func (repo *VoteRepository) UpdateVote(request entities.UpdateVoteRequest, voteId int, optionIdBefore int) (entities.Vote, error) {
	tx, err := repo.DB.Begin()
	if err != nil {
		return entities.Vote{}, err
	}
	defer tx.Rollback()

	// lock options
	queryLockOptions := "SELECT id FROM options WHERE id = $1 AND poll_id = (SELECT poll_id FROM votes WHERE id = $2) FOR UPDATE"
	_, err = tx.Exec(queryLockOptions, request.OptionId, voteId)
	if err != nil {
		return entities.Vote{}, err
	}

	query := "UPDATE votes SET option_id = $1 WHERE id = $2 AND user_id = $3 RETURNING *"
	vote := entities.Vote{}
	row := tx.QueryRow(query, request.OptionId, voteId, request.UserId)
	err = row.Scan(&vote.Id, &vote.UserId, &vote.PollId, &vote.OptionId, &vote.CreatedAt, &vote.UpdatedAt)
	if err != nil {
		return vote, err
	}

	queryUpdateOptionsAdd := "UPDATE options SET vote_count = vote_count + 1 WHERE id = $1"
	_, err = tx.Exec(queryUpdateOptionsAdd, request.OptionId)
	if err != nil {
		return vote, err
	}

	queryUpdateOptionsSubtract := "UPDATE options SET vote_count = vote_count - 1 WHERE id = $1"
	_, err = tx.Exec(queryUpdateOptionsSubtract, optionIdBefore)
	if err != nil {
		return vote, err
	}

	err = tx.Commit()
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
