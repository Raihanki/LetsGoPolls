package handlers

import (
	"database/sql"
	"encoding/json"
	"errors"
	"log"
	"net/http"
	"strconv"

	"github.com/Raihanki/LetsGoPolls/internal/database/repositories"
	"github.com/Raihanki/LetsGoPolls/internal/entities"
	"github.com/Raihanki/LetsGoPolls/internal/helpers"
)

type PollHandler struct {
	PollRepository repositories.PollRepository
}

func (repo *PollHandler) CreatePoll(w http.ResponseWriter, r *http.Request, userId int) {
	request := entities.CreatePollRequest{}
	err := json.NewDecoder(r.Body).Decode(&request)
	defer r.Body.Close()
	if err != nil {
		helpers.JsonResponse(w, 400, "invalid request body", nil)
		return
	}

	poll, err := repo.PollRepository.CreatePoll(request, userId)
	if err != nil {
		log.Printf("error while creating poll: %v", err)
		helpers.JsonResponse(w, 500, "", nil)
		return
	}

	helpers.JsonResponse(w, 201, http.StatusText(http.StatusOK), poll)
}

func (repo *PollHandler) GetPollById(w http.ResponseWriter, r *http.Request) {
	pathV := r.PathValue("pollId")
	pollId, err := strconv.Atoi(pathV)
	if err != nil {
		helpers.JsonResponse(w, 400, "invalid poll id", nil)
		return
	}

	poll, err := repo.PollRepository.GetPollById(pollId)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			helpers.JsonResponse(w, 404, "poll not found", nil)
			return
		}
		log.Printf("error while getting poll: %v", err)
		helpers.JsonResponse(w, 500, "", nil)
		return
	}

	helpers.JsonResponse(w, 200, http.StatusText(http.StatusOK), poll)
}

func (repo *PollHandler) UpdatePoll(w http.ResponseWriter, r *http.Request, userId int) {
	pathV := r.PathValue("pollId")
	pollId, err := strconv.Atoi(pathV)
	if err != nil {
		helpers.JsonResponse(w, 400, "invalid poll id", nil)
		return
	}

	request := entities.UpdatePollRequest{}
	err = json.NewDecoder(r.Body).Decode(&request)
	defer r.Body.Close()
	if err != nil {
		helpers.JsonResponse(w, 400, "invalid request body", nil)
		return
	}

	poll, err := repo.PollRepository.UpdatePoll(request, pollId, userId)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			helpers.JsonResponse(w, 404, "poll not found", nil)
			return
		}
		log.Printf("error while updating poll: %v", err)
		helpers.JsonResponse(w, 500, "", nil)
		return
	}

	helpers.JsonResponse(w, 200, http.StatusText(http.StatusOK), poll)
}

func (repo *PollHandler) DeletePoll(w http.ResponseWriter, r *http.Request, userId int) {
	pathV := r.PathValue("pollId")
	pollId, err := strconv.Atoi(pathV)
	if err != nil {
		helpers.JsonResponse(w, 400, "invalid poll id", nil)
		return
	}

	_, err = repo.PollRepository.GetPollById(pollId)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			helpers.JsonResponse(w, 404, "poll not found", nil)
			return
		}
		log.Printf("error while getting poll: %v", err)
		helpers.JsonResponse(w, 500, "", nil)
		return
	}

	err = repo.PollRepository.DeletePoll(pollId, userId)
	if err != nil {
		log.Printf("error while deleting poll: %v", err)
		helpers.JsonResponse(w, 500, "", nil)
		return
	}

	helpers.JsonResponse(w, 204, http.StatusText(http.StatusNoContent), nil)
}
