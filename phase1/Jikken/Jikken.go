package main

import (
	"fmt"
	"io"
	"os"
	"strings"
)

func Cat(w io.Writer, r io.Reader) error {
	_, err := io.Copy(w, r)
	return err
}

func main() {
	// 文字列を渡す
	Cat(os.Stdout, strings.NewReader("これは文字列\n"))

	// ファイルを渡す（同じ関数）
	f, err := os.Open("/home/akky/dev/backend-roadmap/phase1/Jikken/a.txt")
	if err != nil {
		fmt.Println("エラー:", err)
		return
	}
	defer f.Close()
	Cat(os.Stdout, f)
}
