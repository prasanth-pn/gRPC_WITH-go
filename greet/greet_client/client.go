package main

import (
	"fmt"
	"io"
	"log"
	"machinetest/test/greet/greetpb"

	"context"

	"google.golang.org/grpc"
)

func main() {

	fmt.Println("hellow iam a client")

	cc, err := grpc.Dial("localhost:50051", grpc.WithInsecure())

	if err != nil {
		log.Fatalf("could not connect :%v", err)
	}
	defer cc.Close()
	c := greetpb.NewGreetServiceClient(cc)
	//fmt.Printf("created client  %f\n", c)
//	doUnary(c)
	doServerStreaming(c)
 

}

func doServerStreaming(c greetpb.GreetServiceClient){
	fmt.Println("starting to do a server streaming rpc ....")
req:=&greetpb.GreetManytimesRequest{
	Greeting: &greetpb.Greeting{
		FirstName: "sourav",
		LastName: "kuttappan",
	},
	 
}
	streamResult,err:=c.GreetManyTimes(context.Background(),req) 

	if err!=nil{
		log.Fatalf("error while calling greettimes %v",err)
	}
	for {
	msg,err:=streamResult.Recv()

	if err==io.EOF{
		break
	}
	if err!=nil{
		log.Fatalf("error while reading stream %v",err)
	}

	log.Printf("Response form Greetmany times %v",msg.GetResult())
	}

}
func doUnary(c greetpb.GreetServiceClient){


	fmt.Println("starting to unaru function RPC ......")
	req := &greetpb.GreetRequest{
		Greeting: &greetpb.Greeting{

			FirstName: "prasand",
			LastName:  "padavayal",
		},
	}

	res,err:=c.Greet(context.Background(), req)
	if err!=nil{
		log.Fatalf("error while calling Greet rpc: %v",err)
	}

	log.Printf("response from Greet %v",res.Result)
	 
}
 