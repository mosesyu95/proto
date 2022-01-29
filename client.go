package main

import (
	"context"
	"fmt"
	"github.com/mosesyu95/proto/pb"
	"google.golang.org/grpc"
	"log"
)

func main() {
	conn, err := grpc.Dial("127.0.0.1:8080",grpc.WithInsecure())
	if err != nil {
		log.Fatalln("dailing error: ", err)
	}
	defer conn.Close()
	arithClient := pb.NewArithClient(conn)
	req := &pb.Args{
		A: 2,
		B: 4,
	}
	res, err := arithClient.Calc(context.Background(),req )
	fmt.Printf("%d + %d = %d \n", req.A, req.B, res.Sum)

}
