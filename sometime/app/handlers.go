package app

import (
	"encoding/json"
	"net/http"
	"strings"
	"time"
)

func getCurrentTime(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Content-Type", "application/json")

	query := r.URL.Query().Get("tz")
	if query == "" {
		query = "UTC"
	}
	zones := strings.Split(query, ",")

	timezonemap := map[string]string{}
	for _, zone := range zones {
		loc, err := time.LoadLocation(zone)
		if err != nil {
			w.WriteHeader(http.StatusNotFound)
			json.NewEncoder(w).Encode("invalid timezone")
			return
		}

		timestamp := time.Now().In(loc).String()
		if len(zones) == 1 {
			timezonemap["current_time"] = timestamp
		} else {
			timezonemap[zone] = timestamp
		}
	}

	json.NewEncoder(w).Encode(timezonemap)
}
