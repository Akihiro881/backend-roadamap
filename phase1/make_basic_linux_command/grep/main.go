package main

import (
	"bufio"
	"io"
	"strings"
)

func grep(w io.Writer ,r io.Reader,pattern string, color bool) (str, err error) {
	br := bufio.NewReader(r)

	for {
		line, err := br.ReadString('\n')
		i := strings.Index(pattern)

		if err == io.EOF{
			
			fmt.
		}
	}

}
