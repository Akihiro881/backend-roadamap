package main

import (
	"bufio"
	"flag"
	"fmt"
	"io"
	"os"
)

//cat wc grep のクローンを作る → 標準入力とファイルの両方に対応
//次コマンド対応：cat -n、wc -l -w -c -m、grep -i -v -n

func cat(w io.Writer, r io.Reader) error {
	buf := make([]byte, 3*1024)
	for {
		n, err := r.Read(buf)
		if n > 0 {
			w.Write(buf[:n])
		}
		if err == io.EOF {
			return nil
		}
		if err != nil {
			return err
		}

	}
}

func catn(w io.Writer, r io.Reader) error {
	//sc := bufio.NewScanner(r)
	br := bufio.NewReader(r)
	n := 1
	/**
	for sc.Scan() {
		str := sc.Text()
		fmt.Fprintf(w, "%6d\t%s\n", n, str)
		n++
	}
	return sc.Err()
	**/
	for {
		line, err := br.ReadString('\n')
		fmt.Fprintf(w, "%6d\t%s", n, line)
		if err != nil {
			return err
		}
		if err == io.EOF {
			return nil
		}
		n++
	}
}

func main() {

	n := flag.Bool("n", false, "行番号を表示する")

	flag.Parse()
	files := flag.Args()

	for _, path := range files {
		f, err := os.Open(path)
		if err != nil {
			fmt.Fprintln(os.Stderr, err)
			continue
		}

		if *n {
			catn(os.Stdout, f)
		} else {
			cat(os.Stdout, f)
		}

		f.Close()
	}

}
