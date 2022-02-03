package main

import (
	"fmt"
	"strings"
	"time"
)

func say(s string) {
	for i := 0; i < 5; i++ {
		fmt.Println(s)
	}
}

func spell(s string) {
	strs := strings.Split(s, " ")
	fmt.Println(len(strs))
	for _, str := range strs {
		go say(str)
	}
}

func main() {
	go spell("Channels are a typed conduit through which you can send and receive values with the channel operator, <-.")

	time.Sleep(1_000_000_000)
}
