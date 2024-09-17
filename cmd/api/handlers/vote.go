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

type VoteHandler struct {
	VoteRepository repositories.VoteRepository
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

	option, err := repo.VoteRepository.GetVoteByUserIdAndPollId(userId, pollId)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			helpers.JsonResponse(w, 404, "vote not found", nil)
			return
		}
		log.Printf("error while getting vote: %v", err)
		helpers.JsonResponse(w, 500, "", nil)
		return
	}

	request.UserId = userId
	vote, err := repo.VoteRepository.UpdateVote(request, voteId, option.Id)
	if err != nil {
		log.Printf("error while updating vote: %v", err)
		helpers.JsonResponse(w, 500, "", nil)
		return
	}

	helpers.JsonResponse(w, 200, http.StatusText(http.StatusOK), vote)
}
