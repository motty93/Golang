## 問題
### ポリモーフィズム
Shape interface に Area() float64 メソッドを定義し、Rectangle（幅・高さ）と Circle（半径）に実装してください。
PrintArea(s Shape) 関数を作り、[]Shape をループして全ての面積を出力してください。


## 解説
PrintArea は具象型を知りません。新しい図形（例：Triangle）を追加しても PrintArea 側の修正は不要です。これがinterfaceによる疎結合の本質で、Open/Closed原則（拡張に開かれ、修正に閉じている）にも合致します。
ただし注意点として、[]Shape のように interface型のスライスは、内部で (type, value) のペアを格納するため、純粋な []Rectangle に比べてメモリ・キャッシュ効率は落ちます。
