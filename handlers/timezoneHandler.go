package handlers

import (
	"net/http"
	"smile/models"
)

type Timezone models.Timezone

func (timezone *Timezone) ServeHTTP(w http.ResponseWriter, r *http.Request) {

	if err := serveTemplate(w, "timezone", nil); err != nil {

		http.Error(w, err.Error(), http.StatusInternalServerError)

	}

}
