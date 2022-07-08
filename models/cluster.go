package models

import "gorm.io/gorm"

type Cluster struct {
	gorm.Model
	Name   string
	Server []Server
	UserID uint
}

type CreateClusterRequest struct {
	ID string `json:"owner_id"`
}

type CreateClusterResponse struct {
	Err   `json:"error,omitempty"`
	Value string `json:"value,omitempty"`
}
