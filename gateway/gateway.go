package gateway

import (
	"context"
	"fmt"
	"net/http"
	"time"

	"github.com/grpc-ecosystem/grpc-gateway/runtime"
	"google.golang.org/grpc"

	pb "github.com/Zgxh/grpc-gen/proto/gen/go/grpc"
)

func prepareGateway(ctx context.Context) (http.Handler, error) {
	// grpc dial options
	opts := []grpc.DialOption{
		grpc.WithTimeout(5 * time.Second),
		grpc.WithBlock(),
		grpc.WithInsecure(),
	}

	conn, err := grpc.Dial(Addr, opts...)
	if err != nil {
		return nil, fmt.Errorf("Failed to dail addr: %s when create the gateway, err: %v", Addr, err)
	}

	// json serializer including empty field with default values
	gwMux := runtime.NewServeMux(
		runtime.WithMarshalerOption(runtime.MIMEWildcard, &runtime.JSONPb{OrigName: true, EmitDefaults: true}),
		runtime.WithProtoErrorHandler(runtime.DefaultHTTPProtoErrorHandler),
	)

	// register gateway endpoints
	err = pb.RegisterHelloServiceHandlerServer(ctx, gwMux, conn)
	if err != nil {
		return nil, err
	}

	return gwMux, nil
}
