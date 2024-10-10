package server

import (
	"context"
	pb "gogrpc/pkg/proto"
	"io"
	"log"
	"time"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

// Unary
// C --> S
// C <-- S
// Returns Id*2.
func (s *Server) OneReqOneResp(_ context.Context, in *pb.DemoRequest) (*pb.DemoResponse, error) {
	log.Printf("OneReqOneResp invoked with %v\n", in)

	return &pb.DemoResponse{
		Result: in.Id * 2,
	}, nil
}

// Server streaming
// C --> S (once at start)
// C < - S (stream)
// Returns i={1,3} => Id*i.
func (s *Server) OneReqManyResp(in *pb.DemoRequest, stream pb.DemoService_OneReqManyRespServer) error {
	log.Printf("OneReqManyResp invoked with %v\n", in)

	for i := uint32(1); i <= 3; i++ {
		err := stream.Send(&pb.DemoResponse{
			Result: in.Id * i,
		})
		if err != nil {
			return err
		}
	}

	return nil
}

// Client streaming
// C - > S (stream)
// C <-- S (once at end)
// Returns sum of sended id's.
func (s *Server) ManyReqOneResp(stream pb.DemoService_ManyReqOneRespServer) error {
	log.Printf("ManyReqOneResp invoked\n")

	sum := uint32(0)
	for {
		msg, err := stream.Recv()
		if err == io.EOF {
			break
		} else if err != nil {
			return err
		}

		sum += msg.Id
	}

	return stream.SendAndClose(&pb.DemoResponse{
		Result: sum,
	})
}

// Bi-directional stream
// C - > S (stream)
// C < - S (stream)
// Sends Id*2 for each message.
func (s *Server) ManyReqManyResp(stream pb.DemoService_ManyReqManyRespServer) error {
	log.Printf("ManyReqManyResp invoked\n")

	for {
		msg, err := stream.Recv()
		if err == io.EOF {
			break
		} else if err != nil {
			return err
		}

		if err = stream.Send(&pb.DemoResponse{Result: msg.Id * 2}); err != nil {
			return err
		}
	}

	return nil
}

// Custom error.
var RpcErrCustom = status.New(
	codes.InvalidArgument,
	"custom error message",
)

// Always return error `codes.InvalidArgument`.
func (s *Server) UnaryFail(_ context.Context, in *pb.DemoRequest) (*pb.DemoResponse, error) {
	return nil, RpcErrCustom.Err()
}

// Always answers after 5 seconds.
func (s *Server) UnaryDeadline(_ context.Context, in *pb.DemoRequest) (*pb.DemoResponse, error) {
	time.Sleep(time.Second * 5)

	return &pb.DemoResponse{Result: 1}, nil
}
