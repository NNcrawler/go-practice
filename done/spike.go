package done

import (
	"context"
	"fmt"
	"time"
)

func Run() {
	ctx, cancel:= context.WithCancel(context.Background())
	go func (ctx context.Context) {
		<- ctx.Done()
		fmt.Println("triggered 1")
	}(ctx)

	cancel()
	<- ctx.Done()
	fmt.Println("Gracefull Period")
	<- time.After(3 * time.Second)
}
