package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	"embed"

	"github.com/AlexanderProd/browser-based-image-library/database"
	"github.com/AlexanderProd/browser-based-image-library/routes"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/filesystem"
	"gorm.io/gorm"
)


const BATCH_INSERT_SIZE = 100

var (
	db *gorm.DB
	app *fiber.App
	rootPath string
	port = flag.String("port", "3000", "Port to listen on")
	walk = flag.Bool("walk", true, "If the root path should be scanned.")
	dbPath = flag.String("database", "./database.db", "Path of the database file")
	allowedFileTypes = []string{".png", ".jpg", ".jpeg", ".bmp"}
)

//go:embed frontend/build/*
var embedDirStatic embed.FS


func main() {
	flag.Parse()
	if flag.NArg() == 0 { 
    flag.Usage()
    os.Exit(1)
	}
	rootPath = flag.Arg(0)

	db = database.Connect(*dbPath)

	if (*walk) {
		start := time.Now()
		go walkDir()
		elapsed := time.Since(start)
		fmt.Println("Ran in", elapsed)
	}

	app = fiber.New()
	app.Use(cors.New(cors.Config{
    AllowOrigins: "http://localhost:3000",
    AllowHeaders: "Origin, Content-Type, Accept",
		AllowCredentials: true,
	}))

	app.Use("/", filesystem.New(filesystem.Config{
		Root: http.FS(embedDirStatic),
		PathPrefix: "frontend/build",
		Browse: true,
	}))
	app.Static("/static", rootPath)

	app.Use("/api/files", func (c *fiber.Ctx) error {
		return routes.GetFiles(c, db)
	})
	app.Use("/api/shuffle", func (c *fiber.Ctx) error {
		return routes.Shuffle(c, db)
	})

	log.Fatal(app.Listen(":"+*port))	
}