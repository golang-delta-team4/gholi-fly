package http

import (
	"gholi-fly-maps/internal/paths/domain"
	"gholi-fly-maps/internal/paths/port"

	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
)

type PathHandler struct {
	service port.PathService
}

func NewPathHandler(service port.PathService) *PathHandler {
	return &PathHandler{service: service}
}

func (h *PathHandler) GetAllPaths(c *fiber.Ctx) error {
	paths, err := h.service.GetAllPaths(c.Context())
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Failed to fetch paths"})
	}

	return c.Status(fiber.StatusOK).JSON(paths)
}

func (h *PathHandler) CreatePath(c *fiber.Ctx) error {
	var path domain.Path
	if err := c.BodyParser(&path); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid request payload"})
	}

	if path.SourceTerminalID == uuid.Nil || path.DestinationTerminalID == uuid.Nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid terminal IDs"})
	}

	if err := h.service.CreatePath(c.Context(), &path); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Failed to create path"})
	}

	return c.Status(fiber.StatusCreated).JSON(path)
}

func (h *PathHandler) UpdatePath(c *fiber.Ctx) error {
	id, err := uuid.Parse(c.Params("id"))
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid path ID"})
	}

	var path domain.Path
	if err := c.BodyParser(&path); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid request payload"})
	}
	path.ID = id

	if err := h.service.UpdatePath(c.Context(), &path); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{"message": "Path updated successfully"})
}

func (h *PathHandler) DeletePath(c *fiber.Ctx) error {
	id, err := uuid.Parse(c.Params("id"))
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid path ID"})
	}

	if err := h.service.DeletePath(c.Context(), id); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}

	return c.Status(fiber.StatusNoContent).Send(nil)
}

func (h *PathHandler) FilterPaths(c *fiber.Ctx) error {
	filters := make(map[string]interface{})
	if id := c.Query("id"); id != "" {
		filters["id"] = id
	}
	if sourceTerminalID := c.Query("source_terminal_id"); sourceTerminalID != "" {
		filters["source_terminal_id"] = sourceTerminalID
	}
	if destinationTerminalID := c.Query("destination_terminal_id"); destinationTerminalID != "" {
		filters["destination_terminal_id"] = destinationTerminalID
	}
	if distanceKM := c.Query("distance_km"); distanceKM != "" {
		filters["distance_km"] = distanceKM
	}
	if routeCode := c.Query("route_code"); routeCode != "" {
		filters["route_code"] = routeCode
	}
	if vehicleType := c.Query("vehicle_type"); vehicleType != "" {
		filters["vehicle_type"] = vehicleType
	}

	paths, err := h.service.FilterPaths(c.Context(), filters)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Failed to filter paths"})
	}

	return c.Status(fiber.StatusOK).JSON(paths)
}
func RegisterPathRoutes(app *fiber.App, service port.PathService) {
	handler := NewPathHandler(service)

	// Define the endpoints for paths
	app.Get("/api/v1/paths/all", handler.GetAllPaths)         // GET all paths
	app.Post("/api/v1/paths/new", handler.CreatePath)         // POST create new path
	app.Get("/api/v1/paths/filter", handler.FilterPaths)      // GET filter paths dynamically
	app.Put("/api/v1/paths/update/:id", handler.UpdatePath)   // PUT update path by ID
	app.Delete("/api/v1/paths/delete/:id", handler.DeletePath) // DELETE path by ID
}



// import (
// 	"encoding/json"
// 	"net/http"

// 	"gholi-fly-maps/internal/paths/domain"
// 	"gholi-fly-maps/internal/paths/port"

// 	"github.com/go-chi/chi/v5"
// 	"github.com/google/uuid"
// )

// type PathHandler struct {
// 	service port.PathService
// }

// // NewPathHandler creates a new PathHandler instance.
// func NewPathHandler(service port.PathService) *PathHandler {
// 	return &PathHandler{service: service}
// }

// // GetAllPaths retrieves all paths from the service.
// func (h *PathHandler) GetAllPaths(w http.ResponseWriter, r *http.Request) {
// 	paths, err := h.service.GetAllPaths(r.Context())
// 	if err != nil {
// 		http.Error(w, "Failed to fetch paths", http.StatusInternalServerError)
// 		return
// 	}

// 	w.Header().Set("Content-Type", "application/json")
// 	json.NewEncoder(w).Encode(paths)
// }

// // CreatePath handles the creation of a new path.
// func (h *PathHandler) CreatePath(w http.ResponseWriter, r *http.Request) {
// 	var path domain.Path
// 	if err := json.NewDecoder(r.Body).Decode(&path); err != nil {
// 		http.Error(w, "Invalid request payload", http.StatusBadRequest)
// 		return
// 	}

// 	// Ensure SourceTerminalID and DestinationTerminalID are valid UUIDs
// 	if path.SourceTerminalID == uuid.Nil || path.DestinationTerminalID == uuid.Nil {
// 		http.Error(w, "Invalid terminal IDs", http.StatusBadRequest)
// 		return
// 	}

// 	// Delegate to the service layer
// 	if err := h.service.CreatePath(r.Context(), &path); err != nil {
// 		http.Error(w, "Failed to create path", http.StatusInternalServerError)
// 		return
// 	}

// 	w.WriteHeader(http.StatusCreated)
// 	json.NewEncoder(w).Encode(path)
// }

// // UpdatePath handles the request to update an existing path.
// func (h *PathHandler) UpdatePath(w http.ResponseWriter, r *http.Request) {
// 	// Extract path ID from the URL
// 	idParam := chi.URLParam(r, "id")
// 	id, err := uuid.Parse(idParam)
// 	if err != nil {
// 		http.Error(w, "Invalid path ID", http.StatusBadRequest)
// 		return
// 	}

// 	// Decode the request payload
// 	var path domain.Path
// 	if err := json.NewDecoder(r.Body).Decode(&path); err != nil {
// 		http.Error(w, "Invalid request payload", http.StatusBadRequest)
// 		return
// 	}
// 	path.ID = id // Set the ID from the URL

// 	// Call the service to update the path
// 	err = h.service.UpdatePath(r.Context(), &path)
// 	if err != nil {
// 		http.Error(w, err.Error(), http.StatusBadRequest)
// 		return
// 	}

// 	// Return a success message
// 	w.Header().Set("Content-Type", "application/json")
// 	w.WriteHeader(http.StatusOK)
// 	json.NewEncoder(w).Encode(map[string]string{
// 		"message": "Path updated successfully",
// 	})
// }

// // DeletePath handles the request to delete a path.
// func (h *PathHandler) DeletePath(w http.ResponseWriter, r *http.Request) {
// 	// Extract path ID from the URL
// 	idParam := chi.URLParam(r, "id")
// 	id, err := uuid.Parse(idParam)
// 	if err != nil {
// 		http.Error(w, "Invalid path ID", http.StatusBadRequest)
// 		return
// 	}

// 	// Call the service to delete the path
// 	err = h.service.DeletePath(r.Context(), id)
// 	if err != nil {
// 		http.Error(w, err.Error(), http.StatusBadRequest)
// 		return
// 	}

// 	// Return a success message
// 	w.Header().Set("Content-Type", "application/json")
// 	w.WriteHeader(http.StatusOK)
// 	json.NewEncoder(w).Encode(map[string]string{
// 		"message": "Path deleted successfully",
// 	})
// }
// func (h *PathHandler) FilterPaths(w http.ResponseWriter, r *http.Request) {
// 	// Collect filters from query parameters
// 	filters := make(map[string]interface{})

// 	if id := r.URL.Query().Get("id"); id != "" {
// 		filters["id"] = id
// 	}
// 	if sourceTerminalID := r.URL.Query().Get("source_terminal_id"); sourceTerminalID != "" {
// 		filters["source_terminal_id"] = sourceTerminalID
// 	}
// 	if destinationTerminalID := r.URL.Query().Get("destination_terminal_id"); destinationTerminalID != "" {
// 		filters["destination_terminal_id"] = destinationTerminalID
// 	}
// 	if distanceKM := r.URL.Query().Get("distance_km"); distanceKM != "" {
// 		filters["distance_km"] = distanceKM
// 	}
// 	if routeCode := r.URL.Query().Get("route_code"); routeCode != "" {
// 		filters["route_code"] = routeCode
// 	}
// 	if vehicleType := r.URL.Query().Get("vehicle_type"); vehicleType != "" {
// 		filters["vehicle_type"] = vehicleType
// 	}

// 	// Call service layer
// 	paths, err := h.service.FilterPaths(r.Context(), filters)
// 	if err != nil {
// 		http.Error(w, "Failed to filter paths", http.StatusInternalServerError)
// 		return
// 	}

// 	// Send response
// 	w.Header().Set("Content-Type", "application/json")
// 	json.NewEncoder(w).Encode(paths)
// }

// // RegisterPathRoutes registers path-related routes.
// func RegisterPathRoutes(r chi.Router, service port.PathService) {
// 	handler := NewPathHandler(service)
// 	r.Get("/paths/all", handler.GetAllPaths)
// 	r.Post("/paths/new", handler.CreatePath)
// 	r.Get("/paths/filter", handler.FilterPaths)
// 	r.Put("/paths/update/{id}", handler.UpdatePath)
// 	r.Delete("/paths/delete/{id}", handler.DeletePath)

// }
