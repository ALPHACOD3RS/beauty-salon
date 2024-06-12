package v1

import (
	"github.com/ALPHACOD3RS/Beauty-Salon/internal/models"
	"github.com/ALPHACOD3RS/Beauty-Salon/internal/utils"
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
	"gorm.io/gorm"
)


func CreateServiceHandler(c *fiber.Ctx,  db *gorm.DB) error{
	service := new(models.Service)
	userInfo := utils.GetUserInfoFromJWT(c)
    
    if userInfo.Role != "admin" {
        return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
            "msg": "Unauthorized. Only admins can create services.",
        })
    }

	uuID := uuid.New().String()
	

	if err := c.BodyParser(&service); err != nil{
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"msg": "Something went wrong",
		})
	}

	service.ServiceID = uuID

	if err := db.Create(&service).Error; err != nil{
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"msg":"Couldnt create the service",
		})
	}

	return c.Status(fiber.StatusCreated).JSON(service)
}


func GetAllServiceHandler(c *fiber.Ctx, db *gorm.DB) error{
	var service []models.Service

	db.Find(&service)

	return c.Status(fiber.StatusOK).JSON(service)
}

func GetServiceByIdHandler(c *fiber.Ctx, db *gorm.DB) error{
	sID := c.Params("id")

	service := new(models.Service)

	if err := db.Where("service_id = ?", sID).First(&service).Error; err != nil{
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"msg": "Couldn't get the service",
			"err": err.Error(),	
		})
	}

	return c.Status(fiber.StatusOK).JSON(service)
}


func UpdateServiceHandler(c *fiber.Ctx, db *gorm.DB) error {
	
	sID := c.Params("id")
	userInfo := utils.GetUserInfoFromJWT(c)
    
    if userInfo.Role != "admin" {
        return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
            "msg": "Unauthorized. Only admins can create services.",
        })
    }
	var service models.Service


	if err := db.Where("service_id = ?", sID).First(&service).Error; err != nil{
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"msg": "couldnt get the service",
		})
	}

	if err := c.BodyParser(&service); err != nil{
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"msg": "something went wrong",
			"err": err.Error(),
		})
	}

	db.Save(&service)

	return c.Status(fiber.StatusCreated).JSON(service)

}


func DeleteServiceHandler(c *fiber.Ctx, db *gorm.DB) error{
	sID := c.Params("id")

	userInfo := utils.GetUserInfoFromJWT(c)
    
    if userInfo.Role != "admin" {
        return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
            "msg": "Unauthorized. Only admins can create services.",
        })
    }
	
	var service models.Service

	if err := db.Where("service_id = ?", sID).First(&service).Error; err != nil{
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"msg": "couldnt get the service",
		})
	}

	db.Delete(&service)

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"msg": "Service deleted successfully",
	})
}