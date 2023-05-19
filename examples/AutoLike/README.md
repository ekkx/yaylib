# AutoLike

タイムラインの投稿を自動で「いいね」します。

### 実行方法

以下のスクリプトをコンソール上で実行してください。

```bash
python auto_like.py
```

### 設定

- メールアドレス、パスワードは _.env_ ファイルに設定してください。
- 「いいね」の回数を指定する場合は、start()関数の引数に整数を入力してください。  
  ※ 指定しない場合は、永久的に実行し続けます。

```python
bot.start(30) # 30回「いいね」する場合
```

<!-- - 返信や通話の投稿に対して「いいね」しない場合は 32 行目を以下のように書き直してください。

```python
if post.liked is False and post.reply_to_id is None and post.type != 'call':
``` -->

---

_by @qualia-5w4_
