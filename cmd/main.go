package main

import (
	"log"
	"mf-loan/config"
	"mf-loan/delivery/http"
	"mf-loan/infra"
	"mf-loan/repository"
	"mf-loan/usecase"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/helmet"
	"github.com/gofiber/fiber/v2/middleware/idempotency"
	"github.com/gofiber/fiber/v2/middleware/limiter"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/fiber/v2/middleware/monitor"
	"github.com/gofiber/fiber/v2/middleware/recover"
	"github.com/gofiber/swagger"
)

// @title Fintech API
// @version 1.0
// @description This is a swagger for Fintech API
// @termsOfService http://swagger.io/terms/
// @contact.name API Support
// @contact.email fiber@swagger.io
// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html
// @host localhost:8080
// @BasePath /
func main() {
	config.LoadEnv()
	app := fiber.New()

	// Security Settings
	var HelmetConfig = helmet.Config{
		XSSProtection:             "1",
		ContentTypeNosniff:        "nosniff",
		XFrameOptions:             "SAMEORIGIN",
		ReferrerPolicy:            "no-referrer",
		CrossOriginEmbedderPolicy: "require-corp",
		CrossOriginOpenerPolicy:   "same-origin",
		CrossOriginResourcePolicy: "same-origin",
		OriginAgentCluster:        "?1",
		XDNSPrefetchControl:       "off",
		XDownloadOptions:          "noopen",
		XPermittedCrossDomain:     "none",
	}
	var CorsConfig = cors.Config{
		AllowOrigins: "*", // change this on production server
		AllowHeaders: "Origin, Content-Type, Accept",
	}
	var LimiterConfig = limiter.Config{
		Max:               20,
		Expiration:        30 * time.Second,
		LimiterMiddleware: limiter.SlidingWindow{},
	}
	var IdemConfig = idempotency.Config{
		Lifetime: 5 * time.Second,
	}

	// Setup Middlewares
	app.Use(recover.New())                            // Improve Server availability
	app.Use(idempotency.New(IdemConfig))              // Improve Server availability in unstable network
	app.Use(logger.New())                             // Improve logger (OWASP 10)
	app.Use(helmet.New(HelmetConfig))                 // Improve security (OWASP 10)
	app.Use(cors.New(CorsConfig))                     // Improve security (OWASP 10)
	app.Use(limiter.New(LimiterConfig))               // Improve security (OWASP 10)
	app.Get("/swagger/*", swagger.New(swagger.Config{ // Swagger
		URL:          "/docs/swagger.json",
		DeepLinking:  false,
		DocExpansion: "none",
	}))
	app.Get("/metrics", monitor.New(monitor.Config{ // Add Performance Matrix page
		Title: "Loan Engine Performance Metrics Page",
	}))
	app.Static("/docs", "./docs")

	db := infra.InitDB()

	// Customer components
	customerRepo := repository.NewCustomerRepository(db)
	customerUseCase := usecase.NewCustomerUseCase(customerRepo)
	http.NewCustomerHandler(app, customerUseCase)

	// Tenor components
	tenorRepo := repository.NewTenorRepository(db)
	tenorUseCase := usecase.NewTenorUseCase(tenorRepo, customerRepo)
	http.NewTenorHandler(app, tenorUseCase)

	// Transaction components
	transactionRepo := repository.NewTransactionRepository(db)
	transactionUseCase := usecase.NewTransactionUseCase(transactionRepo, customerRepo)
	http.NewTransactionHandler(app, transactionUseCase)

	log.Fatal(app.Listen(":8080"))
}
