package grpc

import (
	"context"

	"github.com/golang/protobuf/ptypes/empty"

	productpb "github.com/ostapenkoden/microservices-demo/proto/gen/go/demo/product"
	"github.com/ostapenkoden/microservices-demo/service/product"
)

type ProductService struct {
	Repo product.Repository
}

func NewProductService(repo product.Repository) *ProductService {
	return &ProductService{Repo: repo}
}

func (s *ProductService) Find(ctx context.Context, _ *empty.Empty) (*productpb.Products, error) {
	products, err := s.Repo.Products(ctx)
	if err != nil {
		return nil, err
	}

	pbProducts := make([]*productpb.Product, len(products))

	for i, p := range products {
		pbProducts[i] = &productpb.Product{
			ID:           p.ID.String(),
			Name:         p.Name,
			Price:        p.Price,
			AvailableQty: p.AvailableQty,
		}
	}

	return &productpb.Products{Products: pbProducts}, nil
}
