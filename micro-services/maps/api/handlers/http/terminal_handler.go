package http

import (
	"encoding/json"
	"net/http"

	"gholi-fly-maps/internal/terminals/domain"
	"gholi-fly-maps/internal/terminals/port"

	"github.com/go-chi/chi/v5"
	"github.com/google/uuid"
)

type TerminalHandler struct {
	service port.TerminalService
}

func NewTerminalHandler(service port.TerminalService) *TerminalHandler {
	return &TerminalHandler{service: service}
}

func (h *TerminalHandler) CreateTerminal(w http.ResponseWriter, r *http.Request) {
	var terminal domain.Terminal
	if err := json.NewDecoder(r.Body).Decode(&terminal); err != nil {
		http.Error(w, "Invalid request payload", http.StatusBadRequest)
		return
	}

	terminal.ID = uuid.New()
	if err := h.service.CreateTerminal(r.Context(), &terminal); err != nil {
		http.Error(w, "Failed to create terminal", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(terminal)
}

func (h *TerminalHandler) GetTerminalByID(w http.ResponseWriter, r *http.Request) {
	idParam := chi.URLParam(r, "id")
	id, err := uuid.Parse(idParam)
	if err != nil {
		http.Error(w, "Invalid terminal ID", http.StatusBadRequest)
		return
	}

	terminal, err := h.service.GetTerminalByID(r.Context(), id)
	if err != nil {
		http.Error(w, "Terminal not found", http.StatusNotFound)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(terminal)
}

// GetAllTerminals handles the GET request to retrieve all terminals.
func (h *TerminalHandler) GetAllTerminals(w http.ResponseWriter, r *http.Request) {
	terminals, err := h.service.GetAllTerminals(r.Context())
	if err != nil {
		http.Error(w, "Failed to retrieve terminals", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(terminals)
}

func (h *TerminalHandler) UpdateTerminal(w http.ResponseWriter, r *http.Request) {
	// Get the terminal ID from the URL
	idParam := chi.URLParam(r, "id")
	id, err := uuid.Parse(idParam)
	if err != nil {
		http.Error(w, "Invalid terminal ID", http.StatusBadRequest)
		return
	}

	// Decode the request body into a partial terminal update
	var updateData map[string]interface{}
	if err := json.NewDecoder(r.Body).Decode(&updateData); err != nil {
		http.Error(w, "Invalid request payload", http.StatusBadRequest)
		return
	}

	// Validate the fields
	if len(updateData) == 0 {
		http.Error(w, "No update data provided", http.StatusBadRequest)
		return
	}

	// Fetch the existing terminal
	terminal, err := h.service.GetTerminalByID(r.Context(), id)
	if err != nil {
		http.Error(w, "Terminal not found", http.StatusNotFound)
		return
	}

	// Apply updates to the terminal
	if name, ok := updateData["name"].(string); ok && name != "" {
		terminal.Name = name
	}
	if location, ok := updateData["location"].(string); ok && location != "" {
		terminal.Location = location
	}
	if terminalType, ok := updateData["type"].(string); ok {
		terminal.Type = terminalType
	}

	// Update the terminal in the database
	if err := h.service.UpdateTerminal(r.Context(), terminal); err != nil {
		http.Error(w, "Failed to update terminal", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(terminal)
}

func (h *TerminalHandler) DeleteTerminal(w http.ResponseWriter, r *http.Request) {
	// Get the terminal ID from the URL
	idParam := chi.URLParam(r, "id")
	id, err := uuid.Parse(idParam)
	if err != nil {
		http.Error(w, "Invalid terminal ID", http.StatusBadRequest)
		return
	}

	// Fetch the terminal to ensure it exists
	// terminal, err := h.service.GetTerminalByID(r.Context(), id)
	// if err != nil {
	// 	http.Error(w, "Terminal not found", http.StatusNotFound)
	// 	return
	// }

	// Check if the user has permission to delete (placeholder)
	// TODO: Implement user permission check
	// if !userHasPermission(r.Context(), terminal) {
	//     http.Error(w, "User does not have permission to delete this terminal", http.StatusForbidden)
	//     return
	// }

	// Delete the terminal
	if err := h.service.DeleteTerminal(r.Context(), id); err != nil {
		http.Error(w, "Failed to delete terminal", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusNoContent) 
}

func (h *TerminalHandler) SearchTerminals(w http.ResponseWriter, r *http.Request) {
	// Parse query parameters
	queryParams := r.URL.Query()
	idParam := queryParams.Get("id")
	nameParam := queryParams.Get("name")
	cityParam := queryParams.Get("city")
	typeParam := queryParams.Get("type")
	// pagination parameters (not implemented)

	// Build the filter object
	filter := port.TerminalFilter{
		ID:   idParam,
		Name: nameParam,
		City: cityParam,
		Type: typeParam,
		
	}

	// Perform the search using the service
	terminals, err := h.service.SearchTerminals(r.Context(), filter)
	if err != nil {
		http.Error(w, "Failed to search terminals", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(terminals)
}


func RegisterTerminalRoutes(r chi.Router, service port.TerminalService) {
	handler := NewTerminalHandler(service)
	r.Post("/terminals", handler.CreateTerminal)
	r.Get("/terminals/{id}", handler.GetTerminalByID)
	r.Get("/terminals", handler.GetAllTerminals)
	r.Put("/terminals/{id}", handler.UpdateTerminal)
	r.Delete("/terminals/{id}", handler.DeleteTerminal)
	r.Get("/terminals/search", handler.SearchTerminals)
}
