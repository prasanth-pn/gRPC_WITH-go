protoc:
	protoc greet/greetpb/*.proto --go_out=plugins=grpc:.
	protoc calculator/calculatorpb/*.proto --go_out=plugins=grpc:.

run:
	go run greet/greet_server/server.go
client:
	go run greet/greet_client/client.go

cals:
	go run calculator/calculator_server/server.go
calc:
	go run calculator/calculator_client/client.go