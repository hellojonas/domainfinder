package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
	"strings"
)

func main() {
	marks := map[bool]string{
		true:  "✔",
		false: "✘",
	}
	sc := bufio.NewScanner(os.Stdin)

	for sc.Scan() {
		t := sc.Text()
		ok, err := exists(t)

		if err != nil {
			fmt.Fprintln(os.Stderr, err)
			continue
		}

		fmt.Fprintf(os.Stdout, "%-32s %-2v\n", t, marks[!ok])
	}
}

func exists(domain string) (bool, error) {
	whois := "whois.verisign-grs.com" // supports .com, .net & .edu tlds
	match := "no match for"
	conn, err := net.Dial("tcp", whois+":43")

	if err != nil {
		return false, fmt.Errorf("Failed to open connection on '%s'. %s", whois, err.Error())
	}

	defer conn.Close()
	_, err = conn.Write([]byte(domain + "\r\n"))

	if err != nil {
		return false, fmt.Errorf("failed to write data to whois server. %s", err.Error())
	}

	sc := bufio.NewScanner(conn)

	for sc.Scan() {
		t := sc.Text()

		if strings.Contains(strings.ToLower(t), match) {
			return false, nil
		}
	}

	return true, nil
}
