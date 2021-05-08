package main

import (
	"context"
	"net"
	"fmt"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	"github.com/grpc/add/calc"
)

type server struct {}

type Hello struct  {
	a int32
}

func main() {
	lis, err := net.Listen("tcp", ":9090")
	if err != nil {
		panic(err)
	}

	var result []Hello
	result = append(result, Hello{
		a: 32,
	})
	fmt.Printf("Result %v", result)

	fmt.Println("Server running")

	srv := grpc.NewServer()
	calc.RegisterAddServiceServer(srv, &server{})
	reflection.Register(srv)

	if e := srv.Serve(lis); e !=nil {
		panic(e)
	}

	fmt.Println("Server running")
}


var result calc.Response;

func (s *server) Add(ctx context.Context, request *calc.Request)(*calc.Response, error) {
	name, email, phone := request.GetName(), request.GetEmail(), request.GetPhone()

	fmt.Println("Data from clinet %v, %v, %v", name, email, phone)

	obj := &calc.Request{
		Name: name,
		Email: email,
		Phone: phone,
	}

	result.Result = append(result.Result, obj)
	

	return &result, nil
}

func (s *server) GetAll(ctx context.Context, request *calc.Request)(*calc.Response, error) {
	return &result, nil
}