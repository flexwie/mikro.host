package models

import (
	"database/sql/driver"
	"gorm.io/gorm"
)

type DeploymentStatus string

const (
	Pending DeploymentStatus = "pending"
	Success DeploymentStatus = "success"
	Failure DeploymentStatus = "failure"
)

func (e *DeploymentStatus) Scan(value interface{}) error {
	*e = DeploymentStatus(value.([]byte))
	return nil
}

func (e DeploymentStatus) Value() (driver.Value, error) {
	return string(e), nil
}

type Server struct {
	gorm.Model
	Name       string
	ID         string
	PrivateKey string
	IP         string
	State      DeploymentStatus `json:"deployment_status" sql:"type:deployment_status"`
	ClusterID  uint
}

type CreateServerRequest struct {
	Cluster string `json:"cluster_id"`
}

type CreateServerResponse struct {
	Err   `json:"error,omitempty"`
	Value string `json:"value,omitempty"`
}
