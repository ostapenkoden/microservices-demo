package main

import (
	"context"
	"log"
	"net/http"

	"github.com/grpc-ecosystem/grpc-gateway/runtime"
	"google.golang.org/grpc"

	"github.com/ostapenkoden/microservices-demo/proto/gen/go/demo/product"
)

func run() error {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	mux := runtime.NewServeMux()
	opts := []grpc.DialOption{grpc.WithInsecure()}
	if err := product.RegisterProductsServiceHandlerFromEndpoint(ctx, mux, "localhost:3000", opts); err != nil {
		return err
	}

	return http.ListenAndServe(":8080", mux)
}

func main() {
	if err := run(); err != nil {
		log.Fatal(err)
	}
}
