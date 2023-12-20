package main

import (
	"context"
	"google.golang.org/grpc"
	"log"
	pb "rpc-test/proto"
	"time"
)

const (
	defaultName = "world1"
)

func main() {
	conn, err := grpc.Dial("localhost:50051", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}

	defer func(conn *grpc.ClientConn) {
		err := conn.Close()
		if err != nil {
			log.Fatalf("conn Close error : %v", err)
		}
	}(conn)

	// 初始化Greeter服务客户端
	c := pb.NewGreeterClient(conn)

	// 初始化上下文, 设置请求超时时间为1秒
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	// 调用SayHello接口发送一条消息
	res, err := c.SayHello(ctx, &pb.HelloRequest{Name: defaultName})
	if err != nil {
		log.Fatalf("could not greet: %v", err)
	}

	// 打印服务的返回的消息
	log.Printf("Greeting: %s", res.Message)
}
