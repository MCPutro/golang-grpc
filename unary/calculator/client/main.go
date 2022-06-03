package main

import (
	"context"
	"fmt"
	"go-grpc-example2/unary/calculator/proto"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
	"log"
)

func main() {
	tls := true
	var opts []grpc.DialOption

	if tls {
		certFile := "ssl/ca.crt"
		clientTLSFromFile, err := credentials.NewClientTLSFromFile(certFile, "")

		if err != nil {
			log.Fatalf("error when load ca %v", err)
		}

		opts = append(opts, grpc.WithTransportCredentials(clientTLSFromFile))

	}

	conn, err := grpc.Dial("localhost:50051", opts...)
	if err != nil {
		panic(err)
	}
	defer func(conn *grpc.ClientConn) {
		err2 := conn.Close()
		if err2 != nil {
			panic(err2)
		}
	}(conn)

	client := proto.NewCalculatorOperationClient(conn)

	Add(client)
	Subtract(client)
	Multiply(client)
	// Divide(client)
	divide, err2 := client.Divide(context.Background(), &proto.DivideRequest{
		VarA: 15.5,
		VarB: 0,
	})

	if err2 != nil {
		fmt.Println(err2)
	} else {
		fmt.Println("device ", divide.Result)
	}

}

func Add(client proto.CalculatorOperationClient) {
	request := proto.AddRequest{
		VarA: 12.5,
		VarB: -3,
	}

	//call func Add in server
	response, err := client.Add(context.Background(), &request)
	if err != nil {
		panic(err)
	}

	fmt.Println(request.GetVarA(), "Add", request.GetVarB(), "=", response.Result)
}

func Subtract(client proto.CalculatorOperationClient) {
	request := proto.SubtractRequest{
		VarA: 15.5,
		VarB: 0,
	}

	//call func Subtract in server
	response, err := client.Subtract(context.Background(), &request)
	if err != nil {
		panic(err)
	}

	fmt.Println(request.GetVarA(), "Subtract", request.GetVarB(), "=", response.Result)
}

func Multiply(client proto.CalculatorOperationClient) {
	request := proto.MultiplyRequest{
		VarA: 15.5,
		VarB: 0,
	}

	//call func Multiply in server
	response, err := client.Multiply(context.Background(), &request)
	if err != nil {
		panic(err)
	}

	fmt.Println(request.GetVarA(), "Multiply", request.GetVarB(), "=", response.Result)
}

func Divide(client proto.CalculatorOperationClient) {
	request := proto.DivideRequest{
		VarA: 15.5,
		VarB: 3,
	}

	//call func Divide in server
	response, err := client.Divide(context.Background(), &request)
	if err != nil {
		panic(err)
	}

	fmt.Println(request.GetVarA(), "Divide", request.GetVarB(), "=", response.Result)
}
