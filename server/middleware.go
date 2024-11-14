package server

import (
	"encoding/json"
	"fmt"
	"net/http"
	"path/filepath"

	"github.com/tarsh333/go_db/model"
	"github.com/tarsh333/go_db/utils"
)

// Middleware to check and add a directory for each user ID
func checkAndAddDBDirectory(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		var params model.RequestParams
		err := json.NewDecoder(r.Body).Decode(&params)
		if params.Action == "get" {
			next.ServeHTTP(w, r)
			return
			// Add additional middleware logic here, e.g., authentication or logging
		}
		if err != nil {
			http.Error(w, "Invalid request body", http.StatusBadRequest)
			return
		}

		if params.Cluster == "" {
			http.Error(w, "Cluster cannot be empty", http.StatusBadRequest)
			return
		}
		if params.Database == "" {
			http.Error(w, "Database cannot be empty", http.StatusBadRequest)
			return
		}

		// Create the directory if it doesn't already exist
		if err := utils.CreateFolder(filepath.Join("db", params.Cluster, params.Database)); err != nil {
			fmt.Println("failed to add folder ", err)
			return
		}
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
