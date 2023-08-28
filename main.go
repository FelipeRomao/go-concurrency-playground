package main

import (
	"fmt"
	"time"
)

func infiniteCount(thing string) {
	for i := 1; true; i++ {
		fmt.Println(i, thing)
		time.Sleep(time.Second * 1)
	}
}

func main() {
	go infiniteCount("sheep")
	infiniteCount("fish")
}
