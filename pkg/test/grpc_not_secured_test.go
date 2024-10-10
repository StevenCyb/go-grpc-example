package tests

import (
	"context"
	"fmt"
	"gogrpc/pkg/client"
	"gogrpc/pkg/proto"
	"gogrpc/pkg/server"
	"io"
	"sync"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/grpc/status"
)

func TestGRPC_NotSecured(t *testing.T) {
	t.Parallel()

	address := fmt.Sprintf("localhost:%d", GetFreeLocalPort(t))
	createServer(t, address)
	grpcClient := createClient(t, address)

	t.Run("Unary", func(t *testing.T) {
		t.Parallel()

		req := proto.DemoRequest{Id: 1}
		resp, err := grpcClient.OneReqOneResp(context.Background(), &req)
		assert.NoError(t, err)
		assert.Equal(t, req.Id*2, resp.Result)
	})

	t.Run("OneReqManyResp", func(t *testing.T) {
		t.Parallel()

		req := proto.DemoRequest{Id: 1}
		stream, err := grpcClient.OneReqManyResp(context.Background(), &req)
		assert.NoError(t, err)

		index := uint32(1)
		for {
			msg, err := stream.Recv()
			if err == io.EOF {
				break
			} else if err != nil {
				assert.NoError(t, err)
			}

			assert.Equal(t, index*req.Id, msg.Result)
			index++
		}
	})

	t.Run("ManyReqOneResp", func(t *testing.T) {
		t.Parallel()

		stream, err := grpcClient.ManyReqOneResp(context.Background())
		assert.NoError(t, err)

		sum := uint32(0)
		for i := uint32(1); i < 3; i++ {
			sum += i
			assert.NoError(t, stream.Send(&proto.DemoRequest{Id: i}))
		}

		resp, err := stream.CloseAndRecv()
		assert.NoError(t, err)
		assert.Equal(t, sum, resp.Result)
	})

	t.Run("ManyReqManyResp", func(t *testing.T) {
		t.Parallel()

		stream, err := grpcClient.ManyReqManyResp(context.Background())
		assert.NoError(t, err)

		reqs := []*proto.DemoRequest{
			{Id: 1},
			{Id: 2},
			{Id: 3},
		}
		resps := []*proto.DemoResponse{}

		wg := &sync.WaitGroup{}
		wg.Add(2)

		go func() {
			defer wg.Done()

			for _, req := range reqs {
				assert.NoError(t, stream.Send(req))
				time.Sleep(time.Microsecond * 100)
			}

			assert.NoError(t, stream.CloseSend())
		}()

		go func() {
			defer wg.Done()

			for {
				resp, err := stream.Recv()
				if err == io.EOF {
					break
				} else if err != nil {
					assert.NoError(t, err)
				}

				resps = append(resps, resp)
			}
		}()

		wg.Wait()
		assert.Equal(t, len(reqs), len(resps))
		for i := 0; i < len(resps); i++ {
			assert.Equal(t, reqs[i].Id*2, resps[i].Result)
		}
	})

	t.Run("Unary_CustomError", func(t *testing.T) {
		t.Parallel()

		req := proto.DemoRequest{Id: 1}
		_, err := grpcClient.UnaryFail(context.Background(), &req)
		assert.Error(t, err)
		rpcError, ok := status.FromError(err)
		assert.True(t, ok)
		assert.Equal(t, server.RpcErrCustom.Code(), rpcError.Code())
		assert.Equal(t, server.RpcErrCustom.String(), rpcError.String())
	})

	t.Run("Unary_Deadline", func(t *testing.T) {
		t.Parallel()

		ctx, cancel := context.WithTimeout(context.Background(), time.Second)
		defer cancel()

		req := proto.DemoRequest{Id: 1}
		_, err := grpcClient.UnaryDeadline(ctx, &req)
		assert.Error(t, err)
		rpcError, ok := status.FromError(err)
		assert.True(t, ok)
		assert.Equal(t, codes.DeadlineExceeded, rpcError.Code())
	})
}

func createServer(t *testing.T, address string) {
	t.Helper()

	grpcServer, err := server.New(address)
	assert.NoError(t, err)
	t.Cleanup(func() {
		time.Sleep(time.Second)
		grpcServer.Close()
	})
	go func() {
		assert.NoError(t, grpcServer.Start())
	}()
}

func createClient(t *testing.T, address string) *client.Client {
	t.Helper()

	grpcClient, err := client.New(
		address,
		grpc.WithTransportCredentials(insecure.NewCredentials()), // if no security
	)
	assert.NoError(t, err)
	t.Cleanup(func() {
		time.Sleep(time.Second)
		grpcClient.Close()
	})

	return grpcClient
}
