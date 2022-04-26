package main

import (
	"context"
	"fmt"
	"time"
)

func ctxWithTimeOut(ctx context.Context, t time.Duration) {
	for {
		select {
		case <-ctx.Done():
			fmt.Println("Context timeout")
			return
		}
	}
}

func ctxWithCancel(ctx context.Context) {
	for {
		select {
		case <-ctx.Done():
			fmt.Println("Context canceled")
			return
		}
	}
}

func chanel(ch chan struct{}, t time.Duration) {
	select {
	case <-time.After(t):
		close(ch)
		fmt.Println("channel closed")
	}
}
func main() {
	t := time.Second * time.Duration(1)
	ctx := context.Background()
	ctx1, cancel := context.WithTimeout(ctx, t)
	ctx2, Cancel := context.WithCancel(ctx)
	ch := make(chan struct{})

	fmt.Println("Run goroutin with timeout")
	go ctxWithTimeOut(ctx1, t)
	fmt.Println("Run goroutin with cancel")
	go ctxWithCancel(ctx2)
	fmt.Println("Run goroutin via chan")
	go chanel(ch, t)
	fmt.Println("-----------*******------------")

	select {
	case <-time.After(t):
		cancel()
		Cancel()
	}
	time.Sleep(5 * time.Millisecond)

}
