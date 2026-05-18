// イベントストリームで新着チャットメッセージを受信する簡単なボット。
// Ctrl-C で終了します。
//
//	YAY_EMAIL=... YAY_PASSWORD=... go run ./examples/eventstream
package main

import (
	"context"
	"fmt"
	"log"
	"os"
	"os/signal"

	"github.com/ekkx/yaylib/v2"
)

func main() {
	ctx, stop := signal.NotifyContext(context.Background(), os.Interrupt)
	defer stop()

	client := yaylib.NewClient()
	if _, _, err := client.LoginWithEmail(ctx).
		Email(os.Getenv("YAY_EMAIL")).
		Password(os.Getenv("YAY_PASSWORD")).
		Execute(); err != nil {
		log.Fatal("login failed: ", err)
	}

	stream, err := client.OpenEventStream(ctx)
	if err != nil {
		log.Fatal("open event stream failed: ", err)
	}
	defer stream.Close()

	sub, err := stream.Subscribe(ctx, yaylib.ChatRoomChannel())
	if err != nil {
		log.Fatal("subscribe failed: ", err)
	}
	fmt.Println("listening for events (Ctrl-C to stop)")

	for {
		select {
		case <-ctx.Done():
			return
		case ev, ok := <-sub.Events():
			if !ok {
				return
			}
			switch e := ev.(type) {
			case *yaylib.NewMessageEvent:
				fmt.Println("new message:", e.GetId())
			default:
				fmt.Printf("event: %T\n", ev)
			}
		}
	}
}
