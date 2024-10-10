# gRPC in Go: A Comprehensive Guide
Welcome to this comprehensive guide on getting started with gRPC in GoLang. 
This guide will walk you through the basics of gRPC, how to set up a gRPC server and client, and how to implement both secured and unsecured communication channels.

## Prerequisites
Before we dive in, ensure you have the necessary tools installed:
```sh
go install google.golang.org/protobuf/cmd/protoc-gen-go@latest
go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@latest
```

## Creating the Proto File
The first step in using gRPC is to define your service and messages in a `.proto` file. 
For this demo, we have created a proto file that includes all variants of communication, including custom error and deadline cases. 
You can find the demo proto file here: [pkg/proto_spec/dummy.proto](pkg/proto_spec/dummy.proto).

### Communication Variants
There are four main types of communication in gRPC:

#### Unary
A simple request-response interaction.
```
C --> S (once at start)
C <-- S (once at end)
```

#### Server Streaming
The server sends a stream of responses to the client after a client request.
```
C --> S (once at start)
C <-- S (stream)
```

#### Client Streaming
The client sends a stream of requests to the server and the server responses at the end.
```
C --> S (stream)
C <-- S (once at end)
```

#### Bi-directional Streaming
Both the client and server send a stream of messages to each other.
```
C --> S (stream)
C <-- S (stream)
```

## Generating the Code
Once you have defined your proto file, you can generate the Go code for the client and server using the following command:
```sh
protoc --go_out=. --go-grpc_out=. -I pkg/proto_spec pkg/proto_spec/Demo.proto 
```
This command will generate the necessary Go code based on the proto file.
In this case the generated code is generated into [pkg/proto](pkg/proto).

## Creating a Client Wrapper
While it's not required, wrapping the generated gRPC client code can provide a cleaner abstraction and make your code more maintainable. 
In this demo, we wrap the `grpc.ClientConn` and `pb.DemoServiceClient` that were generated in the previous step.

You can take a look at the wrapped client implementation here: [pkg/client/client.go](pkg/client/client.go).

### Example Usage
Here is an example of how to use the wrapped client in your application:

```go
func main() {
    grpcClient, err := client.New(
        serverAddr,
        grpc.WithTransportCredentials(insecure.NewCredentials()),
    )
    AssertNoErr("failed to create client: %w", err)
    defer grpcClient.Close()

    resp, err := grpcClient.OneReqOneResp(context.Background(), &proto.DemoRequest{Id: 1})
    // ...
}
```

## Creating a Server
Creating a gRPC server is straightforward.
Below is an example of how to create and start a gRPC server.
The full demo server is implemented here [pkg/server/server.go](pkg/server/server.go) and its handler here[pkg/server/handler.go](pkg/server/handler.go).

```go
type Server struct {
    pb.DemoServiceServer
    grpcServer *grpc.Server
    listener   net.Listener
}

// New creates a new gRPC server.
func New(listen string, opt ...grpc.ServerOption) (*Server, error) {
    listener, err := net.Listen("tcp", listen)
    if err != nil {
        return nil, fmt.Errorf("failed to listen: %v", err)
    }
    time.Sleep(time.Millisecond * 100) // give the listener time to start
    grpcServer := grpc.NewServer(opt...)
    pb.RegisterDemoServiceServer(grpcServer, &Server{})
    return &Server{
        grpcServer: grpcServer,
        listener:   listener,
    }, nil
}

// Start starts the gRPC server.
func (s *Server) Start() error {
    if err := s.grpcServer.Serve(s.listener); err != nil {
        return fmt.Errorf("failed to serve gRPC: %v", err)
    }
    return nil
}

// Close closes the gRPC server.
func (s *Server) Close() {
    if s.listener != nil {
        s.listener.Close()
    }
    if s.grpcServer != nil {
        s.grpcServer.Stop()
    }
}

// Unary method.
func (s *Server) OneReqOneResp(_ context.Context, in *pb.DemoRequest) (*pb.DemoResponse, error) {
	return &pb.DemoResponse{
		Result: in.Id * 2,
	}, nil
}
```

## Running Tests
We have also included tests to ensure the functionality of both secured and unsecured gRPC communication and to show how the server and clients are used together. 
You can find the test cases here [pkg/test/grpc_not_secured_test.go](pkg/test/grpc_not_secured_test.go).
You can find the test cases for the TLS variant here [pkg/test/grpc_secured_test.go](pkg/test/grpc_secured_test.go).

## Conclusion
This guide has walked you through the basics of setting up a gRPC server and client in Go, including both secured and unsecured communication channels. 
By following these steps, you should be able to implement gRPC in your own projects with ease.

For more detailed information, refer to the official gRPC documentation: [gRPC.io](https://grpc.io/docs/).
