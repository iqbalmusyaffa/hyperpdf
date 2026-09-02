package handler

import (
	"time"

	"cobagolang/backend/internal/compressor"
	"cobagolang/backend/internal/dto"

	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

var startTime = time.Now()

// HealthHandler handles health check endpoint
type HealthHandler struct {
	db         *gorm.DB
	compressor compressor.PDFCompressor
}

// NewHealthHandler creates a new HealthHandler instance
func NewHealthHandler(db *gorm.DB, comp compressor.PDFCompressor) *HealthHandler {
	return &HealthHandler{
		db:         db,
		compressor: comp,
	}
}

// CheckHealth returns system status, database connectivity, and PDF engine status
// @Summary System Health Check
// @Description Check service uptime, database connection, and PDF engine availability
// @Tags Health
// @Produce json
// @Success 200 {object} dto.APIResponse
// @Failure 503 {object} dto.APIErrorResponse
// @Router /health [get]
func (h *HealthHandler) CheckHealth(c *fiber.Ctx) error {
	dbStatus := "connected"
	sqlDB, err := h.db.DB()
	if err != nil || sqlDB.Ping() != nil {
		dbStatus = "disconnected"
	}

	engineStatus := "available"
	if !h.compressor.IsAvailable() {
		engineStatus = "unavailable (binary not found)"
	}

	statusData := fiber.Map{
		"status":          "ok",
		"uptime":          time.Since(startTime).String(),
		"database":        dbStatus,
		"pdf_engine":      h.compressor.Name(),
		"engine_status":   engineStatus,
		"engine_bin_path": h.compressor.GetBinaryPath(),
	}

	if dbStatus == "disconnected" {
		return c.Status(fiber.StatusServiceUnavailable).JSON(
			dto.NewErrorResponse("Database service is unavailable", "Database ping failed"),
		)
	}

	return c.Status(fiber.StatusOK).JSON(
		dto.NewSuccessResponse("System is healthy", statusData),
	)
}
