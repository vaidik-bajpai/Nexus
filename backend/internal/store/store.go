package store

import (
	"github.com/vaidik-bajpai/Nexus/backend/internal/db/db"
)

type Store struct {
	db *db.PrismaClient
}

func NewStore(db *db.PrismaClient) *Store {
	return &Store{db: db}
}

func (s *Store) Close() error {
	return s.db.Disconnect()
}
