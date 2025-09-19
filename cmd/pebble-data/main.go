package main

import (
	"fmt"
	"log"
	"net/http"
	"github.com/DanielKoklev/dbs-project/pkg/raftmeta"
)

func main() {
	fmt.Println("Starting pebble-meta -- metadata service...")
	http.ListenAndServe(":7001", nil)
	log.Fatal("pebble-meta exited")
}