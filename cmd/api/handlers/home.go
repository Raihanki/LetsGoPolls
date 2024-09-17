package handlers

import (
	"net/http"
	"os"

	"github.com/Raihanki/LetsGoPolls/internal/helpers"
)

func GetHome(w http.ResponseWriter, r *http.Request) {
	type AppInfo struct {
		Name         string `json:"name"`
		Version      string `json:"version"`
		ServerHealth string `json:"server_health"`
	}

	ai := AppInfo{
		Name:         os.Getenv("APP_NAME"),
		Version:      os.Getenv("APP_VERSION"),
		ServerHealth: "OK",
	}
	helpers.JsonResponse(w, 200, http.StatusText(http.StatusOK), ai)
}
