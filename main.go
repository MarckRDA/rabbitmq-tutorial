package main

import (
	helloworld "github.com/MarckRDA/rabbitmq-tutorial/hello-world"
)

func main() {
	helloworld.Send()
	helloworld.Receive()
}
