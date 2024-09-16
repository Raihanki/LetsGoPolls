package main

import (
	"net/http"

	"github.com/Raihanki/LetsGoPolls/cmd/api/handlers"
	"github.com/Raihanki/LetsGoPolls/internal/database/repositories"
)

func (cfg *ApplicationConfig) Routes() http.Handler {
	mux := http.NewServeMux()

	mux.HandleFunc("GET /api/v1/home", handlers.GetHome)

	//user routes
	userHandler := handlers.UserHandler{
		UserRepository: repositories.UserRepository{DB: cfg.DB},
	}
	mux.HandleFunc("POST /api/v1/users/register", userHandler.Register)
	mux.HandleFunc("POST /api/v1/users/login", userHandler.Login)

	//poll routes
	pollHandler := handlers.PollHandler{
		PollRepository: repositories.PollRepository{DB: cfg.DB},
	}
	mux.HandleFunc("POST /api/v1/polls", pollHandler.CreatePoll)
	mux.HandleFunc("GET /api/v1/polls/{pollId}", pollHandler.GetPollById)
	mux.HandleFunc("PUT /api/v1/polls/{pollId}", pollHandler.UpdatePoll)
	mux.HandleFunc("DELETE /api/v1/polls/{pollId}", pollHandler.DeletePoll)

	//option routes
	optionHandler := handlers.OptionHandler{
		OptionRepository: repositories.OptionRepository{DB: cfg.DB},
		PollRepository:   repositories.PollRepository{DB: cfg.DB},
	}
	mux.HandleFunc("GET /api/v1/polls/{pollId}/options", optionHandler.GetAllOptions)
	mux.HandleFunc("POST /api/v1/polls/{pollId}/options", optionHandler.CreateOption)
	mux.HandleFunc("GET /api/v1/polls/{pollId}/options/{optionId}", optionHandler.GetOptionById)
	mux.HandleFunc("PUT /api/v1/polls/{pollId}/options/{optionId}", optionHandler.UpdateOption)
	mux.HandleFunc("DELETE /api/v1/polls/{pollId}/options/{optionId}", optionHandler.DeleteOption)

	return mux
}
