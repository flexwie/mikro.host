package models

import "gorm.io/gorm"

type User struct {
	gorm.Model
	Name string
	Mail string
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
	Id string
}

type GetOneResponse struct {
	Err   `json:"error,omitempty"`
	Value User `json:"error,omitempty"`
}
