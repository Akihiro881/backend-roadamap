# backend-roadamap

## 1a-2
## 設計メモ

### 入力を io.Reader で受ける理由

`os.ReadFile(path)` ではなく `os.Open` + `io.Reader` にした。

- 標準入力には「パス」がないので、パスを受け取る設計だと
  `echo hi | ./mycat` が書けない。`os.Stdin` も `*os.File` = `io.Reader` なので、
  `io.Reader` で受ければファイルと標準入力を同じ関数で扱える
- `os.ReadFile` は全部メモリに載せる。`io.Copy` は32KB程度の
  バッファを使い回すので、ファイルサイズに関係なくメモリ使用量が一定

### bufio.Scanner をやめた理由

`-n` の実装で最初 `bufio.Scanner` を使ったが、以下で本物と出力が変わった。

    printf 'abc' > noeol.txt
    diff <(./mycat -n noeol.txt) <(cat -n noeol.txt)

`Scanner.Text()` は改行を取り除いた行を返すため、
"abc" と "abc\n" が区別できない。自分で `\n` を付け直すと、
元のファイルにない改行が増える。

`bufio.Reader.ReadString('\n')` は改行を含んだまま返すのでこれが起きない。
あわせて、Scannerの64KB上限（`bufio.MaxScanTokenSize`）も回避できる。

### -n の出力書式

`cat -A` で本物を調べたところ、右詰め6桁 + タブ + 本文だった。
空白ではなくタブなので `%6d\t%s` とする。