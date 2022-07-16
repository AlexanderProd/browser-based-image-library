package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
	"time"

	"embed"

	"github.com/AlexanderProd/browser-based-image-library/database"
	"github.com/AlexanderProd/browser-based-image-library/routes"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/filesystem"
	"gorm.io/gorm"
)


const PATH = "/Users/alexanderhoerl/Downloads/Dummy_Folder"
const BATCH_INSERT_SIZE = 100

var (
	db *gorm.DB
	app *fiber.App
	port = flag.String("port", "3000", "Port to listen on")
	dbPath = flag.String("database", "./database.db", "Path of the database file")
	allowedFileTypes = []string{".png", ".jpg", ".jpeg", ".bmp"}
)

//go:embed frontend/build/*
var embedDirStatic embed.FS


func main() {
	flag.Parse()

	db = database.Connect(*dbPath)

	start := time.Now()
	
	walkDir()

	elapsed := time.Since(start)
	fmt.Println("Ran in", elapsed)

	app = fiber.New()

	app.Use("/", filesystem.New(filesystem.Config{
		Root: http.FS(embedDirStatic),
		PathPrefix: "frontend/build",
		Browse: true,
	}))
	app.Static("/static", PATH)

	app.Use("/api/files", func (c *fiber.Ctx) error {
		return routes.GetFiles(c, db)
	})
	app.Use("/api/shuffle", func (c *fiber.Ctx) error {
		return routes.Shuffle(c, db)
	})

	log.Fatal(app.Listen(":"+*port))	
}