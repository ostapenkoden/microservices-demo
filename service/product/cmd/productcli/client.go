package main

import (
	"context"
	"fmt"
	"os"

	"google.golang.org/grpc"
	"google.golang.org/protobuf/types/known/emptypb"

	"github.com/ostapenkoden/microservices-demo/proto/gen/go/demo/product"
	"github.com/ostapenkoden/microservices-demo/service/product/envconf"
)

func main() {
	conf := envconf.New()

	conn, err := grpc.Dial(conf.RPC.Address, grpc.WithInsecure())
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
	defer conn.Close()

	client := product.NewProductsServiceClient(conn)

	resp, err := client.Find(context.Background(), &emptypb.Empty{})
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}

	fmt.Printf("%+v", resp.Products)
}
