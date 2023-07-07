## `login()`

ログインするメソッドです。戻り値として認証情報が渡されます。

認証情報はローカルに暗号化された状態で保存され、次回のログイン処理を省略します。

認証情報を復号化するには、初回ログイン時に生成される`secret_key`を引数に設定します。

### Endpoint URL

#### `POST`: `https://api.yay.space/v3/users/login_with_email`

### Parameters

<table>
    <tr>
        <th>Name</th>
        <th>Type</th>
        <th>Description</th>
    </tr>
    <tr>
		<td><code>email</code></td>
		<td><code>str</code></td>
		<td>メールアドレス</td>
	</tr>
    <tr>
		<td><code>password</code></td>
		<td><code>str</code></td>
		<td>パスワード</td>
	</tr>
    <tr>
		<td><code>secret_key</code>(optional)</td>
		<td><code>str</code></td>
		<td>認証情報を復号化するための鍵</td>
	</tr>
</table>

### Returns

#### <a href="https://github.com/qvco/yaylib/blob/main/docs/Object-Models/LoginUserResponse.md">LoginUserResponse</a>

### Examples

初回実行時:

```python
email = "hello@example.com"
password = "password"

api.login(email, password)

>>> Your 'secret_key' for hello@example.com is: wFTwqRSddPzcfs_U2D1NIxFueWwPToVxjA3woDopKWk=
```

それ以降は`secret_key`を引数に設定します。

```python
email = "hello@example.com"
password = "password"
secret_key = "wFTwqRSddPzcfs_U2D1NIxFueWwPToVxjA3woDopKWk="

api.login(email, password, secret_key)

>>> INFO - Successfully logged in as '1826393'
```
