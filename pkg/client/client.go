package client

import (
	"context"
	"crypto/tls"
	"crypto/x509"
	"fmt"
	pb "gogrpc/pkg/proto"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
)

// Client is a gRPC client wrapper.
type Client struct {
	conn       *grpc.ClientConn
	grpcClient pb.DemoServiceClient
}

// New creates a new gRPC client.
func New(address string, opts ...grpc.DialOption) (*Client, error) {
	conn, err := grpc.NewClient(address, opts...)
	if err != nil {
		return nil, fmt.Errorf("failed to dial: %v", err)
	}

	grpcClient := pb.NewDemoServiceClient(conn)

	return &Client{
		grpcClient: grpcClient,
		conn:       conn,
	}, nil
}

// NewTLS creates a new gRPC client with TLS.
func NewTLS(address string, certPool *x509.CertPool) (*Client, error) {
	cred := credentials.NewTLS(&tls.Config{ServerName: "", RootCAs: certPool})

	return New(address, grpc.WithTransportCredentials(cred))
}

// Unary
// C --> S
// C <-- S
// Returns Id*2.
func (c *Client) OneReqOneResp(ctx context.Context, req *pb.DemoRequest) (*pb.DemoResponse, error) {
	return c.grpcClient.OneReqOneResp(ctx, req)
}

// Server streaming
// C --> S (once at start)
// C < - S (stream)
// Returns i={1,3} => Id*i.
func (c *Client) OneReqManyResp(ctx context.Context, req *pb.DemoRequest) (pb.DemoService_OneReqManyRespClient, error) {
	return c.grpcClient.OneReqManyResp(ctx, req)
}

// Client streaming
// C - > S (stream)
// C <-- S (once at end)
// Returns sum of sended id's.
func (c *Client) ManyReqOneResp(ctx context.Context) (pb.DemoService_ManyReqOneRespClient, error) {
	return c.grpcClient.ManyReqOneResp(ctx)
}

// Bi-directional stream
// C - > S (stream)
// C < - S (stream)
// Sends Id*2 for each message.
func (c *Client) ManyReqManyResp(ctx context.Context) (pb.DemoService_ManyReqManyRespClient, error) {
	return c.grpcClient.ManyReqManyResp(ctx)
}

// Always return error `codes.InvalidArgument`.
func (c *Client) UnaryFail(ctx context.Context, req *pb.DemoRequest) (*pb.DemoResponse, error) {
	return c.grpcClient.UnaryFail(ctx, req)
}

// Always answers after 5 seconds.
func (c *Client) UnaryDeadline(ctx context.Context, req *pb.DemoRequest) (*pb.DemoResponse, error) {
	return c.grpcClient.UnaryDeadline(ctx, req)
}

// Close closes the gRPC client.
func (c *Client) Close() {
	if c.conn != nil {
		c.conn.Close()
	}
}
