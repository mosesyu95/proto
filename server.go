package main

import (
	"context"
	"github.com/mosesyu95/proto/pb"
	"google.golang.org/grpc"
	"log"
	"net"
)

// Arith 定义算数算法服务
type Arith struct {

}


func (a *Arith) Calc(ctx context.Context,args *pb.Args) (*pb.Answer, error) {
	log.Printf("receive %+v \n", args)
	var res = &pb.Answer{Sum: args.A + args.B}
	return res ,nil
}

func main() {
	lis, err := net.Listen("tcp", ":8080")
	if err != nil {
		log.Fatalln(err)
	}
	grpcserver := grpc.NewServer()
	pb.RegisterArithServer(grpcserver, &Arith{})

	go grpcserver.Serve(lis)
	select {
	}
}
