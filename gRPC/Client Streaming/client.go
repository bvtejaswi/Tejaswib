package main

import (
	"context"
	"fmt"
	"grpc_course/greetpb"
	"io"
	"log"
	"time"

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
	//doUnary(c)

	//doServerStreaming(c)
	doClientStreaming(c)
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

func doServerStreaming(c greetpb.GreetsServiceClient) {
	fmt.Println("Starting to do a server streaming RPC...")

	req := &greetpb.GreetManyTimesRequest{
		Greeting: &greetpb.Greeting{
			FirstName: "Teju",
			LastName:  "Bondila",
		},
	}
	resStream, err := c.GreetManyTimes(context.Background(), req)
	if err != nil {
		log.Fatalln("Error while calling GreetManyTimes RPC: %v", err)
	}
	for {
		msg, err := resStream.Recv()
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatalf("error while reading stream : %v", err)
		}
		log.Printf("Response from GreetManyTimes: %v", msg.GetResult())

	}
}

func doClientStreaming(c greetpb.GreetsServiceClient) {
	fmt.Println("Starting to do a client streaming RPC...")
	requests := []*greetpb.LongGreetRequest{
		&greetpb.LongGreetRequest{
			Greeting: &greetpb.Greeting{
				FirstName: "Teju",
			},
		},
		&greetpb.LongGreetRequest{
			Greeting: &greetpb.Greeting{
				FirstName: "Cherry",
			},
		},
		&greetpb.LongGreetRequest{
			Greeting: &greetpb.Greeting{
				FirstName: "Betty",
			},
		},
		&greetpb.LongGreetRequest{
			Greeting: &greetpb.Greeting{
				FirstName: "Arch",
			},
		},
		&greetpb.LongGreetRequest{
			Greeting: &greetpb.Greeting{
				FirstName: "Damon",
			},
		},
	}

	stream, err := c.LongGreet(context.Background())
	if err != nil {
		log.Fatalf("error while calling LongGreet: %v", err)
	}

	for _, req := range requests {
		fmt.Printf("Sendind requesrs :%v\n", req)
		stream.Send(req)
		time.Sleep(1000 * time.Millisecond)
	}
	res, err := stream.CloseAndRecv()
	if err != nil {
		log.Fatalf("error while receiving response from LongGreet:%v", err)
	}
	fmt.Printf("LongGreet Response: %v\n", res)

}
