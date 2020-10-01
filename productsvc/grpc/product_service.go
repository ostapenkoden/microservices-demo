package grpc

import (
	"context"

	pb "github.com/ostapenkoden/microservices-demo/pb/products"
	"github.com/ostapenkoden/microservices-demo/productsvc/product"
)

type ProductService struct {
	Repo product.Repository
}

func NewProductService(repo product.Repository) *ProductService {
	return &ProductService{Repo: repo}
}

func (s *ProductService) Find(ctx context.Context, empty *pb.Empty) (*pb.Products, error) {
	products, err := s.Repo.Products(ctx)
	if err != nil {
		return nil, err
	}

	pbProducts := make([]*pb.Product, len(products))

	for i, p := range products {
		pbProducts[i] = &pb.Product{
			ID:           p.ID.String(),
			Name:         p.Name,
			Price:        p.Price,
			AvailableQty: p.AvailableQty,
		}
	}

	return &pb.Products{Products: pbProducts}, nil
}

