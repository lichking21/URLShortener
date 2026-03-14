package handler

import (
	"encoding/json"
	"net/http"
)

type ShortenRequest struct {
	URL string `json:"url"`
}
type ShortenResponse struct {
	ShortCode string `json:"short_code"`
}

// POST short url code
func (h *Handler) ShortUrl(w http.ResponseWriter, r *http.Request) {

	var req ShortenRequest

	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "(ERR) >> Failed to parse json", http.StatusBadRequest)
		return
	}

	defer r.Body.Close()

	shortCode, err := h.service.CreateShortCode(r.Context(), req.URL)
	if err != nil {
		http.Error(w, "(ERR) >> Failed to create short code", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated) // 201 Created
	json.NewEncoder(w).Encode(ShortenResponse{ShortCode: shortCode})
}

// GET long url
func (h *Handler) Redirect(w http.ResponseWriter, r *http.Request) {

	shortCode := r.PathValue("shortCode")
	if shortCode == "" {
		http.Error(w, "(ERR) >> Empty short code", http.StatusBadRequest)
		return
	}

	longUrl, err := h.service.GetLongUrl(r.Context(), shortCode)
	if err != nil {
		http.Error(w, "(ERR) >> URL not found", http.StatusNotFound)
		return
	}

	http.Redirect(w, r, longUrl, http.StatusMovedPermanently)
}
