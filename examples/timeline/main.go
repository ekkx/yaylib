// 認証してタイムラインを取得する最小例。
//
//	YAY_EMAIL=... YAY_PASSWORD=... go run ./examples/timeline
package main

import (
	"context"
	"fmt"
	"log"
	"os"

	"github.com/ekkx/yaylib/v2"
)

func main() {
	ctx := context.Background()
	client := yaylib.NewClient()

	if _, _, err := client.LoginWithEmail(ctx).
		Email(os.Getenv("YAY_EMAIL")).
		Password(os.Getenv("YAY_PASSWORD")).
		Execute(); err != nil {
		log.Fatal("login failed: ", err)
	}

	tl, _, err := client.GetTimeline(ctx, yaylib.NoreplyModeFalse).Number(20).Execute()
	if err != nil {
		log.Fatal("timeline failed: ", err)
	}
	for _, p := range tl.GetPosts() {
		fmt.Println(p.GetId(), p.GetText())
	}
}
