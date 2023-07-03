package main

import (
	"context"
	"errors"
	"flag"

	example "github.com/rpcxio/rpcx-examples"
	"github.com/rpcxio/rpcx-examples/error_customized"
	"github.com/smallnest/rpcx/client"
	"github.com/smallnest/rpcx/server"
)

var (
	addr = flag.String("addr", "localhost:8972", "server address")
)

type Arith struct{}

// the second parameter is not a pointer
func (t *Arith) Mul(ctx context.Context, args example.Args, reply *example.Reply) error {

	d, _ := client.NewPeer2PeerDiscovery("tcp@"+*addr, "")
	xclient := client.NewXClient("Arith", client.Failtry, client.RandomSelect, d, client.DefaultOption)
	defer xclient.Close()

	var s string
	// err := xclient.Call(ctx, "Say", "hello", &s)
	// if err != nil {
	// 	fmt.Println(err)
	// }
	xclient.Call(ctx, "Say", "hello", &s)

	return error_customized.NewError(100, "error from Mul")
}

func (t *Arith) Say(ctx context.Context, args string, reply *string) error {

	return errors.New("error from Say")
}

func main() {
	flag.Parse()

	s := server.NewServer()
	//s.Register(new(Arith), "")
	s.RegisterName("Arith", new(Arith), "")
	err := s.Serve("tcp", *addr)
	if err != nil {
		panic(err)
	}
}
