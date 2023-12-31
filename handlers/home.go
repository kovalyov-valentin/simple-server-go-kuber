package handlers

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/kovalyov-valentin/simple-server-go-kuber/version"
)

func home(buildTime, commit, release string) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		info := struct {
			BuildTime string `json:"build_time"`
			Commit    string `json:"commit"`
			Release   string `json:"release"`
		}{
			version.BuildTime, version.Commit, version.Release,
		}

		body, err := json.Marshal(info)
		if err != nil {
			log.Printf("Could not encode info data: %v", err)
			http.Error(w, http.StatusText(http.StatusServiceUnavailable), http.StatusServiceUnavailable)
			return
		}

		w.Header().Set("Content-Type", "application/json")
		w.Write(body)
	}
}
