// テキスト投稿の例。投稿系のエンドポイントは X-Jwt を必要とします
// （client.GenerateXJwt() で取得し、そのまま渡します）。
//
//	YAY_EMAIL=... YAY_PASSWORD=... go run ./examples/post
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

	res, _, err := client.CreatePost(ctx).
		XJwt(client.GenerateXJwt()).
		PostType("text").
		Text("hello from yaylib").
		Execute()
	if err != nil {
		log.Fatal("create post failed: ", err)
	}
	fmt.Println("posted:", res.GetId())
}
