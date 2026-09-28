package main

import (
	"bufio"
	"flag"
	"fmt"
	"io"
	"os"
	"strings"
)

func main() {
	flag.Parse()

	if flag.NArg() < 1 {
		fmt.Fprintln(os.Stderr, "usage: mygrep [options] pattern [file...]")
		os.Exit(2)
	}

	color := isTerminal(os.Stdout)

	pattern := flag.Arg(0)
	files := flag.Args()[1:]

	hasError := false
	for _, path := range files {
		f, err := os.Open(path)

		if err != nil {
			fmt.Fprintln(os.Stderr, err)
			hasError = true
			continue
		}
		grep(os.Stdout, f, pattern, color)

		f.Close()
	}
	if hasError {
		os.Exit(1)
	}

}

func grep(w io.Writer, r io.Reader, pattern string, color bool) (ismatch bool, err error) {
	br := bufio.NewReader(r)

	matched := false

	for {

		line, err := br.ReadString('\n')
		if color {
			positions := findAll(line, pattern)

			if len(positions) > 0 { //一致する単語がある行
				colorline := colorize(line, positions, len(pattern))
				w.Write([]byte(colorline))
				matched = true
			}
		} else {
			if strings.Contains(line, pattern) {
				w.Write([]byte(line))
				matched = true
			}
		}

		if err == io.EOF {
			return matched, nil
		}

		if err != nil {
			return matched, err
		}
	}
}

func findAll(line string, pattern string) []int {
	if pattern == "" {
		return nil
	}
	positions := []int{}
	start := 0
	for {

		i := strings.Index(line[start:], pattern)

		if i == -1 {
			return positions
		}

		positions = append(positions, start+i)
		start += i + len(pattern)
	}
}

func colorize(line string, positions []int, patLen int) string {
	const (
		colorStart = "\033[1;31m"
		colorEnd   = "\033[0m"
	)
	var b strings.Builder
	prev := 0
	for _, p := range positions {
		b.WriteString(line[prev:p])
		b.WriteString(colorStart)
		b.WriteString(line[p : p+patLen])
		b.WriteString(colorEnd)
		prev = p + patLen
	}
	b.WriteString(line[prev:])
	return b.String()
}

func isTerminal(f *os.File) bool {
	fi, err := f.Stat()
	if err != nil {
		return false
	}
	return fi.Mode()&os.ModeCharDevice != 0
}
