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

type OptionHandler struct {
	OptionRepository repositories.OptionRepository
	PollRepository   repositories.PollRepository
}

func (repo *OptionHandler) GetAllOptions(w http.ResponseWriter, r *http.Request) {
	pathPoll := r.PathValue("pollId")
	pollId, err := strconv.Atoi(pathPoll)
	if err != nil {
		helpers.JsonResponse(w, 400, "invalid poll id", nil)
		return
	}

	options, err := repo.OptionRepository.GetAllOptions(pollId)
	if err != nil {
		log.Printf("error while getting options: %v", err)
		helpers.JsonResponse(w, 500, "", nil)
		return
	}

	helpers.JsonResponse(w, 200, http.StatusText(http.StatusOK), options)
}

func (repo *OptionHandler) CreateOption(w http.ResponseWriter, r *http.Request) {
	pathPoll := r.PathValue("pollId")
	pollId, err := strconv.Atoi(pathPoll)
	if err != nil {
		helpers.JsonResponse(w, 400, "invalid poll id", nil)
		return
	}

	request := entities.CreateOptionRequest{}
	err = json.NewDecoder(r.Body).Decode(&request)
	defer r.Body.Close()
	if err != nil {
		helpers.JsonResponse(w, 400, "invalid request body", nil)
		return
	}

	poll, err := repo.PollRepository.GetPollById(pollId)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			helpers.JsonResponse(w, 404, "poll not found", nil)
			return
		}
		helpers.JsonResponse(w, 500, "", nil)
		return
	}

	option, err := repo.OptionRepository.CreateOption(request, poll.Id) // 1 is a dummy poll ID (default for now)
	if err != nil {
		log.Printf("error while creating option: %v", err)
		helpers.JsonResponse(w, 500, "", nil)
		return
	}

	helpers.JsonResponse(w, 201, http.StatusText(http.StatusCreated), option)
}

func (repo *OptionHandler) GetOptionById(w http.ResponseWriter, r *http.Request) {
	pathOpt := r.PathValue("optionId")
	optionId, err := strconv.Atoi(pathOpt)
	if err != nil {
		helpers.JsonResponse(w, 400, "invalid option id", nil)
		return
	}

	option, err := repo.OptionRepository.GetOptionById(optionId)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			helpers.JsonResponse(w, 404, "option not found", nil)
			return
		}
		log.Printf("error while getting option: %v", err)
		helpers.JsonResponse(w, 500, "", nil)
		return
	}

	helpers.JsonResponse(w, 200, http.StatusText(http.StatusOK), option)
}

func (repo *OptionHandler) UpdateOption(w http.ResponseWriter, r *http.Request) {
	pathOpt := r.PathValue("optionId")
	optionId, err := strconv.Atoi(pathOpt)
	if err != nil {
		helpers.JsonResponse(w, 400, "invalid option id", nil)
		return
	}

	pathPoll := r.PathValue("pollId")
	pollId, err := strconv.Atoi(pathPoll)
	if err != nil {
		helpers.JsonResponse(w, 400, "invalid poll id", nil)
		return
	}

	request := entities.UpdateOptionRequest{}
	err = json.NewDecoder(r.Body).Decode(&request)
	defer r.Body.Close()
	if err != nil {
		helpers.JsonResponse(w, 400, "invalid request body", nil)
		return
	}

	_, err = repo.PollRepository.GetPollById(pollId)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			helpers.JsonResponse(w, 404, "poll not found", nil)
			return
		}
		helpers.JsonResponse(w, 500, "", nil)
		return
	}

	option, err := repo.OptionRepository.UpdateOption(request, optionId)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			helpers.JsonResponse(w, 404, "option not found", nil)
			return
		}
		log.Printf("error while updating option: %v", err)
		helpers.JsonResponse(w, 500, "", nil)
		return
	}

	helpers.JsonResponse(w, 200, http.StatusText(http.StatusOK), option)
}

func (repo *OptionHandler) DeleteOption(w http.ResponseWriter, r *http.Request) {
	pathPoll := r.PathValue("pollId")
	pollId, err := strconv.Atoi(pathPoll)
	if err != nil {
		helpers.JsonResponse(w, 400, "invalid poll id", nil)
		return
	}

	pathOpt := r.PathValue("optionId")
	optionId, err := strconv.Atoi(pathOpt)
	if err != nil {
		helpers.JsonResponse(w, 400, "invalid option id", nil)
		return
	}

	_, err = repo.PollRepository.GetPollById(pollId)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			helpers.JsonResponse(w, 404, "poll not found", nil)
			return
		}
		helpers.JsonResponse(w, 500, "", nil)
		return
	}

	option, err := repo.OptionRepository.GetOptionById(optionId)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			helpers.JsonResponse(w, 404, "option not found", nil)
			return
		}
		log.Printf("error while getting option: %v", err)
		helpers.JsonResponse(w, 500, "", nil)
		return
	}

	err = repo.OptionRepository.DeleteOption(optionId)
	if err != nil {
		log.Printf("error while deleting option: %v", err)
		helpers.JsonResponse(w, 500, "", nil)
		return
	}

	helpers.JsonResponse(w, 200, http.StatusText(http.StatusOK), option)
}
