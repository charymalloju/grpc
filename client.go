package main;

import (
	"context"
	"fmt"
	"google.golang.org/grpc"
	"github.com/grpc/add/calc"
)

func main() {
	
	opts := grpc.WithInsecure();
	cc, err:= grpc.Dial("localhost:9090", opts)
	if err != nil {
		panic(err)
	}

	// defer cc.close()

	client := calc.NewAddServiceClient(cc)

	request := &calc.Request{Name:"chary1", Email: "chary@vitwit.com1", Phone: 9912903729}

	resp, _ := client.Add(context.Background(), request)

	fmt.Printf("Response from sever add details %v", resp);

	resp1, _ := client.GetAll(context.Background(), request)

	fmt.Printf("Response from sever %v", resp1);
}