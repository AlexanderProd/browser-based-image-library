package routes

import (
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"

	"github.com/AlexanderProd/browser-based-image-library/database"
)

func Shuffle(c *fiber.Ctx, db *gorm.DB) error {
	body := struct {
		Limit int 	 `json:"limit"`
		Type  string `json:"type"`
	}{
		Limit: 10,
		Type: "file",
	}

	if err := c.BodyParser(&body); err != nil {
		return err
	}

	var files []database.File
	db.Where(&database.File{Type: "file"}).Order("random()").Preload("Tags").Preload("Categories").Limit(body.Limit).Find(&files)

	return c.JSON(files)
}