package main

import (
	"log"
	"net"
	"net/rpc"
)

const (
	HelloServiceName = "HelloService"
)

type IHelloService interface {
	SayHello(request string, reply *string) error
}

type HelloService struct{}

func (s HelloService) SayHello(request string, reply *string) error {
	*reply = "hello" + request
	return nil
}

func RegisterHelloService(svc IHelloService) error {
	return rpc.RegisterName(HelloServiceName, svc)
}
func main() {

	err := RegisterHelloService(new(HelloService))
	if err != nil {
		return
	}

	listener, err := net.Listen("tcp", ":9999")
	if err != nil {
		log.Fatalln(err)
	}

	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Fatalln("accept error:", err)
		}
		go rpc.ServeConn(conn)
	}
}
