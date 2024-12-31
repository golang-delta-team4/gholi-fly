package http

import (
	"errors"
	"fmt"
	"log"
	"time"
	"vehicle/api/presenter"
	vehicleService "vehicle/internal/vehicle"
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
	var vehicleMatchRequest presenter.MatchMakerRequest
	if err := c.BodyParser(&vehicleMatchRequest); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": fmt.Sprintf("Invalid request payload %v", err)})
	}
	reserveStartDate, err := time.Parse(time.DateOnly, vehicleMatchRequest.ReserveStartDate)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": err})
	}
	reserveEndDate, err := time.Parse(time.DateOnly, vehicleMatchRequest.ReserveEndDate)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": err})
	}
	reservationID, vehicle, err := h.service.MatchVehicle(c.Context(), &domain.MatchMakerRequest{
		TripID:             vehicleMatchRequest.TripID,
		ReserveStartDate:   reserveStartDate,
		ReserveEndDate:     reserveEndDate,
		TripDistance:       vehicleMatchRequest.TripDistance,
		NumberOfPassengers: vehicleMatchRequest.NumberOfPassengers,
		TripType:           domain.TripType(vehicleMatchRequest.TripType),
		MaxPrice:           vehicleMatchRequest.MaxPrice,
		YearOfManufacture:  vehicleMatchRequest.YearOfManufacture,
	})
	if err != nil {
		if errors.Is(err, vehicleService.ErrVehicleNotFound) {
			return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"error": err.Error()})
		}
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"error": err.Error()})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{"reservation_id": reservationID, "vehicle_detail": vehicle})
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
	app.Post("/api/v1/vehicles", handler.CreateVehicle)
	app.Get("/api/v1/vehicles/match", handler.MatchVehicle)
}
