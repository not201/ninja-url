package http

import (
	"embed"
	"encoding/json"
	"io/fs"
	"log"
	"net/http"

	"github.com/not201/ninja-url/internal/core/domain/dto"
	"github.com/not201/ninja-url/internal/core/ports/service"
	"github.com/not201/ninja-url/internal/utils"
)

type handler struct {
	service  service.UrlService
	staticFS embed.FS
}

func NewHandler(service service.UrlService, staticFS embed.FS) *handler {
	return &handler{service: service, staticFS: staticFS}
}

func (h *handler) Health(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]any{
		"message": "ok",
	})
}

func (h *handler) NotFound(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusNotFound)
	json.NewEncoder(w).Encode(map[string]any{
		"message": "route does not exist",
	})
}

func (h *handler) MethodNotAllowed(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusMethodNotAllowed)
	json.NewEncoder(w).Encode(map[string]any{
		"message": "method is not valid",
	})
}

func (h *handler) Shorten(w http.ResponseWriter, r *http.Request) {
	originalUrl := r.FormValue("url")

	if !utils.IsValidUrl(originalUrl) {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(map[string]any{
			"message": "invalid url",
		})
		return
	}

	urlDto := &dto.UrlDto{
		OriginalUrl: originalUrl,
	}

	url, err := h.service.Shorten(r.Context(), urlDto)
	if err != nil {
		log.Print(err)
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(map[string]any{
			"message": "something went wrong!!!",
		})
		return
	}

	shortUrl := "https://" + r.Host + "/" + url.ShortCode

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]any{
		"message": "ok",
		"data": map[string]any{
			"short_url":  shortUrl,
			"expires_at": url.ExpiresAt,
		},
	})

}

func (h *handler) Resolver(w http.ResponseWriter, r *http.Request) {
	shortCode := r.PathValue("shortCode")

	originalURL, err := h.service.Resolver(r.Context(), shortCode)
	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		json.NewEncoder(w).Encode(map[string]any{
			"message": "url not found",
		})
		return
	}

	http.Redirect(w, r, originalURL, http.StatusMovedPermanently)
}

func (h *handler) Static(w http.ResponseWriter, r *http.Request) {
	distFS, err := fs.Sub(h.staticFS, "web/dist")
	if err != nil {
		log.Print(err)
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(map[string]any{
			"message": "something went wrong!!!",
		})
		return
	}

	fileServer := http.FileServer(http.FS(distFS))
	fileServer.ServeHTTP(w, r)
}
