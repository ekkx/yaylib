// テキスト投稿の例。投稿系のエンドポイントは X-Jwt を必要とします
// （client.generate_x_jwt() で取得し、そのまま渡します）。
//
//   YAY_EMAIL=... YAY_PASSWORD=... cargo run --manifest-path examples/rust/post/Cargo.toml

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

    let jwt = client.generate_x_jwt();
    let res = client
        .create_post(
            Some(&jwt),
            None, None, None, None, None, None, None, None, None, // attachment 2-9, attachment
            None, None, None, None, None, None, None, None,       // choices..mention_ids, message_tags
            Some(yaylib::models::PostType::Text),                  // post_type
            None,                                                  // shared_url
            Some("hello from yaylib"),                             // text
            None,                                                  // video_file_name
        )
        .await
        .expect("create post failed");

    println!("posted: {}", res.id.flatten().unwrap_or(0));
}
