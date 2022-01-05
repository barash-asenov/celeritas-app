package handlers

import (
	"myapp/data"
	"net/http"

	"github.com/barash-asenov/celeritas"
)

type Handlers struct {
	App    *celeritas.Celeritas
	Models data.Models
}

func (h *Handlers) Home(rw http.ResponseWriter, r *http.Request) {
	err := h.render(rw, r, "home", nil, nil)
	if err != nil {
		h.App.ErrorLog.Println("error rendering:", err)
	}
}

