package main

import (
	"context"
	"fmt"
	"grpc_course/greetpb"
	"log"

	"google.golang.org/grpc"
)

func main() {
	fmt.Println("Hello Im a client!")

	cc, err := grpc.Dial("localhost:50051", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("couldnt connect: %v", err)
	}
	defer cc.Close()
	c := greetpb.NewGreetsServiceClient(cc)
	doUnary(c)
}



func doUnary(c greetpb.GreetsServiceClient) {
	fmt.Println(("starting to do a unary RPC..."))
	req := &greetpb.GreetRequest{
		Greeting: &greetpb.Greeting{
			FirstName: "Teju",
			LastName:  "Bondila",
		},
	}
	res, err := c.Greet(context.Background(), req)
	if err != nil {
		log.Fatalf("error while calling greet RPC:% v", err)
	}
	log.Printf("Response from greet: %v", res.Result)

}
