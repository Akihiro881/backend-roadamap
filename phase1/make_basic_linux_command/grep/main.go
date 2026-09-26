package main

import (
	"bufio"
	"io"
	"strings"
)

func grep(w io.Writer, r io.Reader, pattern string, color bool) (str, text, err error) {
	br := bufio.NewReader(r)

	wordPoint := []int{}
	for {
		//readstringで読むんじゃなくてreadで読んだほうがよい？
		//bufで読むと単語の途中で切れる可能性あるからなし
		//行単位で読んでその行に対応するインデックスを格納する
		line, err := br.ReadString('\n')
		tansaku(line, pattern, wordPoint, false)

		if err == io.EOF {
			return wordPoint, text, nil
		}

		if err != nil {
			return wordPoint, text, nil
		}
	}
}

func tansaku(line string, pattern string, wordPoint []int, repeatflg bool) wordPoint {
	if strings.Index(line, pattern) == -1 {
		return wordPoint
	}

	i := strings.Index(line, pattern)
	//探索2回目以降
	if repeatflg {
		s := wordPoint[len(wordPoint)-1] + len(pattern) + i
	}

	if repeatflg {
		wordPoint.append(wordPoint, s)
	} else {
		wordPoint.append(wordPoint, i)
	}
	tansaku(line[i+len(pattern):], pattern, wordPoint, true)

}
