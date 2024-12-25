package http

import (
	"encoding/json"
	"net/http"

	"gholi-fly-maps/internal/paths/domain"
	"gholi-fly-maps/internal/paths/port"

	"github.com/go-chi/chi/v5"
	"github.com/google/uuid"
)

type PathHandler struct {
	service port.PathService
}

// NewPathHandler creates a new PathHandler instance.
func NewPathHandler(service port.PathService) *PathHandler {
	return &PathHandler{service: service}
}

// GetAllPaths retrieves all paths from the service.
func (h *PathHandler) GetAllPaths(w http.ResponseWriter, r *http.Request) {
	paths, err := h.service.GetAllPaths(r.Context())
	if err != nil {
		http.Error(w, "Failed to fetch paths", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(paths)
}

// CreatePath handles the creation of a new path.
func (h *PathHandler) CreatePath(w http.ResponseWriter, r *http.Request) {
	var path domain.Path
	if err := json.NewDecoder(r.Body).Decode(&path); err != nil {
		http.Error(w, "Invalid request payload", http.StatusBadRequest)
		return
	}

	// Ensure SourceTerminalID and DestinationTerminalID are valid UUIDs
	if path.SourceTerminalID == uuid.Nil || path.DestinationTerminalID == uuid.Nil {
		http.Error(w, "Invalid terminal IDs", http.StatusBadRequest)
		return
	}

	// Delegate to the service layer
	if err := h.service.CreatePath(r.Context(), &path); err != nil {
		http.Error(w, "Failed to create path", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(path)
}

// UpdatePath handles the request to update an existing path.
func (h *PathHandler) UpdatePath(w http.ResponseWriter, r *http.Request) {
	// Extract path ID from the URL
	idParam := chi.URLParam(r, "id")
	id, err := uuid.Parse(idParam)
	if err != nil {
		http.Error(w, "Invalid path ID", http.StatusBadRequest)
		return
	}

	// Decode the request payload
	var path domain.Path
	if err := json.NewDecoder(r.Body).Decode(&path); err != nil {
		http.Error(w, "Invalid request payload", http.StatusBadRequest)
		return
	}
	path.ID = id // Set the ID from the URL

	// Call the service to update the path
	err = h.service.UpdatePath(r.Context(), &path)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// Return a success message
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(map[string]string{
		"message": "Path updated successfully",
	})
}
// DeletePath handles the request to delete a path.
func (h *PathHandler) DeletePath(w http.ResponseWriter, r *http.Request) {
	// Extract path ID from the URL
	idParam := chi.URLParam(r, "id")
	id, err := uuid.Parse(idParam)
	if err != nil {
		http.Error(w, "Invalid path ID", http.StatusBadRequest)
		return
	}

	// Call the service to delete the path
	err = h.service.DeletePath(r.Context(), id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// Return a success message
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(map[string]string{
		"message": "Path deleted successfully",
	})
}

// RegisterPathRoutes registers path-related routes.
func RegisterPathRoutes(r chi.Router, service port.PathService) {
	handler := NewPathHandler(service)
	r.Get("/paths/all", handler.GetAllPaths)
	r.Post("/paths/new", handler.CreatePath)
	r.Put("/paths/update/{id}", handler.UpdatePath)
	r.Delete("/paths/delete/{id}", handler.DeletePath)

}
