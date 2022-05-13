package handlers

import (
	"errors"
	"go-deploy-gke/api/httpext"
	"net/http"
)

func Index(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		httpext.WriteError(w, http.StatusMethodNotAllowed, errors.New(http.StatusText(http.StatusMethodNotAllowed)))
	} else if r.RequestURI != "/" {
		httpext.WriteError(w, http.StatusNotFound, errors.New(http.StatusText(http.StatusNotFound)))
	} else {
		httpext.WriteJSON(w, http.StatusOK, map[string]interface{}{"message": "OK"})
	}
}
