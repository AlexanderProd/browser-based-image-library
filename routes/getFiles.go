package routes

import (
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"

	"github.com/AlexanderProd/browser-based-image-library/database"
)

func GetFiles(c *fiber.Ctx, db *gorm.DB) error {
	body := struct {
		ParentID string `json:"parentID"`
	}{}

	if err := c.BodyParser(&body); err != nil {
		return err
	}

	var files []database.File
	db.Where("parent_id = ?", body.ParentID).Preload("Tags").Preload("Categories").Find(&files)

	return c.JSON(files)
}