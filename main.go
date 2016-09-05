// FirstApp project main.go
package main

import (
	"log"
	"net/http"
	"smile/handlers"
)

const port = ":9012"

func main() {

	homeHandle := new(handlers.Home)
	partitionHandle := new(handlers.Partition)
	userHandle := new(handlers.User)
	localeHandle := new(handlers.Locale)
	timezoneHandle := new(handlers.Timezone)

	http.Handle("/static/", http.StripPrefix("/static/",
		http.FileServer(http.Dir("./public/assets"))))

	http.Handle("/", homeHandle)
	http.Handle("/partition", partitionHandle)
	http.Handle("/user", userHandle)
	http.Handle("/locale", localeHandle)
	http.Handle("/timezone", timezoneHandle)

	log.Println("Aplicativo Smile iniciado na porta", port)
	http.ListenAndServe(port, nil)

}
