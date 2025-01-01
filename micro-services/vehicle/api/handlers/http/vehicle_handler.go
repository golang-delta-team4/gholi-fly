package http

import (
	"errors"
	"log"
	"time"
	"vehicle/api/presenter"
	"vehicle/config"
	vehicleService "vehicle/internal/vehicle"
	"vehicle/internal/vehicle/domain"
	"vehicle/internal/vehicle/port"

	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
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
		
		return fiber.NewError(fiber.StatusBadRequest, "invalid request payload", err.Error())
	}
	reserveStartDate, err := time.Parse(time.DateOnly, vehicleMatchRequest.ReserveStartDate)
	if err != nil {
		return fiber.NewError(fiber.StatusBadRequest, err.Error())
	}
	reserveEndDate, err := time.Parse(time.DateOnly, vehicleMatchRequest.ReserveEndDate)
	if err != nil {
		return fiber.NewError(fiber.StatusBadRequest, err.Error())
	}
	if reserveStartDate.Before(time.Now()) {
		return fiber.NewError(fiber.StatusBadRequest, "start date can't be sooner than now")
	}
	if reserveEndDate.Before(reserveStartDate) {
		return fiber.NewError(fiber.StatusBadRequest, "end date can't be sooner than start date")
	}
	vehicleUUID, err := uuid.Parse(vehicleMatchRequest.TripID)
	reservationID, vehicle, err := h.service.MatchVehicle(c.Context(), &domain.MatchMakerRequest{
		TripID:             vehicleUUID,
		ReserveStartDate:   reserveStartDate,
		ReserveEndDate:     reserveEndDate,
		TripDistance:       vehicleMatchRequest.TripDistance,
		NumberOfPassengers: vehicleMatchRequest.NumberOfPassengers,
		TripType:           domain.VehicleType(vehicleMatchRequest.TripType),
		MaxPrice:           vehicleMatchRequest.MaxPrice,
		YearOfManufacture:  vehicleMatchRequest.YearOfManufacture,
	})
	if err != nil {
		if errors.Is(err, vehicleService.ErrVehicleNotFound) {
			return fiber.NewError(fiber.StatusNotFound, err.Error())
		}
		return fiber.NewError(fiber.StatusInternalServerError,err.Error())
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
	userUUID := c.Locals("UserUUID")
	vehicle.OwnerID = userUUID.(uuid.UUID)

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

func (h *VehicleHandler) UpdateVehicle(c *fiber.Ctx) error {
	log.Println("Incoming request to create vehicle")
	vehicleId := c.Params("id")
	vehicleUUID, err := uuid.Parse(vehicleId)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid vehicle id",
		})
	}
	var vehicle presenter.Vehicle
	if err := c.BodyParser(&vehicle); err != nil {
		log.Printf("Error parsing body: %v", err)
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid request payload",
		})
	}

	log.Printf("Vehicle parsed: %+v", vehicle)
	domainVehicle := domain.Vehicle{
		ID:                vehicleUUID,
		Capacity:          vehicle.Capacity,
		Speed:             vehicle.Speed,
		Status:            string(vehicle.Status),
		PricePerKilometer: vehicle.PricePerKilometer,
	}
	if err := h.service.UpdateVehicle(c.Context(), &domainVehicle); err != nil {
		log.Printf("Error creating vehicle: %v", err)
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	log.Printf("Vehicle created: %+v", vehicle)
	return c.Status(fiber.StatusCreated).JSON("vehicle updated successfully")
}

func RegisterVehicleRoutes(app *fiber.App, service port.VehicleService, cfg config.Config) {
	handler := NewVehicleHandler(service)
	app.Post("/api/v1/vehicles", newAuthMiddleware([]byte(cfg.Server.Secret)), handler.CreateVehicle)
	app.Patch("/api/v1/vehicles/:id", newAuthMiddleware([]byte(cfg.Server.Secret)), handler.UpdateVehicle)
	app.Post("/api/v1/vehicles/match", handler.MatchVehicle)
}
