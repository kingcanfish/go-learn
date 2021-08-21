package main

import (
	"fmt"
	"log"
	"net/rpc"
)

type HelloServiceClient struct {
	*rpc.Client
}

func (h HelloServiceClient) SayHello(request string, reply *string) error {
	err := h.Client.Call(HelloServiceName+".SayHello", request, reply)
	if err != nil {
		return err
	}
	return nil
}

func DialHelloService(network, address string) (*HelloServiceClient, error) {
	c, err := rpc.Dial(network, address)
	if err != nil {
		return nil, err
	}
	return &HelloServiceClient{Client: c}, nil
}

var _ IHelloService = (*HelloServiceClient)(nil)

func main() {
	client, err := DialHelloService("tcp", "localhost:9999")
	if err != nil {
		log.Fatalln("dail err:", err)
	}
	var reply string
	err = client.SayHello("AssadGuo", &reply)
	if err != nil {
		return
	}
	fmt.Println(reply)
	return
}
