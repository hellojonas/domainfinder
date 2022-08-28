package main

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
	"strings"
	"time"
	"unicode"
)

func main() {
	rand.Seed(time.Now().UTC().UnixNano())
	tlds := []string{".com", ".net"}
	s := bufio.NewScanner(os.Stdin)
	allowed := "abcdefghijklmnopqrstuvwxyz1234567890-_"

	for s.Scan() {
		t := s.Text()
		var newt []rune

		for _, r := range t {
			if unicode.IsSpace(r) {
				r = '-'
			}
			if strings.ContainsRune(allowed, r) {
				newt = append(newt, r)
			}
		}

		d := fmt.Sprintf("%s%s", string(newt), tlds[rand.Intn(len(tlds))])
		fmt.Fprintf(os.Stdout, "%s\n", d)
	}
}
