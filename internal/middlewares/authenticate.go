package middlewares

import (
	"net/http"
	"strconv"
	"strings"

	"github.com/Raihanki/LetsGoPolls/internal/helpers"
)

type AuthenticationMiddlewareFunction func(w http.ResponseWriter, r *http.Request, userId int)

func AuthenticateUsingToken(next AuthenticationMiddlewareFunction) http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		header := r.Header.Get("Authorization")
		if header == "" || !strings.Contains(header, "Bearer") {
			helpers.JsonResponse(w, 401, "unauthorized", nil)
			return
		}

		token := strings.Split(header, " ")[1]
		claims, err := helpers.ValidateToken(token)
		if err != nil {
			helpers.JsonResponse(w, 401, "unauthorized", nil)
			return
		}

		subject, err := claims.GetSubject()
		if err != nil {
			helpers.JsonResponse(w, 401, "unauthorized", nil)
			return
		}

		userId, err := strconv.Atoi(subject)
		if err != nil {
			helpers.JsonResponse(w, 401, "unauthorized", nil)
			return
		}

		next(w, r, userId)
	})
}
