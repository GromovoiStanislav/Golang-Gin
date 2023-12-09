package main

import (
	"gin-example/http_router"
)

const port int = 8080

func main() {
	finished := make(chan bool)

	go http_router.GinNew(port)

	<-finished
}
