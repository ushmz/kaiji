package main

import (
	"flag"
	"fmt"
	"math/rand"
	"time"
)

func main() {
	flag.Parse()
	args := flag.Args()

	cheat := make([]string, 0)
	cheatlen := len(args) * (len(args) + 1) / 2
	for k, v := range args {
		for i := 0; i <= k; i++ {
			cheat = append(cheat, v)
		}
	}
	rand.Seed(time.Now().UnixNano())
	fmt.Println(cheat[rand.Intn(cheatlen)])
}
