package service

import (
	"context"
	"gorm.io/gorm"
	"mikro.host/models"
)

type Service interface {
	CreateCluster(ctx context.Context, c models.CreateClusterRequest) (string, error)
}

type service struct {
	db *gorm.DB
}

func NewService(db *gorm.DB) Service {
	return &service{
		db: db,
	}
}

func (s *service) CreateCluster(ctx context.Context, c models.CreateClusterRequest) (string, error) {
	return "", nil
}
