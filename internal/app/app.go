package app

import (
	"github.com/aws/aws-sdk-go-v2/service/s3"
	"github.com/nurgal1ev/booking-service/internal/config"
	"github.com/nurgal1ev/booking-service/internal/infrastructure/seaweedfs"
	bookingRepo "github.com/nurgal1ev/booking-service/internal/repository/booking"
	"github.com/nurgal1ev/booking-service/internal/repository/property"
	unitRepo "github.com/nurgal1ev/booking-service/internal/repository/unit"
	userRepo "github.com/nurgal1ev/booking-service/internal/repository/user"
	bookingService "github.com/nurgal1ev/booking-service/internal/service/booking"
	propertyService "github.com/nurgal1ev/booking-service/internal/service/property"
	unitService "github.com/nurgal1ev/booking-service/internal/service/unit"
	userService "github.com/nurgal1ev/booking-service/internal/service/user"
	"github.com/nurgal1ev/booking-service/internal/transport/httpv1"
	bookingHandler "github.com/nurgal1ev/booking-service/internal/transport/httpv1/handler/booking"
	propertyHandler "github.com/nurgal1ev/booking-service/internal/transport/httpv1/handler/property"
	unitHandler "github.com/nurgal1ev/booking-service/internal/transport/httpv1/handler/unit"
	userHandler "github.com/nurgal1ev/booking-service/internal/transport/httpv1/handler/user"
	"gorm.io/gorm"
)

func InitHandlers(db *gorm.DB, cfg *config.Config, s3Client *s3.Client) httpv1.Handlers {
	filestorage := seaweedfs.NewFileStorage(s3Client, cfg.S3Config.S3Endpoint, cfg.S3Config.S3PublicEndpoint)

	userRepository := userRepo.NewUserRepo(db)
	userService := userService.NewUserService(userRepository, cfg.Auth.Secret)
	userHandler := userHandler.NewUserHandler(userService, filestorage)

	propertyRepository := property.NewPropertyRepo(db)
	propertyService := propertyService.NewPropertyService(propertyRepository, userRepository)
	propertyHandler := propertyHandler.NewPropertyHandler(propertyService)

	unitRepo := unitRepo.NewUnitRepo(db)
	unitService := unitService.NewUnitService(unitRepo, propertyRepository)
	unitHandler := unitHandler.NewUnitHandler(unitService)

	bookingRepo := bookingRepo.NewBookingRepo(db)
	bookingService := bookingService.NewBookingService(bookingRepo, unitRepo)
	bookingHandler := bookingHandler.NewBookingHandler(bookingService)

	return httpv1.Handlers{
		User:     userHandler,
		Property: propertyHandler,
		Unit:     unitHandler,
		Booking:  bookingHandler,
	}
}
