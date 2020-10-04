package mongo

import (
	"context"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"

	"github.com/ostapenkoden/microservices-demo/service/product"
)

type ProductRepo struct {
	collection *mongo.Collection
}

func NewProductRepo(client *Client) *ProductRepo {
	return &ProductRepo{
		collection: client.Database("products").Collection("products"),
	}
}

func (r ProductRepo) Products(ctx context.Context) ([]product.Product, error) {
	var products []product.Product
	cursor, err := r.collection.Find(ctx, bson.M{})
	if err != nil {
		return nil, err
	}
	return products, cursor.All(ctx, &products)
}
