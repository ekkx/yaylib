// 認証してタイムラインを取得する最小例。
//
//   YAY_EMAIL=... YAY_PASSWORD=... cargo run --manifest-path examples/rust/timeline/Cargo.toml

use yaylib::models::NoreplyMode;

#[tokio::main]
async fn main() {
    let client = yaylib::Client::builder().build();

    client
        .login_with_email()
        .email(std::env::var("YAY_EMAIL").unwrap())
        .password(std::env::var("YAY_PASSWORD").unwrap())
        .execute()
        .await
        .expect("login failed");

    let tl = client
        .get_timeline(NoreplyMode::False, None, None, None, None, Some(20), None, None, None, None, None)
        .await
        .expect("timeline failed");

    for p in tl.posts.flatten().unwrap_or_default() {
        println!(
            "{} {}",
            p.id.flatten().unwrap_or(0),
            p.text.flatten().unwrap_or_default(),
        );
    }
}
