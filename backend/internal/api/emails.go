package api

import (
	"encoding/json"
	"net/http"

	generated "github.com/egeuysall/learn-testing/internal/repository/generated"
	"github.com/egeuysall/learn-testing/internal/services"
	"github.com/jackc/pgx/v5/pgxpool"
)

func SignupHandlerWithEmail(pool *pgxpool.Pool, emailSender services.EmailSender) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req SignupRequest
		if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
			http.Error(w, "invalid request", http.StatusBadRequest)
			return
		}

		queries := generated.New(pool)
		user, err := queries.CreateUser(r.Context(), generated.CreateUserParams{
			Email: req.Email,
			Name:  req.Name,
		})
		if err != nil {
			http.Error(w, "failed to create user", http.StatusInternalServerError)
			return
		}

		// Send welcome email (don't fail signup if email fails)
		emailSender.SendWelcomeEmail(user.Email, user.Name)

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusCreated)
		json.NewEncoder(w).Encode(SignupResponse{
			ID:    user.ID,
			Email: user.Email,
			Name:  user.Name,
		})
	}
}
