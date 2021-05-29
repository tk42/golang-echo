package main

import (
	"flag"
	"fmt"
	"net"

	"golang.org/x/net/context"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"

	gpb "github.com/tk42/golang-echo/proto/autogen/helloworld"
)

var (
	addr = flag.String("addr", "0.0.0.0:5000", "address to connect to")
)

type server struct{ gpb.UnimplementedGreeterServer }

// SayHelloメソッドを実装
func (s *server) SayHello(ctx context.Context, in *gpb.HelloRequest) (*gpb.HelloReply, error) {
	fmt.Printf("Received: %v", in.Name)
	return &gpb.HelloReply{Message: "Hello " + in.Name}, nil
}

func run(ctx context.Context) error {
	fmt.Printf("starting listening %s\n", *addr)
	l, err := net.Listen("tcp", *addr)
	if err != nil {
		fmt.Errorf("Failed to listen %s: %v\n", *addr, err)
		return err
	}

	s := grpc.NewServer()
	gpb.RegisterGreeterServer(s, &server{})
	reflection.Register(s) // mandatory. https://qiita.com/otanu/items/09ef76f2e11b75479105

	fmt.Printf("starting server\n")
	ch := make(chan error, 1)
	go func(ch chan<- error) {
		defer close(ch)
		ch <- s.Serve(l)
	}(ch)
	fmt.Printf("Serving on %s\n", *addr)

	return <-ch
}

func main() {
	flag.Parse()

	ctx := context.Background()
	if err := run(ctx); err != nil {
		panic(err)
	}
}
