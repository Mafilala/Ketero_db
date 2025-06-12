package initializers

import (
	"context"
	"log"
	"os"
	"time"

	"github.com/jackc/pgx/v5/pgxpool"
)

var Pool *pgxpool.Pool

func InitDB() {
	connStr := os.Getenv("DB_URL")
	if connStr == "" {
		log.Fatal("DB_URL environment variable not set")
	}

	config, err := pgxpool.ParseConfig(connStr)
	if err != nil {
		log.Fatalf("Failed to parse DB_URL: %v", err)
	}

	// Production-optimized connection pooling
	if os.Getenv("GO_ENV") == "production" {
		config.MaxConns = 25                // Limit maximum connections
		config.MinConns = 5                 // Maintain minimum connections
		config.MaxConnLifetime = 30 * time.Minute
		config.MaxConnIdleTime = 5 * time.Minute
	}

	// Create connection pool
	Pool, err = pgxpool.NewWithConfig(context.Background(), config)
	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}

	// Verify connection
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	
	if err := Pool.Ping(ctx); err != nil {
		log.Fatalf("Database ping failed: %v", err)
	}

	log.Println("Database connected successfully")
}
