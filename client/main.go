package main

import (
	"context"
	"fmt"
	"log"

	pb "example.com/cat/proto"

	"google.golang.org/grpc"
)

func main() {
	//sampleなのでwithInsecure
	conn, err := grpc.Dial("127.0.0.1:19003", grpc.WithInsecure())
	if err != nil {
		log.Fatal("client connection error:", err)
	}
	defer conn.Close()
	client := pb.NewCatClient(conn)
	message := &pb.GetMyCatMessage{TargetCat: "tama"}
	res, err := client.GetMyCat(context.TODO(), message)
	fmt.Printf("result:%#v \n", res)
	fmt.Println("---------")
	fmt.Printf("result:%#v \n", res.Name)
	fmt.Printf("error::%#v \n", err)
}
