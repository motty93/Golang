## 問題
### 基本的なinterfaceの定義と実装
Animal interface に Sound() string メソッドを定義し、Dog と Cat の2つの構造体に実装してください。
main関数で両方の Sound() を呼び出して出力します。


## 解説
point: 構造的型付け(structural typing)

Goにはimplementsキーワードが存在しない。
メソッドセットがinterfaceの要求を満たしていれば、自動的に「実装している」とみなされます。
これによって、ライブラリ作者が事前にinterfaceを定義していなくても、利用者側で後付けにinterfaceを定義してラップできます。
これがDuck Typingの安全版ともいえるGoの強みです。
