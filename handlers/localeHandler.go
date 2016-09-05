package handlers

import (
	"net/http"
	"smile/models"
)

type Locale models.Locale

func (locale *Locale) ServeHTTP(w http.ResponseWriter, r *http.Request) {

	if err := serveTemplate(w, "locale", nil); err != nil {

		http.Error(w, err.Error(), http.StatusInternalServerError)

	}

}
