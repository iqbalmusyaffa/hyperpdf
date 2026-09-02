package middleware

import (
	"time"

	"cobagolang/backend/internal/config"
	"cobagolang/backend/internal/dto"
	"cobagolang/backend/internal/utils"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/limiter"
	"github.com/gofiber/fiber/v2/middleware/recover"
	"github.com/sirupsen/logrus"
)

// RequestLogger returns a Fiber middleware for Logrus structured logging
func RequestLogger() fiber.Handler {
	log := utils.GetLogger()

	return func(c *fiber.Ctx) error {
		start := time.Now()

		// Process request
		err := c.Next()

		latency := time.Since(start)
		statusCode := c.Response().StatusCode()

		if err != nil {
			if e, ok := err.(*fiber.Error); ok {
				statusCode = e.Code
			}
		}

		entry := log.WithFields(logrus.Fields{
			"status":     statusCode,
			"method":     c.Method(),
			"path":       c.Path(),
			"ip":         c.IP(),
			"latency_ms": latency.Milliseconds(),
			"user_agent": c.Get("User-Agent"),
		})

		if statusCode >= 500 {
			if err != nil {
				entry.WithError(err).Error("HTTP Server Error")
			} else {
				entry.Error("HTTP Server Error")
			}
		} else if statusCode >= 400 {
			if err != nil {
				entry.WithError(err).Warn("HTTP Client Warning")
			} else {
				entry.Warn("HTTP Client Warning")
			}
		} else {
			entry.Info("HTTP Request")
		}

		return err
	}
}

// CORS returns CORS middleware configured for frontend communication
func CORS() fiber.Handler {
	return cors.New(cors.Config{
		AllowOrigins: "*",
		AllowHeaders: "Origin, Content-Type, Accept, Authorization",
		AllowMethods: "GET, POST, PUT, DELETE, OPTIONS",
	})
}

// Recover returns standard recover middleware formatted with APIErrorResponse
func Recover() fiber.Handler {
	return recover.New(recover.Config{
		EnableStackTrace: true,
	})
}

// RateLimiter creates basic rate limiting middleware
func RateLimiter(cfg *config.Config) fiber.Handler {
	return limiter.New(limiter.Config{
		Max:        cfg.RateLimitMax,
		Expiration: time.Duration(cfg.RateLimitDurationSeconds) * time.Second,
		LimitReached: func(c *fiber.Ctx) error {
			return c.Status(fiber.StatusTooManyRequests).JSON(
				dto.NewErrorResponse("Too many requests, please slow down.", "Rate limit exceeded"),
			)
		},
	})
}
