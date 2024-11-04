package server

import (
	"encoding/json"
	"net/http"
	"path/filepath"

	"github.com/tarsh333/go_db/model"
	"github.com/tarsh333/go_db/utils"
)

// Middleware to check and add a directory for each user ID
func checkAndAddUserDirectory(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		var params model.RequestParams
		err := json.NewDecoder(r.Body).Decode(&params)
		if err != nil {
			http.Error(w, "Invalid request body", http.StatusBadRequest)
			return
		}

		if params.Id == "" {
			http.Error(w, "User ID cannot be empty", http.StatusBadRequest)
			return
		}

		// Create the user directory if it doesn't already exist
		utils.CreateFolder(filepath.Join("db", params.Id))

		next.ServeHTTP(w, r)
	})
}

// Middleware to set response type to JSON
func addResponseType(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Add("Content-Type", "application/json")
		next.ServeHTTP(w, r)
	})
}
