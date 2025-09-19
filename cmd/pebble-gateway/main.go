package main

import (
	"flag"
	"fmt"
	"log"
	"github.com/DanielKoklev/dbs-project/pkg/wal"
)

func main() {
	id := flag.String("id", "nil", "node id")
	flag.Parse()
	fmt.Printf("Starting pebble-data node %s...\n", *id)
	wal.Init(*id)
	log.Fatal("pebble-data exited")
}
