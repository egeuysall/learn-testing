// handlers.go (in internal/api/ folder - create it)
package api

import (
	"encoding/json"
	"net/http"

	generated "github.com/egeuysall/learn-testing/internal/repository/generated"
	"github.com/jackc/pgx/v5/pgxpool"
)

type SignupRequest struct {
	Email string `json:"email"`
	Name  string `json:"name"`
}

type SignupResponse struct {
	ID    int32  `json:"id"`
	Email string `json:"email"`
	Name  string `json:"name"`
}

func SignupHandler(pool *pgxpool.Pool) http.HandlerFunc {
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

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusCreated)
		json.NewEncoder(w).Encode(SignupResponse{
			ID:    user.ID,
			Email: user.Email,
			Name:  user.Name,
		})
	}
}
