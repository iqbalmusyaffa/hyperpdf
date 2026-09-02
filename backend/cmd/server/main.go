package main

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"
	"time"

	_ "cobagolang/backend/docs"
	"cobagolang/backend/internal/compressor"
	"cobagolang/backend/internal/config"
	"cobagolang/backend/internal/database"
	"cobagolang/backend/internal/dto"
	"cobagolang/backend/internal/handler"
	"cobagolang/backend/internal/middleware"
	"cobagolang/backend/internal/repository"
	"cobagolang/backend/internal/service"
	"cobagolang/backend/internal/utils"

	"github.com/gofiber/fiber/v2"
	"github.com/sirupsen/logrus"
	fiberSwagger "github.com/swaggo/fiber-swagger"
)

// @title HyperPDF API
// @version 1.0
// @description Next-Gen High-Performance PDF Suite API (Compress, Merge, Split) powered by Go, Fiber, Ghostscript, and PostgreSQL.
// @termsOfService http://swagger.io/terms/

// @contact.name HyperPDF Support
// @contact.email support@hyperpdf.local

// @license.name MIT
// @license.url https://opensource.org/licenses/MIT

// @host localhost:8080
// @BasePath /
// @schemes http https
func main() {
	// 1. Load Configuration
	cfg, err := config.LoadConfig("")
	if err != nil {
		fmt.Printf("Fatal: failed to load configuration: %v\n", err)
		os.Exit(1)
	}

	// 2. Initialize Structured Logger
	log := utils.InitLogger(cfg.AppEnv)
	log.WithFields(logrus.Fields{
		"env":      cfg.AppEnv,
		"port":     cfg.AppPort,
		"db_host":  cfg.DatabaseHost,
		"max_size": fmt.Sprintf("%d MB", cfg.MaxFileSizeMB),
	}).Info("Starting PDF Compressor Service")

	// 3. Connect to Database
	db, err := database.InitDB(cfg)
	if err != nil {
		log.WithError(err).Warn("PostgreSQL connection failed. If running outside Docker, start PostgreSQL or use docker-compose.")
	} else {
		// Run migrations
		if err := database.RunMigrations(cfg, "migrations"); err != nil {
			log.WithError(err).Warn("golang-migrate failed, attempting GORM AutoMigrate as fallback...")
			if autoErr := database.AutoMigrate(db); autoErr != nil {
				log.WithError(autoErr).Error("AutoMigrate also failed")
			}
		}
	}

	// 4. Initialize Compressor Engine
	pdfComp, err := compressor.NewCompressor(cfg.CompressorEngine, cfg.GhostscriptBinary, cfg.QPDFBinary)
	if err != nil {
		log.WithError(err).Fatal("Failed to initialize PDF compression engine")
	}

	if pdfComp.IsAvailable() {
		log.WithFields(logrus.Fields{
			"engine": pdfComp.Name(),
			"path":   pdfComp.GetBinaryPath(),
		}).Info("PDF compression engine is ready")
	} else {
		log.WithField("engine", pdfComp.Name()).Warn(
			"PDF engine binary not found in PATH. Install Ghostscript or run in Docker.",
		)
	}

	// 5. Initialize Layers (Repo -> Service -> Handler)
	var jobRepo repository.JobRepository
	if db != nil {
		jobRepo = repository.NewJobRepository(db)
	}

	pdfService := service.NewPDFService(jobRepo, pdfComp, cfg)
	pdfHandler := handler.NewPDFHandler(pdfService)
	healthHandler := handler.NewHealthHandler(db, pdfComp)

	// 6. Initialize Fiber App
	app := fiber.New(fiber.Config{
		AppName:      "PDF Compressor API",
		BodyLimit:    cfg.MaxFileSizeMB * 1024 * 1024,
		ReadTimeout:  time.Duration(cfg.CompressTimeoutSeconds+30) * time.Second,
		WriteTimeout: time.Duration(cfg.CompressTimeoutSeconds+30) * time.Second,
		ErrorHandler: func(c *fiber.Ctx, err error) error {
			code := fiber.StatusInternalServerError
			if e, ok := err.(*fiber.Error); ok {
				code = e.Code
			}
			return c.Status(code).JSON(dto.NewErrorResponse(err.Error(), "Request processing failed"))
		},
	})

	// 7. Register Global Middlewares
	app.Use(middleware.Recover())
	app.Use(middleware.CORS())
	app.Use(middleware.RequestLogger())
	app.Use(middleware.RateLimiter(cfg))

	// 8. Register Routes
	app.Get("/", func(c *fiber.Ctx) error {
		return c.Status(fiber.StatusOK).JSON(dto.NewSuccessResponse("HyperPDF API is running", fiber.Map{
			"service":      "HyperPDF API",
			"version":      "1.0.0",
			"frontend_url": "http://localhost:3000",
			"swagger_docs": "http://localhost:8080/swagger/index.html",
			"health_check": "http://localhost:8080/health",
			"endpoints": fiber.Map{
				"compress": "POST /api/v1/pdf/compress",
				"merge":    "POST /api/v1/pdf/merge",
				"split":    "POST /api/v1/pdf/split",
				"get_job":  "GET /api/v1/pdf/jobs/:id",
				"download": "GET /api/v1/pdf/jobs/:id/download",
				"delete":   "DELETE /api/v1/pdf/jobs/:id",
			},
		}))
	})

	app.Get("/health", healthHandler.CheckHealth)

	// Swagger documentation route
	app.Get("/swagger/*", fiberSwagger.WrapHandler)

	api := app.Group("/api/v1")
	pdfGroup := api.Group("/pdf")
	{
		pdfGroup.Post("/compress", pdfHandler.CompressPDF)
		pdfGroup.Post("/merge", pdfHandler.MergePDF)
		pdfGroup.Post("/split", pdfHandler.SplitPDF)
		pdfGroup.Get("/jobs/:id", pdfHandler.GetJob)
		pdfGroup.Get("/jobs/:id/download", pdfHandler.DownloadCompressedPDF)
		pdfGroup.Delete("/jobs/:id", pdfHandler.DeleteJob)
	}

	// 404 Fallback Handler
	app.Use(func(c *fiber.Ctx) error {
		return c.Status(fiber.StatusNotFound).JSON(
			dto.NewErrorResponse("Endpoint not found", fmt.Sprintf("Cannot %s %s", c.Method(), c.Path())),
		)
	})

	// 9. Graceful Shutdown Setup
	serverShutdown := make(chan os.Signal, 1)
	signal.Notify(serverShutdown, os.Interrupt, syscall.SIGTERM)

	go func() {
		<-serverShutdown
		log.Info("Shutting down PDF Compressor server gracefully...")
		if err := app.ShutdownWithTimeout(10 * time.Second); err != nil {
			log.WithError(err).Error("Error while shutting down server")
		}
	}()

	// 10. Start Server
	listenAddr := fmt.Sprintf(":%s", cfg.AppPort)
	log.Infof("Server listening on %s (http://localhost:%s)", listenAddr, cfg.AppPort)
	if err := app.Listen(listenAddr); err != nil {
		log.WithError(err).Info("Server stopped listening")
	}
}
