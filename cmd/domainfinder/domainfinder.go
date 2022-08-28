package main

import (
	"log"
	"os"
	"os/exec"
)

var cmdchain = []*exec.Cmd{
	exec.Command("synonym"),
	exec.Command("sprinkle"),
	exec.Command("coolify"),
	exec.Command("domainify"),
	exec.Command("available"),
}

func main() {
	cmdchain[0].Stdin = os.Stdin
	cmdchain[len(cmdchain)-1].Stdout = os.Stdout

	for i := 0; i < len(cmdchain)-1; i++ {
		curcmd := cmdchain[i]
		nextcmd := cmdchain[i+1]
		stdout, err := curcmd.StdoutPipe()

		if err != nil {
			log.Fatalln(err)
		}

		nextcmd.Stdin = stdout
	}

	for _, cmd := range cmdchain {
		if err := cmd.Start(); err != nil {
			log.Fatalf("failed starting command chain. %s", err.Error())
			return
		}

		defer cmd.Process.Kill()
	}

	for _, cmd := range cmdchain {
		if err := cmd.Wait(); err != nil {
			log.Fatalf("error waiting command chain. %s", err.Error())
		}
	}
}
