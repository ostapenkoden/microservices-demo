package grpc

import (
	"net"

	"google.golang.org/grpc"

	pb "github.com/ostapenkoden/microservices-demo/pb/products"
)

type Server struct {
	*grpc.Server
}

func NewServer(srv *ProductService) *Server {
	s := grpc.NewServer()
	pb.RegisterProductsServiceServer(s, srv)
	return &Server{s}
}

func (s Server) ListenAndServe(addr string) error {
	lis, err := net.Listen("tcp", addr)
	if err != nil {
		return err
	}
	return s.Serve(lis)
}