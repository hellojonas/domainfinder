package main

import (
	"bufio"
	"fmt"
	"log"
	"os"

	"github.com/hellojonas/domyne/pkg/thesarus"
)

func main() {
	bht := thesarus.New(os.Getenv("BHT_KEY"))
	s := bufio.NewScanner(os.Stdin)

	for s.Scan() {
		t := s.Text()

		syns, err := bht.Synonyms(t)

		if err != nil {
			log.Fatalln(err.Error())
		}

		for _, syn := range syns {
			fmt.Fprintln(os.Stdout, syn)
		}
	}
}
