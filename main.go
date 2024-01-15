package main

import (
	"log"

	"github.com/skywalkeretw/master-container-builder/pkg"
)

func main() {
	log.Println("Start")
	pkg.ListenToQueue("rpc_queue")
	log.Println("Fin")

}
