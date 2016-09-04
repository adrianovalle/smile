package handlers

import (
	"net/http"
	"smile/models"
)

type Partition models.Partition

func (partition *Partition) ServeHTTP(w http.ResponseWriter, r *http.Request) {

	if err := serveTemplate(w, "partition", nil); err != nil {

		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

}
