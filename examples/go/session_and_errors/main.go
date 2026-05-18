// セッション永続化とエラー処理の例。
//
// WithSessionStore を渡すと、ログイン状態がファイルに保存され、次回以降は
// キャッシュされたセッションが再利用されます（毎回ログインしない）。
// 非 2xx は *yaylib.APIError として返り、CodeOf / ErrorResponseOf で
// 中身を判定できます。
//
//	YAY_EMAIL=... YAY_PASSWORD=... go run ./examples/go/session_and_errors
package main

import (
	"context"
	"errors"
	"fmt"
	"log"
	"os"

	"github.com/ekkx/yaylib/v2"
)

func main() {
	ctx := context.Background()

	store, err := yaylib.NewSessionStore("yay-session.json")
	if err != nil {
		log.Fatal("session store: ", err)
	}
	client := yaylib.NewClient(yaylib.WithSessionStore(store))

	// 2 回目以降の実行では保存済みセッションが使われます。
	if _, _, err := client.LoginWithEmail(ctx).
		Email(os.Getenv("YAY_EMAIL")).
		Password(os.Getenv("YAY_PASSWORD")).
		Execute(); err != nil {
		log.Fatal("login failed: ", err)
	}

	_, _, err = client.GetTimeline(ctx, yaylib.NoreplyModeFalse).Number(20).Execute()
	if err != nil {
		// grpc 的なコード分岐
		fmt.Println("error code:", yaylib.CodeOf(err))
		// サーバのエラー封筒（ban_until / retry_in など）
		if r := yaylib.ErrorResponseOf(err); r != nil {
			fmt.Println("message:", r.Message)
		}
		// 生のステータス / ボディが要るとき
		var apiErr *yaylib.APIError
		if errors.As(err, &apiErr) {
			fmt.Println("raw body:", string(apiErr.Body()))
		}
		return
	}
	fmt.Println("ok")
}
