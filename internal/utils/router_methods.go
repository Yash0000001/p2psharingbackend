package utils

import (
	"net/http"
)

func POSTOnly(handler http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPost {
			SendError(w, http.StatusMethodNotAllowed, "Invalid Request method", nil)
			return
		}
		handler(w, r)
	}
}

func GETOnly(handler http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodGet {
			SendError(w, http.StatusMethodNotAllowed, "Invalid Request method", nil)
			return
		}
		handler(w, r)
	}
}

func PUTOnly(handler http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPut {
			SendError(w, http.StatusMethodNotAllowed, "Invalid Request method", nil)
			return
		}
		handler(w, r)
	}
}

func PATCHOnly(handler http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPatch {
			SendError(w, http.StatusMethodNotAllowed, "Invalid Request method", nil)
			return
		}
		handler(w, r)
	}
}

func DELETEOnly(handler http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodDelete {
			SendError(w, http.StatusMethodNotAllowed, "Invalid Request method", nil)
			return
		}
		handler(w, r)
	}
}
