package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type Product struct {
    Id       	primitive.ObjectID `json:"id"`
    Name     	string             `json:"name" validate:"required"`
    Description string             `json:"description" validate:"required"`
    Price    	float64            `json:"price" validate:"required"`
}