package main

import (
	"fmt"
	"io"
	"log"
	"machinetest/test/greet/greetpb"
	"time"

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
	//doServerStreaming(c)
	doClientStreaming(c)
 

}
func doClientStreaming(c greetpb.GreetServiceClient){
stream,err:= c.LongGreet(context.Background())

requests:=[]*greetpb.LongGreetRequest{
	&greetpb.LongGreetRequest{
		Greeting: &greetpb.Greeting{
			FirstName: "prasanth",

		},

	},
	&greetpb.LongGreetRequest{
		Greeting: &greetpb.Greeting{
			FirstName: "brother",

		},

	},
	&greetpb.LongGreetRequest{
		Greeting: &greetpb.Greeting{
			FirstName: "varun chakravarthy",

		},

	},
	&greetpb.LongGreetRequest{
		Greeting: &greetpb.Greeting{
			FirstName: "arun",

		},

	},
	&greetpb.LongGreetRequest{
		Greeting: &greetpb.Greeting{
			FirstName: "Abid",

		},

	},
	&greetpb.LongGreetRequest{
		Greeting: &greetpb.Greeting{
			FirstName: "sourav",

		},

	},

}

if err!=nil{
	log.Fatal("error while calling long greet")
}
for _,req:=range requests{
	fmt.Printf("sending req:  %v\n",req)

	stream.Send(req)
	time.Sleep(time.Second)

}
res,err:=stream.CloseAndRecv()
if err!=nil{
	log.Fatalf("Error while reciving response %v",err)
}
fmt.Printf("LongGreet response %v",res)


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
 