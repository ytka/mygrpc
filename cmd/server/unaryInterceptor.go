package main

import (
	"context"
	"log"

	hellopb "mygrpc/pkg/grpc"

	"google.golang.org/grpc"
)

func myUnaryServerInterceptor1(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (interface{}, error) {
	log.Println("[pre] my unary server interceptor 1: ", info.FullMethod) // ハンドラの前に割り込ませる前処理
	res, err := handler(ctx, req)                                         // 本来の処理
	m := ""
	if helloRes, ok := res.(*hellopb.HelloResponse); ok {
		m = helloRes.GetMessage()
	}
	log.Println("[post] my unary server interceptor 1: ", m) // ハンドラの後に割り込ませる後処理
	return res, err
}
