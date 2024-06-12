package v1

import (
	model "github.com/ALPHACOD3RS/Beauty-Salon/internal/models"
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
	"gorm.io/gorm"
)


func CreateAppointmentHandler(c *fiber.Ctx, db *gorm.DB) error {
    appointment := new(model.Appointment)
    uuId := uuid.New().String()

    if err := c.BodyParser(&appointment); err != nil {
        return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
            "msg": "Invalid data",
            "err": err.Error(),
        })
    }

    appointment.AppointmentID = uuId

    var user model.User
    if err := db.Where("user_id = ?", appointment.UserID).First(&user).Error; err != nil {
        return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
            "msg": "User not found",
            "err": err.Error(),
        })
    }

    // Check if the service exists
    var service model.Service
    if err := db.Where("service_id = ?", appointment.ServiceID).First(&service).Error; err != nil {
        return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
            "msg": "Service not found",
            "err": err.Error(),
        })
    }

    appointment.User = user
    appointment.Service = service

    if err := db.Create(&appointment).Error; err != nil {
        return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
            "msg": "Could not create appointment",
            "err": err.Error(),
        })
    }

    return c.Status(fiber.StatusCreated).JSON(appointment)
}


func GetAllAppointmentsHandler(c *fiber.Ctx, db *gorm.DB) error{
	var appointments []model.Appointment

	if err := db.Preload("User").Preload("Service").Find(&appointments).Error; err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Failed to fetch appointments",
		})
	}

	return c.Status(fiber.StatusOK).JSON(appointments)

}
