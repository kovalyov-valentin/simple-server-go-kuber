package handlers

import (
	"net/http"
	"sync/atomic"
)

func readyz(isReady *atomic.Value) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if isReady == nil || !isReady.Load().(bool) {
			http.Error(w, http.StatusText(http.StatusServiceUnavailable), http.StatusServiceUnavailable)
			return 
		}
		w.WriteHeader(http.StatusOK)
	}
}