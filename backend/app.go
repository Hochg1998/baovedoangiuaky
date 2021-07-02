package main

import (
	"github.com/TechMaster/golang/15GoMySQL/config"
	"github.com/TechMaster/golang/15GoMySQL/repo"
	"github.com/TechMaster/golang/15GoMySQL/routes"
	_ "github.com/go-sql-driver/mysql"
	"github.com/gofiber/fiber/v2"

	"github.com/gofiber/fiber/v2/middleware/cors"
)

func main() {
	err := config.LoadConfig()
	if err != nil {
		panic(err.Error())
	}
	repo.Connect(config.Config)
	repo.InitMasterData()

	app := fiber.New(fiber.Config{
		Prefork:       false,
		CaseSensitive: true,
		StrictRouting: true,
	})

	// Default config
	app.Use(cors.New())

	// Or extend your config for customization
	app.Use(cors.New(cors.Config{
		AllowOrigins: "http://localhost:8080",
		AllowHeaders: "Origin, Content-Type, Accept",
	}))

	app.Static("/public", "./public", fiber.Static{ //http://localhost:3000/public OR http://localhost:3000/public/dog.jpeg
		Compress:  true,
		ByteRange: true,
		Browse:    true,
		MaxAge:    3600,
	})

	routes.ConfigRouter(app)

	app.Listen(":3000")
}
