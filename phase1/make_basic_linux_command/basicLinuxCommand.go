package main

import (
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

func main() {

	var command string = os.Args[1]
	var filepath string = os.Args[2]
	fmt.Println(command)
	fmt.Println(filepath)
	if command == "cat" {
		f, err := os.Open(filepath)
		if err != nil {
			fmt.Println(err)
		}

		caterr := cat(os.Stdout, f)
		fmt.Println(caterr)
	}

}
