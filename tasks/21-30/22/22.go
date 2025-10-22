package main

import "fmt"

func main() {
	msgChan := make(chan string, 1)
	close(msgChan)
	stopChan := make(chan struct{})
	close(stopChan)

	select {
	case msg, ok := <-msgChan:
		if !ok {
			fmt.Println("channel closed")
			return
		}
		fmt.Println("msg sent", msg)
	case <-stopChan:
		fmt.Println("stop signal recived") // stop signal recived
	}
}
