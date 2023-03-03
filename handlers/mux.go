package handlers

import (
	"net/http"
	"odo/dep"
)

func Mux(d *dep.Dependencies) *http.ServeMux {

	mux := http.NewServeMux()

	// catalog static files
	static := http.StripPrefix("/static/", http.FileServer(http.Dir("./static/")))
	mux.Handle("/static/", static)

	// catalog page
	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		catalog(w, r, d)
	})

	// dynamically add handlers
	for _, handler := range d.Handlers {
		server := http.StripPrefix(handler.Name, http.FileServer(http.Dir(handler.Dir)))
		mux.Handle(handler.Name, server)
	}

	// logs handler
	mux.Handle("/toys.log", http.FileServer(http.Dir("logs/")))

	return mux
}
