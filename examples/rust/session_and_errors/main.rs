// セッション永続化とエラー処理の例。
//
// FileSessionStore を渡すと、ログイン状態がファイルに保存され、次回以降は
// キャッシュされたセッションが再利用されます（毎回ログインしない）。
// 非 2xx は ApiError として返り、code_of / error_response_of で
// 中身を判定できます。
//
//   YAY_EMAIL=... YAY_PASSWORD=... cargo run --manifest-path examples/rust/session_and_errors/Cargo.toml

use yaylib::models::NoreplyMode;

#[tokio::main]
async fn main() {
    let store = yaylib::FileSessionStore::new("yay-session.json")
        .await
        .expect("session store");

    let client = yaylib::Client::builder().session_store(store).build();

    // 2 回目以降の実行では保存済みセッションが使われます。
    client
        .login_with_email()
        .email(std::env::var("YAY_EMAIL").unwrap())
        .password(std::env::var("YAY_PASSWORD").unwrap())
        .execute()
        .await
        .expect("login failed");

    let result = client
        .get_timeline(NoreplyMode::False, None, None, None, None, Some(20), None, None, None, None, None)
        .await;

    match result {
        Ok(_) => println!("ok"),
        Err(ref e) => {
            // grpc 的なコード分岐
            println!("error code: {:?}", yaylib::code_of(e));
            // サーバのエラー封筒（ban_until / retry_in など）
            if let Some(r) = yaylib::error_response_of(e) {
                println!("message: {}", r.message);
            }
            // 生のステータス / ボディが要るとき
            if let yaylib::ApiError::Response { status, body, .. } = e {
                println!("status: {status}, raw body: {}", String::from_utf8_lossy(body));
            }
        }
    }
}
