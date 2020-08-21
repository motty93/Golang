package main

import (
	"fmt"

	"github.com/markcheno/go-quote"
	"github.com/markcheno/go-talib"
)

func main() {
	// yahoo株価spyを取得
	spy, _ := quote.NewQuoteFromYahoo(
		"spy", "2019-01-01", "2019-12-31", quote.Daily, true)
	fmt.Println(spy.CSV())
	rsi2 := talib.Rsi(spy.Close, 2)
	fmt.Println(rsi2)
	mva := talib.Ema(spy.Close, 14)
	fmt.Println(mva)
}
