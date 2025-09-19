package main

import (
	"fmt"
	"log"
	"net"
	"github.com/DanielKoklev/dbs-project/pkg/nbdtarget"
)

func main() {
	fmt.Println("Starting pebble-gateway -- NBD server...")
	ln, err := net.Listen("tcp", ":10809")
	
	if err != nil {
		log.Fatal(err)
	}
	
	backend := nbdtarget.NewMockBackend(2 << 30) // 2 GiB mock
	srv := &nbdtarget.Server{BE: backend}
	log.Fatal(srv.Serve(ln))
}