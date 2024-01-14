package main

import (
	"github.com/skywalkeretw/master-container-builder/pkg"
)

func main() {
	rabbit := pkg.NewRabbitMQ()
	rabbit.Dial()
	rabbit.ReceiveMessages("rpc_queue")

}
