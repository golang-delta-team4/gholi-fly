package http

import (
	"fmt"

	"gholi-fly-hotel/app"
	"gholi-fly-hotel/config"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/recover"
)

func Run(appContainer app.App, cfg config.ServerConfig) error {
	router := fiber.New()
	router.Use(recover.New())
	api := router.Group("/api/v1/hotel",
		setUserContext,
		newAuthMiddleware([]byte(cfg.Secret)),
	)
	registerGlobalRoutes(api)

	registerHotelAPI(appContainer, api)
	registerRoomAPI(appContainer, api)
	registerBookingAPI(appContainer, api)
	registerStaffAPI(appContainer, api)

	return router.Listen(fmt.Sprintf(":%d", cfg.HttpPort))
}

func registerGlobalRoutes(router fiber.Router) {
	router.Get("/health", func(c *fiber.Ctx) error {
		return c.Status(fiber.StatusOK).JSON(fiber.Map{
			"status":  "OK",
			"message": "gholi-hotels-api is running",
		})
	})
}

func registerHotelAPI(appContainer app.App, router fiber.Router) {
	hotelSvcGetter := hotelServiceGetter(appContainer)
	router.Post("/", setTransaction(appContainer.DB()), CreateHotel(hotelSvcGetter))
	router.Get("/", setTransaction(appContainer.DB()), GetAllHotels(hotelSvcGetter))
	router.Get("/owner/:owner_id", setTransaction(appContainer.DB()), GetAllHotelsByOwnerID(hotelSvcGetter))
	router.Get("/:id", setTransaction(appContainer.DB()), GetHotelByID(hotelSvcGetter))
	// router.Put("/:id", setTransaction(appContainer.DB()), UpdateHotelByID(hotelSvcGetter))
	// router.Delete("/:id", setTransaction(appContainer.DB()), DeleteHotelByID(hotelSvcGetter))
}

func registerRoomAPI(appContainer app.App, router fiber.Router) {
	roomlSvcGetter := roomServiceGetter(appContainer)
	roomRouter := router.Group("/room")
	roomRouter.Post("/:hotel_id", setTransaction(appContainer.DB()), CreateRoomByHotelID(roomlSvcGetter))
	roomRouter.Get("/hotel/:hotel_id", setTransaction(appContainer.DB()), GetAllRoomsByHotelID(roomlSvcGetter))
	roomRouter.Get("/:id", setTransaction(appContainer.DB()), GetRoomByID(roomlSvcGetter))
	// roomRouter.Put("/:id", setTransaction(appContainer.DB()), UpdateRoomByID(roomlSvcGetter))
	roomRouter.Delete("/:id", setTransaction(appContainer.DB()), DeleteRoomByID(roomlSvcGetter))
}

func registerBookingAPI(appContainer app.App, router fiber.Router) {
	bookingSvcGetter := bookingServiceGetter(appContainer)
	bookingRouter := router.Group("/booking")
	bookingRouter.Post("user/:hotel_id", setTransaction(appContainer.DB()), CreateUserBookingByHotelID(bookingSvcGetter))
	bookingRouter.Post("/:hotel_id", setTransaction(appContainer.DB()), CreateBookingByHotelID(bookingSvcGetter))
	bookingRouter.Get("/room/:room_id", setTransaction(appContainer.DB()), GetAllBookingsByRoomID(bookingSvcGetter))
	// bookingRouter.Get("/user/:user_id", setTransaction(appContainer.DB()), GetAllBookingsByUserID(hotelSvcGetter))
	// bookingRouter.Get("/:id", setTransaction(appContainer.DB()), GetBookingByID(hotelSvcGetter))
	// bookingRouter.Put("/:id", setTransaction(appContainer.DB()), UpdateBookingByID(hotelSvcGetter))
	// bookingRouter.Delete("/:id", setTransaction(appContainer.DB()), DeleteBookingByID(hotelSvcGetter))
}

func registerStaffAPI(appContainer app.App, router fiber.Router) {
	stafflSvcGetter := staffServiceGetter(appContainer)
	staffRouter := router.Group("/staff")
	staffRouter.Post("/:hotel_id", setTransaction(appContainer.DB()), CreateStaffByHotelID(stafflSvcGetter))
	staffRouter.Get("/hotel/:hotel_id", setTransaction(appContainer.DB()), GetAllStaffsByHotelID(stafflSvcGetter))
	staffRouter.Get("/:id", setTransaction(appContainer.DB()), GetStaffByID(stafflSvcGetter))
	// staffRouter.Put("/:id", setTransaction(appContainer.DB()), UpdateStaffByID(stafflSvcGetter))
	// staffRouter.Delete("/:id", setTransaction(appContainer.DB()), DeleteStaffByID(stafflSvcGetter))
}
