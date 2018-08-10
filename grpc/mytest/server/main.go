package main

import (
	"fmt"
	"net"

        "golang.org/x/net/context"
        "google.golang.org/grpc"
        pb "google.golang.org/grpc/examples/mytest/mytest"
)

var (
	addr = "127.0.0.1:12345"
)

//实现在proto中定义的服务，即实现.pb.go中所定义的server接口
type myserver struct {
}

//所定义的方法要与server的interface中声明的一致
func (s *myserver) GetNum(ctx context.Context, seed *pb.Seed) (*pb.Reply, error) {
	res := fmt.Sprintf("seed: %v", seed.Seed)
	return &pb.Reply{Reply : res}, nil
}

func main() {
	listen, err := net.Listen("tcp", addr)//监听
	if err != nil {
		fmt.Println(err)
		return
	}
	s := grpc.NewServer()//创建gRPC server

	pb.RegisterFirstServer(s, &myserver{})//将所实现的服务注册到服务器上
	
	err = s.Serve(listen)//在指定的IP和端口上处理请求
	if err != nil {
		fmt.Println(err)
	}
	return 
}
