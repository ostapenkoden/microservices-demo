package product

import (
	"context"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Product struct {
	ID           primitive.ObjectID `bson:"_id"`
	Name         string             `bson:"name"`
	Price        float64            `bson:"price"`
	AvailableQty int32              `bson:"available_qty"`
}

type Repository interface {
	Products(ctx context.Context) ([]Product, error)
}
