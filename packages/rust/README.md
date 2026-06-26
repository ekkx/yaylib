<div align="center">

<img src="https://github.com/ekkx/yaylib/assets/77382767/45c45b21-d812-4cad-8f27-315ffef53201" alt="yaylib" height="160px">

### yaylib

**好きでつながるバーチャルワールド - Yay! の API ライブラリ**

<a href="https://github.com/ekkx/yaylib/tree/master/packages/rust"><img src="https://img.shields.io/badge/Rust-000000?style=flat-square&logo=rust&logoColor=white" alt="Rust"></a>

<sub>あらゆる操作の自動化や、ボットの開発が可能です。</sub>

<a href="https://discord.gg/MEuBfNtqRN">Discord に参加</a>

</div>

Yay! の API を Rust から扱うための非公式 SDK です。

### インストール

```sh
cargo add yaylib
```

Rust 1.88 以上。非同期ランタイムは [tokio](https://tokio.rs/) を使用します。

### クイックスタート

```rust
use yaylib::models::NoreplyMode;

#[tokio::main]
async fn main() -> Result<(), Box<dyn std::error::Error>> {
    let client = yaylib::Client::builder().build();

    // ログイン（セッションは透過的にキャッシュされます）
    client.login_with_email()
        .email("...")
        .password("...")
        .execute()
        .await?;

    // タイムラインを取得
    let tl = client.get_timeline(
        NoreplyMode::False,
        None, None, None, None, Some(20),
        None, None, None, None, None,
    ).await?;
    for p in tl.posts.flatten().unwrap_or_default() {
        println!("{} {}", p.id.flatten().unwrap_or(0), p.text.flatten().unwrap_or_default());
    }

    // 投稿する
    let jwt = client.generate_x_jwt();
    client.create_post(
        Some(&jwt), None, None, None, None, None, None, None, None, None,
        None, None, None, None, None, None, None, None,
        Some(yaylib::models::PostType::Text),
        None, Some("hello from yaylib"), None,
    ).await?;

    Ok(())
}
```

すべてのオペレーションは `client.<operation>` として直接呼び出せます。

### イベントストリーム

```rust
let stream = client.open_event_stream().await?;
let mut sub = stream.subscribe(yaylib::chat_room_channel()).await.unwrap();

// 新着メッセージを受信
while let Some(event) = sub.recv().await {
    match event {
        yaylib::Event::NewMessage(msg) => {
            println!("new message: {:?}", msg.message.id);
        }
        _ => {}
    }
}
```

### サンプル

実行できるサンプルを [`examples/`](https://github.com/ekkx/yaylib/tree/master/examples/rust) に用意しています（全言語共通）。

- 認証 + タイムライン取得
- テキスト投稿
- イベントストリームの簡単なボット
- セッション永続化とエラー処理

```bash
YAY_EMAIL=... YAY_PASSWORD=... cargo run --manifest-path examples/rust/timeline/Cargo.toml
```

### ⚖️ ライセンス

<p align="center">
  <a href="https://github.com/ekkx">
    <img src="https://github.com/ekkx/yaylib/assets/77382767/5d6aef18-5d98-4c9b-9f54-791308b393af" width="256" height="256">
  </a>
</p>

<p align="center">
  <strong>MIT © <a href="https://github.com/ekkx">ekkx</a></strong>
</p>
