package main

import (
	"flag"
	"log"
	"time"

	"github.com/rpcxio/rpcx-etcd/serverplugin"
	example "github.com/rpcxio/rpcx-examples"
	"github.com/smallnest/rpcx/server"
)

var (
	addr     = flag.String("addr", "localhost:8972", "server address")
	etcdAddr = flag.String("etcdAddr", "localhost:2379", "etcd address")
	basePath = flag.String("base", "/rpcx_test", "prefix path")
)

func main() {
	flag.Parse()

	s := server.NewServer()
	addRegistryPlugin(s)

	s.RegisterName("Arith", new(example.Arith), "")
	err := s.Serve("tcp", *addr)
	if err != nil {
		panic(err)
	}
}

func addRegistryPlugin(s *server.Server) {
	r := &serverplugin.EtcdV3RegisterPlugin{
		ServiceAddress: "tcp@" + *addr,
		EtcdServers:    []string{*etcdAddr},
		BasePath:       *basePath,
		UpdateInterval: time.Second,
		Expired:        time.Second * 2,
	}
	err := r.Start()
	if err != nil {
		log.Fatal(err)
	}
	s.Plugins.Add(r)
}
