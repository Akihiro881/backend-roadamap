package main

import (
	"fmt"
	"os"
)

//cat wc grep のクローンを作る → 標準入力とファイルの両方に対応
//次コマンド対応：cat -n、wc -l -w -c -m、grep -i -v -n

func check(filepath string, e error) bool {
	if e != nil {
		fmt.Printf("cat: %s: No such file or directory", filepath)
		return false
	}

	return true
}

func cat(filepath string) bool {
	data, err := os.ReadFile(filepath)
	if check(filepath, err) {
		fmt.Println("no file")
		return false
	}
	fmt.Println(string(data))
	return true
}

func main() {

	var command string = os.Args[1]
	var filepath string = os.Args[2]
	fmt.Println(command, filepath)
	if command == "cat" {
		fmt.Println("test")
		cat(filepath)
	}
}
