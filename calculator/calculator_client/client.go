package main

import (
	"context"
	"fmt"
	"io"
	"log"
	"machinetest/test/calculator/calculatorpb"

	"google.golang.org/grpc"
)

func main() {

	fmt.Println("iam a client")
	cc, err := grpc.Dial("localhost:50051", grpc.WithInsecure()) //creating clent connection

	if err != nil {
		log.Fatalf("could not connect : %v", err)
	}
	defer cc.Close()
	c := calculatorpb.NewCalculatorServiceClient(cc)

//	doUnary(c)
	doServerStreaming(c)

}
func doUnary(c calculatorpb.CalculatorServiceClient) {
	fmt.Println("starting to unary function RPC ...")
	req := &calculatorpb.SumRequest{
		FirstNumber:  23,
		SecondNumber: 3453453,
	}
	res, err := c.Sum(context.Background(), req)
	if err != nil {
		log.Fatalf("error while calling sum RPC %v", err)
	}

	log.Printf("Response from sum: %v", res)

}

func doServerStreaming(c calculatorpb.CalculatorServiceClient){

	fmt.Println("starting to do server streaming msg in RPC....")
	req:=&calculatorpb.PrimeNumberDecompositionRequest{
		Number: 1234534,
	}
streamResult,err:=c.PrimeNumberDecomposition(context.Background(),req)

if err!=nil{
	log.Fatalf("error while calling prime number decomposition %v",err)

}
for {

msg,err:=streamResult.Recv()

	if err==io.EOF{
		break
	}
if err!=nil{
	log.Fatalf("error while reading stream %v",err)
}

log.Printf("Response from Greetmany times %v",msg.GetPrimFactor())




}


}
