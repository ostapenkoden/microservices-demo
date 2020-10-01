package mongo

import (
	"context"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

type Client struct {
	*mongo.Client
}

func NewClient(ctx context.Context, uri string) (*Client, error) {
	c, err := mongo.Connect(ctx, options.Client().ApplyURI(uri))
	if err != nil {
		return nil, err
	}
	if err = c.Ping(ctx, readpref.Primary()); err != nil {
		return nil, err
	}
	return &Client{c}, nil
}

func (c *Client) Disconnect(ctx context.Context) error {
	return c.Client.Disconnect(ctx)
}
