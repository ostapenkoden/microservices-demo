package grpc

import (
	"context"

	"github.com/golang/protobuf/ptypes/empty"

	"github.com/ostapenkoden/microservices-demo/inventory/product"
	"github.com/ostapenkoden/microservices-demo/proto/demo/inventory"
)

type ProductService struct {
	Repo product.Repository
}

func NewProductService(repo product.Repository) *ProductService {
	return &ProductService{Repo: repo}
}

func (s *ProductService) Find(ctx context.Context, _ *empty.Empty) (*inventory.Products, error) {
	products, err := s.Repo.Products(ctx)
	if err != nil {
		return nil, err
	}

	pbProducts := make([]*inventory.Product, len(products))

	for i, p := range products {
		pbProducts[i] = &inventory.Product{
			ID:           p.ID.String(),
			Name:         p.Name,
			Price:        p.Price,
			AvailableQty: p.AvailableQty,
		}
	}

	return &inventory.Products{Products: pbProducts}, nil
}
