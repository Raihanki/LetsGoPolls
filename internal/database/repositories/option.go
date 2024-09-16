package repositories

import (
	"database/sql"

	"github.com/Raihanki/LetsGoPolls/internal/entities"
)

type OptionRepository struct {
	DB *sql.DB
}

func (repo *OptionRepository) GetAllOptions(pollId int) ([]entities.Option, error) {
	query := "SELECT id, poll_id, body, vote_count, created_at, updated_at FROM options WHERE poll_id = $1"
	rows, err := repo.DB.Query(query, pollId)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	options := []entities.Option{}
	for rows.Next() {
		option := entities.Option{}
		err := rows.Scan(&option.Id, &option.PollId, &option.Body, &option.VoteCount, &option.CreatedAt, &option.UpdatedAt)
		if err != nil {
			return nil, err
		}
		options = append(options, option)
	}

	return options, nil
}

func (repo *OptionRepository) CreateOption(request entities.CreateOptionRequest, pollId int) (entities.Option, error) {
	query := "INSERT INTO options (poll_id, body) VALUES ($1, $2) RETURNING *"
	option := entities.Option{}
	row := repo.DB.QueryRow(query, pollId, request.Body)
	err := row.Scan(&option.Id, &option.PollId, &option.Body, &option.VoteCount, &option.CreatedAt, &option.UpdatedAt)
	if err != nil {
		return option, err
	}

	return option, nil
}

func (repo *OptionRepository) GetOptionById(id int) (entities.Option, error) {
	query := "SELECT id, poll_id, body, vote_count, created_at, updated_at FROM options WHERE id = $1"
	row := repo.DB.QueryRow(query, id)

	option := entities.Option{}
	err := row.Scan(&option.Id, &option.PollId, &option.Body, &option.VoteCount, &option.CreatedAt, &option.UpdatedAt)
	if err != nil {
		return option, err
	}

	return option, nil
}

func (repo *OptionRepository) UpdateOption(request entities.UpdateOptionRequest, optionId int) (entities.Option, error) {
	query := "UPDATE options SET body = $1 WHERE id = $2 RETURNING *"
	option := entities.Option{}
	row := repo.DB.QueryRow(query, request.Body, optionId)
	err := row.Scan(&option.Id, &option.PollId, &option.Body, &option.VoteCount, &option.CreatedAt, &option.UpdatedAt)
	if err != nil {
		return option, err
	}

	return option, nil
}

func (repo *OptionRepository) DeleteOption(optionId int) error {
	query := "DELETE FROM options WHERE id = $1"
	_, err := repo.DB.Exec(query, optionId)
	if err != nil {
		return err
	}

	return nil
}
