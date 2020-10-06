package main

import (
	"context"
	"fmt"
	"os"

	"github.com/ostapenkoden/microservices-demo/inventory/envconf"
	"github.com/ostapenkoden/microservices-demo/inventory/grpc"
	"github.com/ostapenkoden/microservices-demo/inventory/mongo"
)

func run() error {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	conf := envconf.New()

	mongoClient, err := mongo.NewClient(ctx, conf.MongoDB.ConnectionURI)
	if err != nil {
		return fmt.Errorf("init mongo: %v", err)
	}
	defer mongoClient.Disconnect(ctx)

	productRepo := mongo.NewProductRepo(mongoClient)
	productService := grpc.NewProductService(productRepo)
	srv := grpc.NewServer(productService)

	if err := srv.ListenAndServe(conf.RPC.Address); err != nil {
		return fmt.Errorf("listen and serve: %v", err)
	}

	return nil
}

func main() {
	if err := run(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
