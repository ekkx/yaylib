// イベントストリームで新着チャットメッセージを受信する簡単なボット。
// Ctrl-C で終了します。
//
//   YAY_EMAIL=... YAY_PASSWORD=... cargo run --manifest-path examples/rust/event_stream/Cargo.toml

use yaylib::{chat_room_channel, Event};

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

    let stream = client
        .open_event_stream()
        .await
        .expect("open event stream failed");

    let mut sub = stream
        .subscribe(chat_room_channel())
        .await
        .expect("subscribe failed");

    println!("listening for events (Ctrl-C to stop)");

    loop {
        tokio::select! {
            event = sub.recv() => {
                match event {
                    Some(Event::NewMessage(msg)) => {
                        let id = msg.message.id.flatten().unwrap_or(0);
                        println!("new message: {id}");
                    }
                    Some(ev) => println!("event: {ev:?}"),
                    None => break,
                }
            }
            _ = tokio::signal::ctrl_c() => break,
        }
    }
}
