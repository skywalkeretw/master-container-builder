package main

import (
	"log"

	"github.com/skywalkeretw/master-container-builder/pkg"
)

func main() {
	log.Println("Start")
	pkg.ListenToQueue("imagebuilder")
	log.Println("Fin")

}
