package gateway

import (
	"context"
	pb "grpc-server-go/proto/gen/go/grpc/hello"
	"strconv"

	gwruntime "github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"google.golang.org/grpc"
)

const (
	IP             = "0.0.0.0"
	Port           = 10000
	GrpcServerAddr = ":8090"
)

var (
	Addr string = IP + ":" + strconv.Itoa(Port)

	Apis = []func(context.Context, *gwruntime.ServeMux, *grpc.ClientConn) error{
		pb.RegisterHelloServiceHandler,
	}
)
