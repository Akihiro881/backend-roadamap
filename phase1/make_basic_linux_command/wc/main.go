package main

import (
	"bufio"
	"flag"
	"fmt"
	"io"
	"os"
	"strings"
)

func wc(r io.Reader) (lines, words, bytes int, err error) {
	br := bufio.NewReader(r)
	for {
		line, err := br.ReadString('\n')

		bytes += len(line)
		words += len(strings.Fields(line))
		if strings.HasSuffix(line, "\n") {
			lines++
		}

		if err == io.EOF {
			return lines, words, bytes, nil
		}
		if err != nil {
			return lines, words, bytes, err
		}

	}
}

func format(lines, words, bytes int, showL, showW, showC bool, name string) string {
	var parts []string
	if showL {
		parts = append(parts, fmt.Sprintf("%8d", lines))
	}
	if showW {
		parts = append(parts, fmt.Sprintf("8%d", words))
	}
	if showC {
		parts = append(parts, fmt.Sprintf("%8d", bytes))
	}
	s := strings.Join(parts, "")
	if name != "" {
		s += " " + name
	}
	return s
}

func main() {
	l := flag.Bool("l", false, "行数を表示する")
	w := flag.Bool("w", false, "単語数を表示する")
	c := flag.Bool("c", false, "バイト数を表示する")
	flag.Parse()
	files := flag.Args()

	//オプションが1つも指定されなければ全部表示
	if !*l && !*w && !*c {
		*l, *w, *c = true, true, true
	}

	if len(files) == 0 {
		lines, words, bytes, err := wc(os.Stdin)
		if err != nil {
			fmt.Fprintln(os.Stderr, err)
			os.Exit(1)
		}
		fmt.Fprintln(os.Stdout, format(lines, words, bytes, *l, *w, *c, ""))
		return
	}

	hasError := false
	for _, path := range files {
		f, err := os.Open(path)
		if err != nil {
			fmt.Fprintln(os.Stderr, err)
			hasError = true
			continue
		}
		lines, words, bytes, err := wc(f)
		f.Close()
		if err != nil {
			fmt.Fprintln(os.Stderr, err)
			hasError = true
			continue
		}
		fmt.Fprintln(os.Stdout, format(lines, words, bytes, *l, *w, *c, path))
	}
	if hasError {
		os.Exit(1)
	}
}
