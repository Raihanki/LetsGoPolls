package handlers

import (
	"net/http"

	"github.com/Raihanki/LetsGoPolls/internal/helpers"
)

func GetHome(w http.ResponseWriter, r *http.Request) {
	helpers.JsonResponse(w, 200, http.StatusText(http.StatusOK), nil)
}
