package main

import (
	"log"

	"github.com/skywalkeretw/master-container-builder/pkg"
)

func main() {
	log.Println("Start")
	pkg.ListenToQueue("imageBuilder")
	log.Println("Fin")

}
