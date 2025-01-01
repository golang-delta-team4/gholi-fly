package http

import (
	"gholi-fly-maps/internal/terminals/domain"
	"gholi-fly-maps/internal/terminals/port"

	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
)

type TerminalHandler struct {
	service port.TerminalService
}

func NewTerminalHandler(service port.TerminalService) *TerminalHandler {
	return &TerminalHandler{service: service}
}

func (h *TerminalHandler) CreateTerminal(c *fiber.Ctx) error {
	var terminal domain.Terminal
	if err := c.BodyParser(&terminal); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid request payload"})
	}

	if terminal.Location == "" {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Location is required"})
	}

	terminal.ID = uuid.New()
	if err := h.service.CreateTerminal(c.Context(), &terminal); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Failed to create terminal"})
	}

	return c.Status(fiber.StatusCreated).JSON(terminal)
}

func (h *TerminalHandler) GetTerminalByID(c *fiber.Ctx) error {
	id, err := uuid.Parse(c.Params("id"))
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid terminal ID"})
	}

	terminal, err := h.service.GetTerminalByID(c.Context(), id)
	if err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"error": "Terminal not found"})
	}

	return c.Status(fiber.StatusOK).JSON(terminal)
}

func (h *TerminalHandler) GetAllTerminals(c *fiber.Ctx) error {
	terminals, err := h.service.GetAllTerminals(c.Context())
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Failed to retrieve terminals"})
	}

	return c.Status(fiber.StatusOK).JSON(terminals)
}

func (h *TerminalHandler) UpdateTerminal(c *fiber.Ctx) error {
	id, err := uuid.Parse(c.Params("id"))
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid terminal ID"})
	}

	var terminal domain.Terminal
	if err := c.BodyParser(&terminal); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid request payload"})
	}
	terminal.ID = id

	if err := h.service.UpdateTerminal(c.Context(), &terminal); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Failed to update terminal"})
	}

	return c.Status(fiber.StatusOK).JSON(terminal)
}

func (h *TerminalHandler) DeleteTerminal(c *fiber.Ctx) error {
	id, err := uuid.Parse(c.Params("id"))
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid terminal ID"})
	}

	if err := h.service.DeleteTerminal(c.Context(), id); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Failed to delete terminal"})
	}

	return c.Status(fiber.StatusNoContent).Send(nil)
}

// SearchTerminals handles the dynamic search for terminals.
func (h *TerminalHandler) SearchTerminals(c *fiber.Ctx) error {
	// Parse query parameters into filter
	filter := port.TerminalFilter{
		ID:   c.Query("id"),
		Name: c.Query("name"),
		City: c.Query("city"),
		Type: c.Query("type"),
	}

	// Validate filters
	if filter.ID == "" && filter.Name == "" && filter.City == "" && filter.Type == "" {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "At least one filter parameter must be provided",
		})
	}

	// Delegate to the service layer
	terminals, err := h.service.SearchTerminals(c.Context(), filter)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Failed to search terminals: " + err.Error(),
		})
	}

	return c.Status(fiber.StatusOK).JSON(terminals)
}

func RegisterTerminalRoutes(app *fiber.App, service port.TerminalService) {
	handler := NewTerminalHandler(service)

	// Define the endpoints for terminals
	app.Get("/api/v1/terminals/all", handler.GetAllTerminals)
	app.Get("/api/v1/terminals/:id", handler.GetTerminalByID)
	app.Post("/api/v1/terminals/new", handler.CreateTerminal)
	app.Put("/api/v1/terminals/update/:id", handler.UpdateTerminal)
	app.Delete("/api/v1/terminals/delete/:id", handler.DeleteTerminal)
	app.Get("/api/v1/terminals/search", handler.SearchTerminals)
}
