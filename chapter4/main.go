package main

import (
	"bufio"
	"flag"
	"fmt"
	"io"
	"os"
	"strings"
)

// go-cutコマンドを実装しよう
func main() {
	var (
		dflg = flag.String("d", "\t", "field delimiter character instead of the tab character.")
		fflg = flag.Int("f", 1, "Output fields are separated by a single occurrence of the field delimiter character.")
	)
	flag.Parse()

	var (
		args   = flag.Args()
		reader io.Reader
	)
	if len(args) > 0 {
		var err error
		reader, err = os.Open(args[0])
		if err != nil {
			fmt.Println(err)
		}
	} else {
		reader = os.Stdin
	}
	stdin := bufio.NewScanner(reader)
	for stdin.Scan() {
		text := stdin.Text()
		splitText := strings.Split(text, *dflg)
		if len(splitText) > *fflg-1 {
			fmt.Println(splitText[*fflg-1])
		} else {
			fmt.Println("")
		}
	}
}
