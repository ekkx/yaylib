## `change_email()`

メールアドレスを変更します。

### Endpoint URL

#### `PUT`: `https://api.yay.space/v1/users/change_email`

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
		<td><code>email_grant_token</code>(optional)</td>
		<td><code>str</code></td>
		<td>認証用トークン</td>
	</tr>
</table>

### Returns

#### <a href="">LoginUpdateResponse<a>

### Examples

```python
>>> email = "メールアドレス"
>>> password = "パスワード"
>>> api.change_email(email=email, password=password)
```
