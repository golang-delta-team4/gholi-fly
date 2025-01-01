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

	pattt, err := h.service.CreatePath(c.Context(), &path)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Failed to create path"})
	}

	return c.Status(fiber.StatusCreated).JSON(pattt)
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
	app.Get("/api/v1/paths/all", handler.GetAllPaths)          // GET all paths
	app.Post("/api/v1/paths/new", handler.CreatePath)          // POST create new path
	app.Get("/api/v1/paths/filter", handler.FilterPaths)       // GET filter paths dynamically
	app.Put("/api/v1/paths/update/:id", handler.UpdatePath)    // PUT update path by ID
	app.Delete("/api/v1/paths/delete/:id", handler.DeletePath) // DELETE path by ID
}
