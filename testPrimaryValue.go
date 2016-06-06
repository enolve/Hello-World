package main

import (
	"flag"
	"fmt"
	"os"
	"strconv"
)

var goal int //求此数以内的素数

/*
	第一个协程：将不能被2整数的数添加到第二个协程的channel中，第一个数为基数(3)
	第二：将不能被3整除的数添加到下个协程中，以此循环
*/
func primetask(c chan int) {
	p := <-c
	if p > goal {
		os.Exit(0)
	}
	fmt.Println(p)

	nc := make(chan int)
	go primetask(nc)

	for {
		i := <-c
		if i%p != 0 {
			nc <- i
		}
	}
}

func main() {
	flag.Parse()

	args := flag.Args()
	if args != nil && len(args) > 0 {
		var err error
		goal, err = strconv.Atoi(args[0])
		if err != nil {
			goal = 100
		}
	} else {
		goal = 100
	}

	fmt.Println("goal=", goal)
	c := make(chan int)
	go primetask(c)
	for i := 2; ; i++ {
		c <- i  
	}
}
