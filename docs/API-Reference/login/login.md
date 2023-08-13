## `login()`

ログインするメソッドです。戻り値として認証情報が渡されます。

認証情報はローカルストレージに保存され、次回のログイン処理を省略します。

認証情報を暗号化して保存するには、`Client`クラスを初期化する際に、`encrypt_cookie`引数を`True`に設定してください。

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

#### <a href="https://github.com/qvco/yaylib/blob/master/docs/Object-Models/LoginUserResponse.md">LoginUserResponse</a>

### Examples

- 認証情報を暗号化しない場合

```python
import yaylib

api = yaylib.Client()

email = "your_email"
password = "your_password"

api.login(email, password)

>>> INFO - Successfully logged in as '1826393'
```

- 認証情報を暗号化する場合

初回実行時:

```python
import yaylib

api = yaylib.Client(encrypt_cookie=True)

email = "your_email"
password = "your_password"

api.login(email, password)

>>> Your 'secret_key' for your_email is: wFTwqRSddPzcfs_U2D1NIxFueWwPToVxjA3woDopKWk=
```

以降は`secret_key`を引数に設定します。

```python
import yaylib

api = yaylib.Client(encrypt_cookie=True)

email = "your_email"
password = "your_password"
secret_key = "wFTwqRSddPzcfs_U2D1NIxFueWwPToVxjA3woDopKWk="

api.login(email, password, secret_key)

>>> INFO - Successfully logged in as '1826393'
```
