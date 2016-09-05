package handlers

import (
	"net/http"
	"smile/models"
)

type User models.User

func (user *User) ServeHTTP(w http.ResponseWriter, r *http.Request) {

	if err := serveTemplate(w, "user", nil); err != nil {

		http.Error(w, err.Error(), http.StatusInternalServerError)

	}

}
