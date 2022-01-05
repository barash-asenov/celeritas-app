package handlers

import (
	"github.com/justinas/nosurf"
	"net/http"
)

func (h *Handlers) ShowCachePage(rw http.ResponseWriter, r *http.Request) {
	err := h.render(rw, r, "cache", nil, nil)
	if err != nil {
		h.App.ErrorLog.Println("error rendering: ", err)
	}
}

func (h *Handlers) SaveInCache(rw http.ResponseWriter, r *http.Request) {
	var userInput struct {
		Name  string `json:"name"`
		Value string `json:"value"`
		CSRF  string `json:"csrf_token"`
	}

	err := h.App.ReadJSON(rw, r, &userInput)
	if err != nil {
		h.App.Error500(rw)
		return
	}

	if !nosurf.VerifyToken(nosurf.Token(r), userInput.CSRF) {
		h.App.Error500(rw)
		return
	}

	err = h.App.Cache.Set(userInput.Name, userInput.Value)
	if err != nil {
		h.App.Error500(rw)
		return
	}

	var resp struct {
		Error   bool   `json:"error"`
		Message string `json:"message"`
	}

	resp.Error = false
	resp.Message = "Saved in cache"

	_ = h.App.WriteJSON(rw, http.StatusCreated, resp)
}

func (h *Handlers) GetFromCache(rw http.ResponseWriter, r *http.Request) {
	var userInput struct {
		Name string `json:"name"`
		CSRF string `json:"csrf_token"`
	}
	var msg string
	var inCache = true

	err := h.App.ReadJSON(rw, r, &userInput)
	if err != nil {
		h.App.Error500(rw)
		return
	}

	if !nosurf.VerifyToken(nosurf.Token(r), userInput.CSRF) {
		h.App.Error500(rw)
		return
	}

	fromCache, err := h.App.Cache.Get(userInput.Name)
	if err != nil {
		msg = "Not found in cache"
		inCache = false
	}

	var resp struct {
		Error   bool   `json:"error"`
		Message string `json:"message"`
		Value   string `json:"value"`
	}

	if inCache {
		resp.Error = false
		resp.Message = "Success"
		resp.Value = fromCache.(string)
	} else {
		resp.Error = true
		resp.Message = msg
	}

	_ = h.App.WriteJSON(rw, http.StatusCreated, resp)
}

func (h *Handlers) DeleteFromCache(rw http.ResponseWriter, r *http.Request) {
	var userInput struct {
		Name string `json:"name"`
		CSRF string `json:"csrf_token"`
	}

	err := h.App.ReadJSON(rw, r, &userInput)
	if err != nil {
		h.App.Error500(rw)
		return
	}

	if !nosurf.VerifyToken(nosurf.Token(r), userInput.CSRF) {
		h.App.Error500(rw)
		return
	}

	err = h.App.Cache.Forget(userInput.Name)
	if err != nil {
		h.App.Error500(rw)
		return
	}

	var resp struct {
		Error   bool   `json:"error"`
		Message string `json:"message"`
	}

	resp.Error = false
	resp.Message = "Deleted from cache (if it existed)"

	_ = h.App.WriteJSON(rw, http.StatusCreated, resp)
}

func (h *Handlers) EmptyCache(rw http.ResponseWriter, r *http.Request) {
	var userInput struct {
		CSRF string `json:"csrf_token"`
	}

	err := h.App.ReadJSON(rw, r, &userInput)
	if err != nil {
		h.App.Error500(rw)
		return
	}

	if !nosurf.VerifyToken(nosurf.Token(r), userInput.CSRF) {
		h.App.Error500(rw)
		return
	}

	err = h.App.Cache.Empty()
	if err != nil {
		h.App.Error500(rw)
		return
	}

	var resp struct {
		Error   bool   `json:"error"`
		Message string `json:"message"`
	}

	resp.Error = false
	resp.Message = "Emptied cache"

	_ = h.App.WriteJSON(rw, http.StatusCreated, resp)
}
