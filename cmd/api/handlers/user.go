package handlers

import (
	"encoding/json"
	"log"
	"net/http"
	"time"

	"github.com/Raihanki/LetsGoPolls/internal/database/repositories"
	"github.com/Raihanki/LetsGoPolls/internal/entities"
	"github.com/Raihanki/LetsGoPolls/internal/helpers"
	"golang.org/x/crypto/bcrypt"
)

type UserHandler struct {
	UserRepository repositories.UserRepository
}

func (repo *UserHandler) Register(w http.ResponseWriter, r *http.Request) {
	request := entities.CreateUserRequest{}
	err := json.NewDecoder(r.Body).Decode(&request)
	defer r.Body.Close()
	if err != nil {
		helpers.JsonResponse(w, 400, "invalid request body", nil)
		return
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(request.Password), bcrypt.DefaultCost)
	if err != nil {
		log.Printf("error while hashing password: %v", err)
		helpers.JsonResponse(w, 500, "", nil)
		return
	}

	request.Password = string(hashedPassword)
	user, err := repo.UserRepository.CreateUser(request)
	if err != nil {
		log.Printf("error while creating user: %v", err)
		helpers.JsonResponse(w, 500, "", nil)
		return
	}

	type response struct {
		Id    int    `json:"id"`
		Name  string `json:"name"`
		Email string `json:"email"`
	}
	responseUser := response{
		Id:    user.Id,
		Name:  user.Name,
		Email: user.Email,
	}
	helpers.JsonResponse(w, 201, http.StatusText(http.StatusCreated), responseUser)
}

func (repo *UserHandler) Login(w http.ResponseWriter, r *http.Request) {
	request := entities.LoginUserRequest{}
	err := json.NewDecoder(r.Body).Decode(&request)
	defer r.Body.Close()
	if err != nil {
		helpers.JsonResponse(w, 400, "invalid request body", nil)
		return
	}

	user, err := repo.UserRepository.GetUserByEmail(request.Email)
	if err != nil {
		log.Printf("error while getting user: %v", err)
		helpers.JsonResponse(w, 400, "invalid email or password", nil)
		return
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(request.Password))
	if err != nil {
		helpers.JsonResponse(w, 400, "invalid email or password", nil)
		return
	}

	type response struct {
		Token     string `json:"token"`
		ExpiresIn string `json:"expires_in"`
	}
	token, err := helpers.GenerateToken(user.Id)
	if err != nil {
		log.Printf("error while generating token: %v", err)
		helpers.JsonResponse(w, 500, "", nil)
		return
	}

	responseUser := response{
		Token:     token,
		ExpiresIn: time.Now().Add(time.Hour * 24).Format("2006-01-02 15:04:05"),
	}
	helpers.JsonResponse(w, 200, http.StatusText(http.StatusOK), responseUser)
}
