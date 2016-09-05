package handlers

import "net/http"

type Home struct {
}

func (home *Home) ServeHTTP(w http.ResponseWriter, r *http.Request) {

	if err := serveTemplate(w, "home", nil); err != nil {

		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

}
