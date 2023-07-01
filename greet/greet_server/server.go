package main

import (
	"fmt"
	"io"
	"log"
	"machinetest/test/greet/greetpb"
	"strconv"
	"time"

	"net"

	"context"

	"google.golang.org/grpc"
)

type server struct {
}

func (*server) Greet(ctx context.Context, req *greetpb.GreetRequest) (*greetpb.GreetResponse, error) {

	fmt.Println("greet functoin was invoked",req)
	firstName := req.GetGreeting().FirstName

	result := "hellow" + firstName

	res := greetpb.GreetResponse{
		Result: result,
	}
	return &res, nil
}
func (*server)GreetManyTimes(req *greetpb.GreetManytimesRequest,stream greetpb.GreetService_GreetManyTimesServer)( error){

	fmt.Println("Greet many times function invoked with",req)
	fname:=req.GetGreeting().GetFirstName()

	for i:=0;i<10;i++{
		resutlt:="hello "+fname+" Number" +strconv.Itoa(i)
		res:=&greetpb.GreetManytimesResponse{
			Result:resutlt ,
		}
		stream.Send(res)
		time.Sleep(time.Second)
	}
	return nil

}


func (*server)LongGreet(stream greetpb.GreetService_LongGreetServer) error{

	fmt.Println("long greet function is invoked with  streaming request")
	result:=""
for {
	req,err:=stream.Recv()
	if err==io.EOF{
		//we have finished the reading the client stream
		return stream.SendAndClose(&greetpb.LongGreetResponse{
			Result: result,
		})
	}

	if err!=nil{
		log.Fatalf("Error while reading client stream log greet function %v",err)
	}
	firstName:=req.GetGreeting().GetFirstName()
	result+="Hello "+ firstName+"! "
	
}
 
}

func main() {
	fmt.Println("hello word")
	lis, err := net.Listen("tcp", "0.0.0.0:50051") 
	if err != nil {
		log.Fatalf("failed to listern %v", err)  
	}
	s := grpc.NewServer()
	greetpb.RegisterGreetServiceServer(s, &server{})
	if err := s.Serve(lis); err != nil {
		log.Fatalf("faile to serve :%v", err)
	} 
}
