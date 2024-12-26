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

func (h *VehicleHandler) CreateVehicle(c *fiber.Ctx) error {
	log.Println("Incoming request to create vehicle")

	var vehicle domain.Vehicle
	if err := c.BodyParser(&vehicle); err != nil {
		log.Printf("Error parsing body: %v", err)
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid request payload",
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
	app.Post("/vehicles", handler.CreateVehicle)
}
