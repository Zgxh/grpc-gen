package main

import (
	"grpc-server-go/gateway"
)

// // start the grpc gateway server
// func StartGrpcGatewayServer() {
// 	conn, err := net.Listen("tcp", ":"+gateway.Port)
// 	if err != nil {
// 		panic(err)
// 	}
// 	defer conn.Close()

// 	server := gateway.NewGatewayServer(gateway.Addr, conn)

// 	server.start()
// }

func main() {
	gateway.StartGateway()
}
