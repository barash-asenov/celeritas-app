package handlers

import (
	"myapp/data"
	"net/http"
	"time"

	"github.com/CloudyKit/jet/v6"
	"github.com/barash-asenov/celeritas"
)

type Handlers struct {
	App    *celeritas.Celeritas
	Models data.Models
}

func (h *Handlers) Home(rw http.ResponseWriter, r *http.Request) {
	defer h.App.LoadTime(time.Now())
	err := h.render(rw, r, "home", nil, nil)
	if err != nil {
		h.App.ErrorLog.Println("error rendering:", err)
	}
}

func (h *Handlers) GoPage(rw http.ResponseWriter, r *http.Request) {
	err := h.App.Render.GoPage(rw, r, "home", nil)
	if err != nil {
		h.App.ErrorLog.Println("error rendering:", err)
	}
}

func (h *Handlers) JetPage(rw http.ResponseWriter, r *http.Request) {
	err := h.App.Render.JetPage(rw, r, "jet-template", nil, nil)
	if err != nil {
		h.App.ErrorLog.Println("error rendering:", err)
	}
}

func (h *Handlers) SessionTest(rw http.ResponseWriter, r *http.Request) {
	myData := "bar"

	h.App.Session.Put(r.Context(), "foo", myData)

	myValue := h.App.Session.GetString(r.Context(), "foo")

	vars := make(jet.VarMap)
	vars.Set("foo", myValue)

	err := h.App.Render.JetPage(rw, r, "sessions", vars, nil)
	if err != nil {
		h.App.ErrorLog.Println("error rendering:", err)
	}
}

func (h *Handlers) JSON(rw http.ResponseWriter, r *http.Request) {
	type Payload struct {
		ID int64 `json:"id"`
		Name string `json:"name"`
		Hobbies []string `json:"hobbies"`
	}

	var payload Payload

	payload.ID = 10
	payload.Name = "Jack Jones"
	payload.Hobbies = []string{"karate", "computer", "programming", "tennis", "gym"}

	err := h.App.WriteJSON(rw, http.StatusOK, payload)
	if err != nil {
		h.App.ErrorLog.Println(err)
	}
}

func (h *Handlers) XML(rw http.ResponseWriter, r *http.Request)  {
	type Payload struct {
		ID int64 `xml:"id"`
		Name string `xml:"name"`
		Hobbies []string `xml:"hobbies>hobby"`
	}

	var payload Payload

	payload.ID = 10
	payload.Name = "Jack Jones"
	payload.Hobbies = []string{"karate", "computer", "programming", "tennis", "gym"}

	err := h.App.WriteXML(rw, http.StatusOK, payload)
	if err != nil {
		h.App.ErrorLog.Println(err)
	}
}

func (h *Handlers) DownloadFile(rw http.ResponseWriter, r *http.Request) {
	h.App.DownloadFile(rw, r, "./public/images", "celeritas.jpg")
}
