package main

import (
	"fmt"
	"time"
	"math/rand"

	"golang.org/x/net/context"
	"google.golang.org/grpc"
	pb "google.golang.org/grpc/examples/mytest/mytest"
)

var (
	addr = "127.0.0.1:12345"
)

func main() {
	conn, err := grpc.Dial(addr, grpc.WithInsecure())//创建到gRPC服器的连接
	if err != nil {
		fmt.Println(err)
		return
	}
	defer conn.Close()
	
	c := pb.NewFirstClient(conn)//创建gRPC客户端
	ctx, cancle := context.WithTimeout(context.Background(), time.Second*5)
	defer cancle()
	rand.Seed(time.Now().Unix())
	seed := pb.Seed{Seed : rand.Int31()}
	res, err := c.GetNum(ctx, &seed)//调用gRPC,同步阻塞的：
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(res.Reply)
}
