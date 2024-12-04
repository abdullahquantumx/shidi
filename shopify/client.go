package shopify

import (
	"github.com/Shridhar2104/logilo/shopify/pb"
	"google.golang.org/grpc"
)

type Client struct {
	conn      *grpc.ClientConn
	service   pb.ShopifyServiceClient
}

// NewClient initializes the Shopify app and GRPC connection
func NewClient(grpcURL string) (*Client, error) {
	conn, err := grpc.Dial(grpcURL, grpc.WithInsecure())
	if err != nil {
		return nil, err
	}




	return &Client{
		conn:      conn,
		service:   pb.NewShopifyServiceClient(conn),
	}, nil
}


// Close cleans up the GRPC connection
func (c *Client) Close() error {
	return c.conn.Close()
}
