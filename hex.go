package main

import (
	"bufio"
	"encoding/hex"
	"fmt"
	"io"
	"os"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println("you must supply a file name")
		os.Exit(1)
	}

	fp, err := os.Open(os.Args[1])
	if err != nil {
		panic(err)
	}

	reader := bufio.NewReader(fp)
	w := hex.Dumper(os.Stdout)

	buf := make([]byte, 1024)
	for {
		n, err := reader.Read(buf)
		if err != nil && err != io.EOF {
			panic(err)
		}
		if n == 0 {
			break
		}
		w.Write(buf[:n])
	}
}
