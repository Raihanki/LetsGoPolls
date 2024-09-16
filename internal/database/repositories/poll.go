package repositories

import (
	"database/sql"

	"github.com/Raihanki/LetsGoPolls/internal/entities"
)

type PollRepository struct {
	DB *sql.DB
}

func (repo *PollRepository) CreatePoll(request entities.CreatePollRequest, userId int) (entities.Poll, error) {
	query := "INSERT INTO polls (user_id, title, description, end_date, is_active) VALUES ($1, $2, $3) RETURNING *"
	poll := entities.Poll{}
	row := repo.DB.QueryRow(query, userId, request.Title, request.Description, request.EndDate, request.IsActive)
	err := row.Scan(&poll.Id, &poll.UserId, &poll.Title, &poll.Description, &poll.EndDate, &poll.IsActive, &poll.WinnerOptionId, &poll.CreatedAt, &poll.UpdatedAt)
	if err != nil {
		return poll, err
	}

	return poll, nil
}

func (repo *PollRepository) GetPollById(id int) (entities.Poll, error) {
	query := "SELECT id, user_id, title, description, end_date, is_active, winner_option_id, created_at, updated_at FROM polls WHERE id = $1"
	row := repo.DB.QueryRow(query, id)

	poll := entities.Poll{}
	err := row.Scan(&poll.Id, &poll.UserId, &poll.Title, &poll.Description, &poll.EndDate, &poll.IsActive, &poll.WinnerOptionId, &poll.CreatedAt, &poll.UpdatedAt)
	if err != nil {
		return poll, err
	}

	return poll, nil
}

func (repo *PollRepository) GetPollsByUserId(userId int) ([]entities.Poll, error) {
	query := "SELECT id, user_id, title, description, end_date, is_active, winner_option_id, created_at, updated_at FROM polls WHERE user_id = $1"
	rows, err := repo.DB.Query(query, userId)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	polls := []entities.Poll{}
	for rows.Next() {
		poll := entities.Poll{}
		err := rows.Scan(&poll.Id, &poll.UserId, &poll.Title, &poll.Description, &poll.EndDate, &poll.IsActive, &poll.WinnerOptionId, &poll.CreatedAt, &poll.UpdatedAt)
		if err != nil {
			return nil, err
		}
		polls = append(polls, poll)
	}

	return polls, nil
}

func (repo *PollRepository) UpdatePoll(request entities.UpdatePollRequest, pollId int) (entities.Poll, error) {
	query := "UPDATE polls SET title = $1, description = $2, end_date = $3, is_active = $4 WHERE id = $5 RETURNING *"
	poll := entities.Poll{}
	row := repo.DB.QueryRow(query, request.Title, request.Description, request.EndDate, request.IsActive, pollId)
	err := row.Scan(&poll.Id, &poll.UserId, &poll.Title, &poll.Description, &poll.EndDate, &poll.IsActive, &poll.WinnerOptionId, &poll.CreatedAt, &poll.UpdatedAt)
	if err != nil {
		return poll, err
	}

	return poll, nil
}

func (repo *PollRepository) DeletePoll(pollId int) error {
	query := "DELETE FROM polls WHERE id = $1"
	_, err := repo.DB.Exec(query, pollId)
	if err != nil {
		return err
	}

	return nil
}
