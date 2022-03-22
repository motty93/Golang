## contextによる解決
「複数ゴルーチン間で安全に、そして簡単に情報伝達を行いたい」ときはチャネルだけで伝達を実現するのは難しい


- 処理の締め切りを伝達
- キャンセル信号の伝播
- リクエストスコープ値の伝達
の３つについて、ゴルーチン上で起動される関数の第一引数に、`context.Context`型を１つ渡すだけで簡単に実現できる


## contextの定義

```go
type Context interface {
  Deadline() (deadline time.Time, ok bool)
  Done() <-chan struct{}
  Err() error
  Value(key interface{}) interface{}
}
```
