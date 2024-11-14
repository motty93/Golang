package main

import "fmt"

type Candle struct {
	open, close, high, low int
}

func (c *Candle) Reset() {
	c.open = 0
	c.close = 0
	c.high = 0
	c.low = 0
}

func main() {
	candles := map[int]*Candle{}
	count := 5
	for i := 0; i < count; i++ {
		candles[i] = &Candle{open: i}
	}

	for i, candle := range candles {
		candle.Reset()
		candles[i] = &Candle{
			open: 100,
		}
		fmt.Println(candle)
	}

	for _, candle := range candles {
		fmt.Println(candle.open)
	}
}

func resetCandle() {
}
