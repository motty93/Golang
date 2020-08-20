package main

import (
	"context"
	"fmt"
	"time"

	"golang.org/x/sync/semaphore"
)

// 2つ許可する
var s *semaphore.Weighted = semaphore.NewWeighted(2)

func longProcess(ctx context.Context) {
	// 処理を終えたい場合はTryAcquireを使う
	isAcquire := s.TryAcquire(1)
	if !isAcquire {
		fmt.Println("could not get lock")
		return
	}
	// 2つ同時に走らせて、もう1つは止める
	// if err := s.Acquire(ctx, 1); err != nil {
	// 	fmt.Println(err)
	// 	return
	// }
	defer s.Release(1)

	fmt.Println("wait...")
	time.Sleep(1 * time.Second)
	fmt.Println("done")
}

// semaphoreはgoroutineが走っている数を限定できる
func main() {
	ctx := context.TODO()
	go longProcess(ctx)
	go longProcess(ctx)
	go longProcess(ctx)
	time.Sleep(5 * time.Second)
}
