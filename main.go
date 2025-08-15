package main

import (
	"fmt"
	"log"

	"github.com/kotjuz/go_distributed_file_storage/p2p"
)

func main() {
	tr := p2p.NewTCPTransport(":3000")

	if err := tr.ListenAndAccept(); err != nil {
		log.Fatal(err)
	}

	fmt.Println("WE Gucci")
	select {}
}
