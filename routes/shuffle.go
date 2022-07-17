package routes

import (
	"strconv"

	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"

	"github.com/AlexanderProd/browser-based-image-library/database"
)

func Shuffle(c *fiber.Ctx, db *gorm.DB) error {
	limit, _ := strconv.Atoi(c.Query("limit", "10")) 

	var files []database.File
	db.Where(&database.File{Type: "file"}).Order("random()").Preload("Tags").Preload("Categories").Limit(limit).Find(&files)

	return c.Status(fiber.StatusOK).JSON(files)
}