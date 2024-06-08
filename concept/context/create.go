package context

import (
	"context"
	"fmt"
	"time"
)

type User struct {
	UserId int
}

func RunCreateContext() {

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	ctx = context.WithValue(ctx, "UserId", 123)
	defer cancel()

	go PerformTask(ctx)

	select {
	case <-ctx.Done():
		{
			fmt.Println("timeout")
		}
	}

}

func PerformTask(ctx context.Context) {

	select {
	case <-time.After(1 * time.Second):
		{
			fmt.Println("task completed successfully", ctx.Value("UserId"))
		}
	}
}
