package apis

import (
	"context"

	"google.golang.org/grpc"

	pb "github.com/Zgxh/grpc-gen/proto/gen/go/grpc/hello"
	gwruntime "github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
)

func GetApis() []func(context.Context, *gwruntime.ServeMux, *grpc.ClientConn) error {
	return []func(context.Context, *gwruntime.ServeMux, *grpc.ClientConn) error{
		pb.RegisterHelloServiceHandler,
	}
}
