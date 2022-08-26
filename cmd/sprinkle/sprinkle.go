package main

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
	"time"
)

func main() {
	transforms := []string{
		"%sapp",
		"%ssite",
		"lets%s",
		"%stime",
		"get%s",
		"go%s",
		"%shq",
	}

	s := bufio.NewScanner(os.Stdin)

	rand.Seed(time.Now().UTC().UnixNano())

	for s.Scan() {
		t := transforms[rand.Intn(len(transforms))]
		n := fmt.Sprintf(t, s.Text())
		fmt.Fprintf(os.Stdout, "%s\n", n)
	}
}
