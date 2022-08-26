package main

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
	"strings"
	"time"
)

var (
	dupvowel = true
	rmvowel  = false
)

func randbool() bool {
	return rand.Intn(2) == 1
}

func main() {
	rand.Seed(time.Now().UTC().UnixNano())
	s := bufio.NewScanner(os.Stdin)

	for s.Scan() {
		w := []byte(s.Text())
		vi := -1

		for i, v := range w {
			switch strings.ToLower(string(v)) {
			case "a", "e", "i", "o", "u":
				if randbool() {
					vi = i
					if randbool() {
						break
					}
				}
			}
		}

		if vi >= 0 {
			switch randbool() {
			case dupvowel:
				w = append(w[:vi+1], w[vi:]...)
			case rmvowel:
				w = append(w[:vi], w[vi+1:]...)
			}
		}

		fmt.Fprintln(os.Stdout, string(w))
	}
}
