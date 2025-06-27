package main

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/Mafilala/ketero/backend/initializers"
	"github.com/Mafilala/ketero/backend/routes"
	"github.com/gin-gonic/gin"
)

func init() {
	initializers.LoadEnvVariables()
	initializers.InitDB()
}

func main() {
	// Configure Gin for production
	if os.Getenv("GO_ENV") == "production" {
		gin.SetMode(gin.ReleaseMode) // Disable debug mode
	}

	r := gin.Default()

	// Register routes
	routes.RegisterMeasure(r)
	routes.RegisterClothingType(r)
	routes.RegisterClothing(r)
	routes.RegisterClient(r)
	routes.RegisterStatus(r)
	routes.RegisterAddClothing(r)
	routes.RegisterClothingMeasures(r)
	routes.RegisterOrderRoutes(r)
	routes.RegisterOrderMeasureRoutes(r)
	routes.RegisterPriceDetailRoutes(r)
	routes.RegisterOrderDetailRoutes(r)
	routes.RegisterUser(r)
	// Get port from environment with fallback
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080" // Default for local development
	}

	// Configure HTTP server with graceful shutdown
	server := &http.Server{
		Addr:    ":" + port,
		Handler: r,
	}

	// Graceful shutdown routine
	go func() {
		if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("Server failed: %v", err)
		}
	}()

	// Wait for interrupt signal
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit
	log.Println("Shutting down server...")

	// Create context with timeout
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	// Shutdown server
	if err := server.Shutdown(ctx); err != nil {
		log.Fatal("Server forced to shutdown:", err)
	}

	// Close database pool
	if initializers.Pool != nil {
		initializers.Pool.Close()
	}

	log.Println("Server exited")
}
