# Golang
## 型(type)
* bool
* int8/int16/int32/int64
* uint8/uint16/uint32/uint64
* float32/float64
* complex64/complex128
* byte: 1バイトデータ(uint8と同義)
* rune: 1文字(int32と同義)
* uint: uint32 or uint64
* int: int32 or int64
* uintptr: ポインタを表現するのに十分な非負整数
* string

## 演算子
* x + y    加算 (文字列の連結にも利用)
* x - y    減算
* x * y    乗算
* x / y    除算
* x % y    除算の余り
* x & y    論理積(AND)
* x | y    論理和(OR)
* x ^ y    排他的論理和(XOR)
* x &^ y    x AND (NOT y)
* x << y    yビット左にシフト
* x >> y    yビット右にシフト
* x = y    xにyを代入
* x := y    xにyを代入(初期化の使用可能)
* x++    x = x + 1 と同義
* x--    x = x - 1 と同義
* x += y    x = x + y と同義
* x -= y    x = x - y と同義
* x *= y    x = x * y と同義
* x /= y    x = x / y と同義
* x %= y    x = x % y と同義
* x &= y    x = x & y と同義
* x |= y    x = x | y と同義
* x ^= y    x = x ^ y と同義
* x &^= y    x = x &^ y と同義
* x <<= y    x = x << y と同義
* x >>= y    x = x >> y と同義
* x && y    xかつy(AND)
* x || y    xまたはy(OR)
* !x    xがtrueの場合false/falseの場合true(NOT)
* x == y    xとyが等しければ
* x != y    xとyが等しくなければ
* x < y    yがxより大きければ
* x < y    yがx以上であれば
* x > y    yがxより小さければ
* x >= y    yがx以下であれば
* ch <- x    chチャネルにxを送信
* x =<- ch  chチャネルからxに受信

## リテラル・値

* nil  無しを示す特別な値
* true  真偽値の真
* false  真偽値の偽
* 1234  整数
* 1_234  整数(カンマ区切りの代わりに_を使用。_は無視される)
* 0777  8進数
* 0o755  8進数(0Oも可)
* 0x89ab  16進数(0Xも可)
* 0b1111  2進数(0Bも可)
* 123.4  小数
* 1.23e4  浮動小数点数(1.23E4も可)
* 1.23i  虚数
* "ABC"  文字列
* 'A'  文字(rune)

## エスケープシーケンス

* \a    ベル(U+0007)
* \b    バックスペース(U+0008)
* \t    タブ(U+0009)
* \n    改行(U+000A)
* \v    垂直タブ(U+000B)
* \f    フォームフィード(U+000C)
* \r    キャリッジリターン(U+000D)
* \"    ダブルクォート(U+0022)
* \'    シングルクォート(U+0027)
* \\    バックスラッシュ(U+005C)
* \x42    ASCII文字(U+0000～U+00FF)
* \u30A2    Unicode(U+0000～U+FFFF)
* \U0001F604  Unicode(U+0000～U+10FFFF)

## 配列(array)
個数が決まっていて変更不可のものを配列という
個数変更できないがメモリ効率がよい

```
a1 := [3]string{}
a1[0] = "red"
a2[1] = "blue"
a3[2] = "green"
```

初期化時にあたりを設定することもできる。
初期化で個数が決まる場合は...と省略可能(3と明示してもおｋ)

```
a1 := [3]string{"red", "blue", "green"}
b1 := [...]string{"red", "blue", "green"}
```

## スライス(slice)
可変な配列のことをスライスと呼ぶ
不可変な配列よりもメモリ効率や速度は若干落ちる
```
a1 := []string{}
// apendで追加
a1 = append(a1, "red")
a1 = append(a1, "green")
a1 = append(a1, "black")
a1 = append(a1, "blue")

// lenは長さ、capは容量を求める
fmt.Println(len(a1), cap(a1))
```

makeを用いてメモリの確保が可能。
容量超過時の再確保を減らして速度を早めることができる

```
a2 := make([]string, 0, 1024)
a2 = append(a2, "red")
a2 = append(a2, "green")
a2 = append(a2, "black")
a2 = append(a2, "blue")
```

### スライスの操作
`slice_array.go`

|操作|意味|
|---|---|
|Slice[start:end]|startからend-1まで|
|Slice[start:]|startから末尾まで|
|Slice[:end]|先頭からend-1まで|
|Slice[:]|先頭から最後まで|

## Struct(構造体)
クラスの代わり。
構造体にはメンバ変数のみを定義し、クラスメソッドに相当する関数は関数名の前に(thisに相当する変数 *構造体名)をつけて定義する。

## interface(インタフェース)
インタフェースはポリモーフィズムを実装するための機能
`interface.go`を参照

### interface {}型
型のないインタフェース`interface {}`はany型のように使用可能
`interface2.go`を参照

```
// 本来の定義方法
p1 := map[string]int{
  "x": 10,
  "y": 200,
}

// valueを可変に定義できる
p2 := map[string]interface{}{
  "name": "motty",
  "age":  90,
}
```

interface{}はanyの用に使えるという特徴を活かし、任意の型の値を持つマップを定義することができる。


## ポインタ(pointer)

変数が格納されているメモリのアドレス。C言語と同様に、オブジェクトのポインタを参照するには`&`を、ポインタの中身を参照するには`*`を用いる。

変数の値を渡すことを「値渡し」
変数のポインタを渡すことを「参照渡し」と呼ぶ。
値渡しでは値のコピーしか渡していないので、元の変数を変更することは出来ませんが、ポインタ渡しであれば関数の中で変数の値を変更することが可能。


## 領域確保(new)
`new()`を用いて領域を動的に確保し、その領域へのポインタを得ることができる。
確保した領域は参照されなくなったあとにガーベジコレクションにより自動的に解放される。

