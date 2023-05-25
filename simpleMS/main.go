package main

import (
	"bufio"
	"context"
	"fmt"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/grpc/grpclog"
	"net"
	"os"
	"simpleMS/reverse"
)

func Server() {

	listener, err := net.Listen("tcp", ":5300")

	if err != nil {
		grpclog.Fatalf("failed to listen: %v", err)
	}

	opts := []grpc.ServerOption{}
	grpcServer := grpc.NewServer(opts...)

	reverse.RegisterReverseServer(grpcServer, &server{})
	grpcServer.Serve(listener)
}

type server struct {
	reverse.UnimplementedReverseServer
}

func (s *server) Do(c context.Context, request *reverse.Request) (response *reverse.Response, err error) {
	n := 0
	// Сreate an array of runes to safely reverse a string.
	runes := make([]rune, len(request.Message))

	for _, r := range request.Message {
		runes[n] = r
		n++
	}

	// Reverse using runes.
	runes = runes[0:n]

	for i := 0; i < n/2; i++ {
		runes[i], runes[n-1-i] = runes[n-1-i], runes[i]
	}

	output := string(runes)
	response = &reverse.Response{
		Message: output,
	}

	return response, nil
}

func Client() {
	opts := []grpc.DialOption{
		grpc.WithTransportCredentials(insecure.NewCredentials()),
	}
	fmt.Println("Write some text below(if you want to stop - type: exit")
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {

		args := scanner.Text()
		if args == "exit" {
			return
		}
		conn, err := grpc.Dial("127.0.0.1:5300", opts...)

		if err != nil {
			grpclog.Fatalf("fail to dial: %v", err)
		}

		defer conn.Close()

		client := reverse.NewReverseClient(conn)
		request := &reverse.Request{
			Message: args,
		}
		response, err := client.Do(context.Background(), request)

		if err != nil {
			grpclog.Fatalf("fail to dial: %v", err)
		}

		fmt.Println(response.Message)
	}
}

func main() {
	go Server()
	Client()
}
