package db

import (
	"krriya/internal/adapter/ports"

	"go.mongodb.org/mongo-driver/mongo"
)

type RepoImpl struct {
	db *mongo.Database
}

func NewAdminRepo(db *mongo.Database) ports.Repo {
	return &RepoImpl{
		db: db,
	}
}
