package handlers

import (
	"net/http"
	"odo/dep"
)

func catalog(w http.ResponseWriter, r *http.Request, d *dep.Dependencies) {
	// handle method
	if r.Method != http.MethodGet {
		http.NotFound(w, r)
		return
	}

	// return template
	if err := d.Tmp.ExecuteTemplate(w, "catalog.html", d.Handlers); err != nil {
		http.Error(w, "Internal Server Error", 500)
		d.Log.Error(err)
		return
	}
}
