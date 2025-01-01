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
	)
	registerGlobalRoutes(api)

	registerHotelAPI(appContainer, api, cfg)
	registerRoomAPI(appContainer, api, cfg)
	registerBookingAPI(appContainer, api, cfg)
	registerStaffAPI(appContainer, api, cfg)

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

func registerHotelAPI(appContainer app.App, router fiber.Router, cfg config.ServerConfig) {
	hotelSvcGetter := hotelServiceGetter(appContainer)
	router.Post("/", setTransaction(appContainer.DB()), newAuthMiddleware([]byte(cfg.Secret)), CreateHotel(hotelSvcGetter))
	router.Get("/", setTransaction(appContainer.DB()), newAuthMiddleware([]byte(cfg.Secret)), GetAllHotels(hotelSvcGetter))
	router.Get("/owner/:owner_id", setTransaction(appContainer.DB()), newAuthMiddleware([]byte(cfg.Secret)), GetAllHotelsByOwnerID(hotelSvcGetter))
	router.Get("/:id", setTransaction(appContainer.DB()), newAuthMiddleware([]byte(cfg.Secret)), GetHotelByID(hotelSvcGetter))
	// router.Put("/:id", setTransaction(appContainer.DB()), UpdateHotelByID(hotelSvcGetter))
	// router.Delete("/:id", setTransaction(appContainer.DB()), DeleteHotelByID(hotelSvcGetter))
}

func registerRoomAPI(appContainer app.App, router fiber.Router, cfg config.ServerConfig) {
	roomlSvcGetter := roomServiceGetter(appContainer)
	roomRouter := router.Group("/room")
	roomRouter.Post("/:hotel_id", setTransaction(appContainer.DB()), newAuthMiddleware([]byte(cfg.Secret)), CreateRoomByHotelID(roomlSvcGetter))
	roomRouter.Get("/hotel/:hotel_id", setTransaction(appContainer.DB()), newAuthMiddleware([]byte(cfg.Secret)), GetAllRoomsByHotelID(roomlSvcGetter))
	roomRouter.Get("/:id", setTransaction(appContainer.DB()), newAuthMiddleware([]byte(cfg.Secret)), GetRoomByID(roomlSvcGetter))
	// roomRouter.Put("/:id", setTransaction(appContainer.DB()),newAuthMiddleware([]byte(cfg.Secret)), UpdateRoomByID(roomlSvcGetter))
	roomRouter.Delete("/:id", setTransaction(appContainer.DB()), newAuthMiddleware([]byte(cfg.Secret)), DeleteRoomByID(roomlSvcGetter))
}

func registerBookingAPI(appContainer app.App, router fiber.Router, cfg config.ServerConfig) {
	bookingSvcGetter := bookingServiceGetter(appContainer)
	bookingRouter := router.Group("/booking")
	bookingRouter.Post("user/:hotel_id", setTransaction(appContainer.DB()), newAuthMiddleware([]byte(cfg.Secret)), CreateUserBookingByHotelID(bookingSvcGetter))
	bookingRouter.Post("/:hotel_id", setTransaction(appContainer.DB()), CreateBookingByHotelID(bookingSvcGetter))
	bookingRouter.Get("/room/:room_id", setTransaction(appContainer.DB()), newAuthMiddleware([]byte(cfg.Secret)), GetAllBookingsByRoomID(bookingSvcGetter))
	bookingRouter.Get("/user/", setTransaction(appContainer.DB()), newAuthMiddleware([]byte(cfg.Secret)), GetAllBookingsByUserID(bookingSvcGetter))
	bookingRouter.Get("/:id", setTransaction(appContainer.DB()), newAuthMiddleware([]byte(cfg.Secret)), GetBookingByID(bookingSvcGetter))
	bookingRouter.Patch("/cancel/user/:factor_id", setTransaction(appContainer.DB()), newAuthMiddleware([]byte(cfg.Secret)), CancelUserBookingByID(bookingSvcGetter))
	bookingRouter.Patch("/cancel/:factor_id", setTransaction(appContainer.DB()), CancelBookingByID(bookingSvcGetter))
	bookingRouter.Patch("/approve/user/:factor_id", setTransaction(appContainer.DB()), newAuthMiddleware([]byte(cfg.Secret)), ApproveUserBookingByID(bookingSvcGetter))
	// bookingRouter.Put("/:id", setTransaction(appContainer.DB()), newAuthMiddleware([]byte(cfg.Secret)), UpdateBookingByID(bookingSvcGetter))
	bookingRouter.Delete("/:id", setTransaction(appContainer.DB()), newAuthMiddleware([]byte(cfg.Secret)), DeleteBookingByID(bookingSvcGetter))
}

func registerStaffAPI(appContainer app.App, router fiber.Router, cfg config.ServerConfig) {
	stafflSvcGetter := staffServiceGetter(appContainer)
	staffRouter := router.Group("/staff")
	staffRouter.Post("/:hotel_id", setTransaction(appContainer.DB()), newAuthMiddleware([]byte(cfg.Secret)), CreateStaffByHotelID(stafflSvcGetter))
	staffRouter.Get("/hotel/:hotel_id", setTransaction(appContainer.DB()), newAuthMiddleware([]byte(cfg.Secret)), GetAllStaffsByHotelID(stafflSvcGetter))
	staffRouter.Get("/:id", setTransaction(appContainer.DB()), newAuthMiddleware([]byte(cfg.Secret)), GetStaffByID(stafflSvcGetter))
	// staffRouter.Put("/:id", setTransaction(appContainer.DB()), UpdateStaffByID(stafflSvcGetter))
	// staffRouter.Delete("/:id", setTransaction(appContainer.DB()), DeleteStaffByID(stafflSvcGetter))
}
