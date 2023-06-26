package main

import (
	"context"
	"fmt"
	"log"
	"machinetest/test/calculator/calculatorpb"

	"net"

	"google.golang.org/grpc"
)

type server struct{}

func (*server)Sum(ctx context.Context,req *calculatorpb.SumRequest)(*calculatorpb.SumResponse,error){
	fmt.Println("calculator function was invoked",req)

	fnum:=req.FirstNumber
	snum:=req.SecondNumber

	//sum:=fnum+snum
sum:=calculatorpb.SumResponse{
	SumResult: fnum+snum,
}


	return &sum,nil 
}
func (*server)PrimeNumberDecomposition(req *calculatorpb.PrimeNumberDecompositionRequest,
	stream calculatorpb.CalculatorService_PrimeNumberDecompositionServer)error{
	
	fmt.Println("server prime number decompositon invoked  ",req)
		num:=req.GetNumber()

		var  divisior=int64(2)

		for num>1{
			if num%divisior==0{
				stream.Send(&calculatorpb.PrimeNumberDecompositionResponse{
					PrimFactor: divisior,
				})
				num=num/divisior

			}else{
				divisior++
				fmt.Printf("division has increased %v\n",divisior)
			}

		}
return nil
}

func main() {
	fmt.Println("calculator server started")

	lis, err := net.Listen("tcp", "0.0.0.0:50051")
	if err != nil {
		log.Fatalf("failed to listen %v", err)
	}
	s := grpc.NewServer()
	calculatorpb.RegisterCalculatorServiceServer(s,&server{})

	if err:=s.Serve(lis);err!=nil{
		log.Fatalf("filed to serve : %v ",err)
	}

}
