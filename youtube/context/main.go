package main

import (
	"context"
	"fmt"
	"log"
	"time"
)

func main() {
	start := time.Now()
	ctx := context.Background()
	val, err := fetchUserData(ctx, 1)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("result:", val)
	fmt.Println("took:", time.Since(start))
}

type Response struct {
	value int
	err   error
}

func fetchUserData(ctx context.Context, userID int) (int, error) {
	ctx, cancel := context.WithTimeout(ctx, time.Millisecond*200)
	defer cancel()
	respch := make(chan Response)

	go func() {
		val, err := fetchThirdPartyStuffWhichCanBeSlow()
		respch <- Response{val, err}
	}()

	for {
		select {
		case <-ctx.Done():
			return 0, fmt.Errorf("fetching data from third party took to long.")
		case res := <-respch:
			return res.value, res.err
		}
	}
}

func fetchThirdPartyStuffWhichCanBeSlow() (int, error) {
	time.Sleep(time.Millisecond * 500)

	return 666, nil
}
