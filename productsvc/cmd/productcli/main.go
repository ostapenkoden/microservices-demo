package main

import (
	"context"
	"fmt"
	"os"

	"google.golang.org/grpc"

	pb "github.com/ostapenkoden/microservices-demo/pb/products"
	"github.com/ostapenkoden/microservices-demo/productsvc/envconf"
)

func main() {
	conf := envconf.New()

	conn, err := grpc.Dial(conf.RPC.Address, grpc.WithInsecure())
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
	defer conn.Close()

	client := pb.NewProductsServiceClient(conn)

	resp, err := client.Find(context.Background(), &pb.Empty{})
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}

	fmt.Printf("%+v", resp.Products)
}
