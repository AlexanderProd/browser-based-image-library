package main

import (
	"fmt"
	"net/http"
	"time"

	"embed"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/filesystem"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)


const PATH = "/Users/alexanderhoerl/Downloads/Dummy_Folder"
const BATCH_INSERT_SIZE = 200
var allowedFileTypes = []string{".png", ".jpg", ".jpeg", ".bmp"}
var db *gorm.DB
var app *fiber.App

//go:embed frontend/dist/*
var embedDirStatic embed.FS

func main() {
	var err error
	db, err = gorm.Open(sqlite.Open("./database.db"), &gorm.Config{Logger: logger.Default.LogMode(logger.Silent)})
  if err != nil {
    panic("failed to connect database")
  }

	db.AutoMigrate(&File{}, &FilePath{}, &Tag{})

	start := time.Now()
	
	walkDir()
	/* var file File
	db.Preload("FilePath").First(&file, "id = ?", "e856f6ea521daf0272fe8ebee027680a3ec0ccba7a5deefb6bd0ae08c42dd5c3")
	fmt.Println(file.Type)
	fmt.Println(file.FilePath.Path)

	tag := Tag{Name: "test"}
	db.Clauses(clause.OnConflict{DoNothing: true}).Create(&tag)
	db.Model(&file).Association("Tags").Append(&tag) */

	elapsed := time.Since(start)
	fmt.Println("Ran in", elapsed)

	app = fiber.New()

	app.Use("/", filesystem.New(filesystem.Config{
		Root: http.FS(embedDirStatic),
		PathPrefix: "frontend/dist",
		Browse: true,
	}))

	//app.Listen(":3000")
}