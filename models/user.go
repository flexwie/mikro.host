package models

import "gorm.io/gorm"

type User struct {
	gorm.Model
	Name    string
	Mail    string
	Cluster []Cluster
}

type CreateRequest struct {
	Name string `json:"name"`
	Mail string `json:"mail"`
}

type CreateResponse struct {
	Err   `json:"error,omitempty"`
	Value User `json:"value,omitempty"`
}

type GetAllResponse struct {
	Err   `json:"error,omitempty"`
	Value []User `json:"value,omitempty"`
}

type GetOneRequest struct {
	Id uint `json:"id"`
}

type GetOneResponse struct {
	Err   `json:"error,omitempty"`
	Value User `json:"value,omitempty"`
}

type UpdateRequest struct {
	User
}

type UpdateResponse struct {
	Err   `json:"error,omitempty"`
	Value User `json:"value,omitempty"`
}

type DeleteRequest struct {
	Id uint `json:"id"`
}

type DeleteResponse struct {
	Err `json:"error,omitempty"`
}
