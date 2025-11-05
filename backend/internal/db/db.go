package db

import (
	"log"

	"github.com/vaidik-bajpai/Nexus/backend/internal/db/db"
)

func NewDB() *db.PrismaClient {
	client := db.NewClient()
	if err := client.Connect(); err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}
	return client
}
