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
	filepathes := os.Args[1:] // 引数だけを切り出す

	fmt.Println(command)
	fmt.Println("以下catコマンドだよ")
	if command == "cat" {
		for _, filepath := range filepathes {

			//fmt.Println(filepath)

			f, err := os.Open(filepath)
			if err != nil {
				fmt.Println(err)
			}
			cat(os.Stdout, f)
			//fmt.Println(caterr)
		}

	}

}
