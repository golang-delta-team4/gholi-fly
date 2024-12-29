package http

import (
	"log"
	"vehicle/internal/vehicle/domain"
	"vehicle/internal/vehicle/port"

	"github.com/gofiber/fiber/v2"
)

type VehicleHandler struct {
	service port.VehicleService
}

func NewVehicleHandler(service port.VehicleService) *VehicleHandler {
	return &VehicleHandler{service: service}
}

func (h *VehicleHandler) MatchVehicle(c *fiber.Ctx) error {
	// Parse query parameters using Fiber's QueryParser
	var tripRequest domain.TripRequest
	if err := c.QueryParser(&tripRequest); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid query parameters"})
	}

	// Validate the parsed TripRequest
	if tripRequest.TripType == "" || tripRequest.MinPassengers <= 0 {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Missing or invalid required parameters"})
	}

	vehicle, err := h.service.MatchVehicle(c.Context(), &tripRequest)
	if err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"error": err.Error()})
	}

	return c.Status(fiber.StatusOK).JSON(vehicle)
}

func (h *VehicleHandler) CreateVehicle(c *fiber.Ctx) error {
	log.Println("Incoming request to create vehicle")

	// Parse body using Fiber's BodyParser
	var vehicle domain.Vehicle
	if err := c.BodyParser(&vehicle); err != nil {
		log.Printf("Error parsing body: %v", err)
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid request payload",
		})
	}

	// Validate required fields
	if vehicle.UniqueCode == "" || vehicle.Capacity <= 0 || vehicle.YearOfManufacture <= 0 {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Missing or invalid required fields",
		})
	}

	log.Printf("Vehicle parsed: %+v", vehicle)

	if err := h.service.CreateVehicle(c.Context(), &vehicle); err != nil {
		log.Printf("Error creating vehicle: %v", err)
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	log.Printf("Vehicle created: %+v", vehicle)
	return c.Status(fiber.StatusCreated).JSON(vehicle)
}

func RegisterVehicleRoutes(app *fiber.App, service port.VehicleService) {
	handler := NewVehicleHandler(service)

	// Register routes with Fiber
	app.Post("/api/v1/vehicles", handler.CreateVehicle)
	app.Post("/api/v1/vehicles/match", handler.MatchVehicle)
}
