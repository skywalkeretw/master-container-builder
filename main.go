package main

import (
	"github.com/skywalkeretw/master-container-builder/pkg"
)

func main() {
	pkg.ListenToQueue("rpc_queue")

}
