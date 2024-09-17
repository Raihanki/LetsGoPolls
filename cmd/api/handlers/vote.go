package handlers

import (
	"database/sql"
	"encoding/json"
	"errors"
	"log"
	"net/http"
	"strconv"
	"time"

	"github.com/Raihanki/LetsGoPolls/internal/database/repositories"
	"github.com/Raihanki/LetsGoPolls/internal/entities"
	"github.com/Raihanki/LetsGoPolls/internal/helpers"
)

type VoteHandler struct {
	VoteRepository   repositories.VoteRepository
	OptionRepository repositories.OptionRepository
	PollRepository   repositories.PollRepository
}

func (repo *VoteHandler) Store(w http.ResponseWriter, r *http.Request, userId int) {
	paramPoll := r.PathValue("pollId")
	pollId, err := strconv.Atoi(paramPoll)
	if err != nil {
		helpers.JsonResponse(w, 400, "invalid poll id", nil)
		return
	}

	request := entities.CreateVoteRequest{}
	err = json.NewDecoder(r.Body).Decode(&request)
	defer r.Body.Close()
	if err != nil {
		helpers.JsonResponse(w, 400, "invalid request body", nil)
		return
	}

	checkPoll, err := repo.PollRepository.GetPollById(pollId)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			helpers.JsonResponse(w, 404, "poll not found", nil)
			return
		}
		log.Printf("error while getting poll: %v", err)
		helpers.JsonResponse(w, 500, "", nil)
		return
	}

	if time.Now().After(checkPoll.EndDate) {
		helpers.JsonResponse(w, 400, "poll has ended", nil)
		return
	}

	checkOption, err := repo.OptionRepository.GetOptionById(request.OptionId)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			helpers.JsonResponse(w, 404, "option not found", nil)
			return
		}
		log.Printf("error while getting poll: %v", err)
		helpers.JsonResponse(w, 500, "", nil)
		return
	}

	//check if user already voted
	_, err = repo.VoteRepository.GetVoteByUserIdAndPollId(userId, pollId)
	if err == nil {
		helpers.JsonResponse(w, 400, "user already voted", nil)
		return
	}

	if checkOption.PollId != pollId {
		helpers.JsonResponse(w, 400, "option not found in this poll", nil)
		return
	}

	request.UserId = userId
	request.PollId = pollId
	vote, err := repo.VoteRepository.CreateVote(request)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			helpers.JsonResponse(w, 400, "invalid poll id", nil)
			return
		}
		log.Printf("error while creating vote: %v", err)
		helpers.JsonResponse(w, 500, "", nil)
		return
	}

	helpers.JsonResponse(w, 201, http.StatusText(http.StatusOK), vote)
}

func (repo *VoteHandler) Update(w http.ResponseWriter, r *http.Request, userId int) {
	paramPoll := r.PathValue("pollId")
	pollId, err := strconv.Atoi(paramPoll)
	if err != nil {
		helpers.JsonResponse(w, 400, "invalid poll id", nil)
		return
	}

	paramVote := r.PathValue("voteId")
	voteId, err := strconv.Atoi(paramVote)
	if err != nil {
		helpers.JsonResponse(w, 400, "invalid vote id", nil)
		return
	}

	request := entities.UpdateVoteRequest{}
	err = json.NewDecoder(r.Body).Decode(&request)
	defer r.Body.Close()
	if err != nil {
		helpers.JsonResponse(w, 400, "invalid request body", nil)
		return
	}

	checkPoll, err := repo.PollRepository.GetPollById(pollId)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			helpers.JsonResponse(w, 404, "poll not found", nil)
			return
		}
		log.Printf("error while getting poll: %v", err)
		helpers.JsonResponse(w, 500, "", nil)
		return
	}

	if time.Now().After(checkPoll.EndDate) {
		helpers.JsonResponse(w, 400, "poll has ended", nil)
		return
	}

	checkOption, err := repo.OptionRepository.GetOptionById(request.OptionId)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			helpers.JsonResponse(w, 404, "option not found", nil)
			return
		}
		log.Printf("error while getting poll: %v", err)
		helpers.JsonResponse(w, 500, "", nil)
		return
	}

	vote, err := repo.VoteRepository.GetVoteByUserIdAndPollId(userId, pollId)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			helpers.JsonResponse(w, 404, "vote not found", nil)
			return
		}
		log.Printf("error while getting vote: %v", err)
		helpers.JsonResponse(w, 500, "", nil)
		return
	}

	if request.OptionId == vote.OptionId {
		helpers.JsonResponse(w, 400, "same option", nil)
		return
	}

	if checkOption.PollId != pollId {
		helpers.JsonResponse(w, 400, "option not found in this poll", nil)
		return
	}

	request.UserId = userId
	updatedVote, err := repo.VoteRepository.UpdateVote(request, voteId, vote.OptionId)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			helpers.JsonResponse(w, 404, "vote not found", nil)
			return
		}
		log.Printf("error while updating vote: %v", err)
		helpers.JsonResponse(w, 500, "", nil)
		return
	}

	helpers.JsonResponse(w, 200, http.StatusText(http.StatusOK), updatedVote)
}

func (repo *VoteHandler) Show(w http.ResponseWriter, r *http.Request, userId int) {
	paramPoll := r.PathValue("pollId")
	pollId, err := strconv.Atoi(paramPoll)
	if err != nil {
		helpers.JsonResponse(w, 400, "invalid poll id", nil)
		return
	}

	vote, err := repo.VoteRepository.GetVoteByUserIdAndPollId(userId, pollId)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			helpers.JsonResponse(w, 404, "vote not found", nil)
			return
		}
		log.Printf("error while getting vote: %v", err)
		helpers.JsonResponse(w, 500, "", nil)
		return
	}

	helpers.JsonResponse(w, 200, http.StatusText(http.StatusOK), vote)
}
