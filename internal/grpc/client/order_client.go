package client

import (
	"context"
	"food-delivery-backend/pkg/config"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

type OrderServiceClient struct {
	conn *grpc.ClientConn
}

func NewOrderServiceClient(cfg *config.Config) (*OrderServiceClient, error) {
	conn, err := grpc.Dial(cfg.GRPC.OrderAddr, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		return nil, err
	}
	return &OrderServiceClient{conn: conn}, nil
}

func (c *OrderServiceClient) Close() error { return c.conn.Close() }
func (c *OrderServiceClient) PlaceOrder(ctx context.Context, in any) (any, error) {
	_ = ctx
	return in, nil
}
func (c *OrderServiceClient) GetOrder(ctx context.Context, in any) (any, error) {
	_ = ctx
	return in, nil
}
func (c *OrderServiceClient) CancelOrder(ctx context.Context, in any) (any, error) {
	_ = ctx
	return in, nil
}
func (c *OrderServiceClient) UpdateOrderStatus(ctx context.Context, in any) (any, error) {
	_ = ctx
	return in, nil
}
func (c *OrderServiceClient) GetOrderTracking(ctx context.Context, in any) (any, error) {
	_ = ctx
	return in, nil
}
func (c *OrderServiceClient) GetUserOrders(ctx context.Context, in any) (any, error) {
	_ = ctx
	return in, nil
}
