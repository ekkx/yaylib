"""
MIT License

Copyright (c) 2023 ekkx

Permission is hereby granted, free of charge, to any person obtaining a copy
of this software and associated documentation files (the "Software"), to deal
in the Software without restriction, including without limitation the rights
to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
copies of the Software, and to permit persons to whom the Software is
furnished to do so, subject to the following conditions:

The above copyright notice and this permission notice shall be included in all
copies or substantial portions of the Software.

THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE
SOFTWARE.
"""

import hashlib
import base64


from typing import Optional
from cryptography.fernet import Fernet


class Crypto:
    """暗号化を行うクラス"""

    def __init__(self, password: Optional[str] = None) -> None:
        self.__encryption_key: Optional[Fernet] = None

        if password is not None:
            self.__encryption_key = self.__generate_key(password)

    @staticmethod
    def __generate_key(password: str) -> Fernet:
        hashed = hashlib.sha256(password.encode()).digest()
        key = base64.urlsafe_b64encode(hashed[:32])
        return Fernet(key)

    @staticmethod
    def hash(text: str) -> str:
        """`sha256` を用いてハッシュ化を行います"""
        return hashlib.sha256(text.encode()).hexdigest()

    def encrypt(self, text: str) -> str:
        """設定されたパスワードを元に文字列を暗号化します

        Raises:
            AttributeError: パスワードが設定されていない場合
        """
        if self.__encryption_key is None:
            raise AttributeError("Encryption key is not defined.")
        encoded = text.encode()
        encrypted = self.__encryption_key.encrypt(encoded)
        return encrypted.decode()

    def decrypt(self, text: str) -> str:
        """設定されたパスワードを元に文字列を復号化します

        Raises:
            AttributeError: パスワードが設定されていない場合
        """
        if self.__encryption_key is None:
            raise AttributeError("Encryption key is not defined.")
        decrypted = self.__encryption_key.decrypt(text)
        return decrypted.decode()
