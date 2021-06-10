package main

import (
	"errors"
	"fmt"
	"os"
	"strings"
)

func echo(args []string) (txt string, err error) {
	if len(args) < 2 {
		return "", errors.New("no message to echo")
	}
	return strings.Join(args[1:], " "), nil
}
func main() {
	txt, err := echo(os.Args)
	if err == nil {
		fmt.Println(txt + "najs lan")
	} else {
		fmt.Fprintf(os.Stderr, "%+v\n", err)
		os.Exit(1)
	}
}
