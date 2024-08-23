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

import base64
import hashlib
import sqlite3
from dataclasses import dataclass
from queue import Queue
from typing import Optional

from cryptography.fernet import Fernet

from . import utils


@dataclass(slots=True)
class LocalUser:
    """ストレージ内のユーザー型"""

    user_id: int
    email: str
    device_uuid: str
    access_token: str
    refresh_token: str


class Crypto:
    """暗号化を行うクラス"""

    def __init__(self, password: Optional[str] = None) -> None:
        self.__encryption_key: Optional[Fernet] = None
        if password is not None:
            self.__encryption_key = self.generate_key(password)

    @staticmethod
    def generate_key(password: str) -> Fernet:
        """鍵を生成する

        Args:
            password (str):

        Returns:
            Fernet: 鍵
        """
        hashed = hashlib.sha256(password.encode()).digest()
        key = base64.urlsafe_b64encode(hashed[:32])
        return Fernet(key)

    @staticmethod
    def hash(text: str) -> str:
        """`sha256` を用いてハッシュ化を行う

        Returns:
            str:
        """
        return hashlib.sha256(text.encode()).hexdigest()

    def set_encryption_key(self, key: Fernet) -> None:
        """鍵を設定する

        Args:
            key (Fernet): 鍵
        """
        self.__encryption_key = key

    def has_encryption_key(self) -> bool:
        """鍵が設定されているか確認する

        Returns:
            bool: 鍵が設定されている場合は True, そうでない場合 False
        """
        return self.__encryption_key is not None

    def encrypt(self, text: str) -> str:
        """設定された鍵から文字列を暗号化する

        Note:
            鍵が設定されていない場合は、引数で受け取った文字列をそのまま返す

        Args:
            text (str):

        Returns:
            str: 暗号化された文字列
        """
        if self.__encryption_key is None:
            return text
        encoded = text.encode()
        encrypted = self.__encryption_key.encrypt(encoded)
        return encrypted.decode()

    def decrypt(self, text: str) -> str:
        """設定された鍵から文字列を復号化する

        Note:
            鍵が設定されていない場合は、引数で受け取った文字列をそのまま返す

        Args:
            text (str):

        Returns:
            str: 複合化された文字列
        """
        if self.__encryption_key is None:
            return text
        decrypted = self.__encryption_key.decrypt(text)
        return decrypted.decode()


class SQLiteConnectionPool:
    """`sqlite3` のコネクションマネージャー"""

    def __init__(self, db_path, pool_size=5):
        self.__pool = Queue(maxsize=pool_size)
        for _ in range(pool_size):
            self.__pool.put(sqlite3.connect(db_path))

    def get_connection(self) -> sqlite3.Connection:
        """コネクションを取得する"""
        return self.__pool.get()

    def return_connection(self, conn: sqlite3.Connection) -> None:
        """コネクションを返却する"""
        self.__pool.put(conn)


class Storage:
    """クライアントのステートのデータベース操作を行う"""

    def __init__(self, path: str, pool_size=5):
        self.__pool = SQLiteConnectionPool(path, pool_size)

        conn = self.__pool.get_connection()
        try:
            cursor = conn.cursor()
            cursor.execute(
                """
                CREATE TABLE IF NOT EXISTS users (
                    id INTEGER PRIMARY KEY,
                    email TEXT NOT NULL,
                    device_uuid TEXT NOT NULL,
                    access_token TEXT NOT NULL,
                    refresh_token TEXT NOT NULL
                );
                """
            )
            conn.commit()
        finally:
            self.__pool.return_connection(conn)

    def get_user(
        self, user_id: Optional[int] = None, email: Optional[str] = None
    ) -> Optional[LocalUser]:
        """ユーザーを取得する"""
        conn = self.__pool.get_connection()
        try:
            cursor = conn.cursor()

            where = ""
            params = []
            if user_id is not None:
                where = "id = ?"
                params.append(user_id)
            elif email is not None:
                where = "email = ?"
                params.append(email)

            cursor.execute(f"SELECT * FROM users WHERE {where}", params)
            user = cursor.fetchone()
        finally:
            self.__pool.return_connection(conn)

        if user is None:
            return None

        return LocalUser(
            user_id=user[0],
            email=user[1],
            device_uuid=user[2],
            access_token=user[3],
            refresh_token=user[4],
        )

    def create_user(self, user: LocalUser) -> bool:
        """ユーザーを作成する"""
        conn = self.__pool.get_connection()
        try:
            cursor = conn.cursor()
            cursor.execute(
                "INSERT INTO users (id, email, device_uuid, access_token, refresh_token) VALUES (?, ?, ?, ?, ?)",
                (
                    user.user_id,
                    user.email,
                    user.device_uuid,
                    user.access_token,
                    user.refresh_token,
                ),
            )
            conn.commit()
            return True
        except sqlite3.IntegrityError:
            return False
        finally:
            self.__pool.return_connection(conn)

    def update_user(
        self,
        user_id: int,
        *,
        email: Optional[str] = None,
        device_uuid: Optional[str] = None,
        access_token: Optional[str] = None,
        refresh_token: Optional[str] = None,
    ) -> bool:
        """ユーザーを更新する"""
        conn = self.__pool.get_connection()
        try:
            cursor = conn.cursor()
            updates = []

            if email is not None:
                updates.append("email = ?")
            if device_uuid is not None:
                updates.append("device_uuid = ?")
            if access_token is not None:
                updates.append("access_token = ?")
            if refresh_token is not None:
                updates.append("refresh_token = ?")

            sql = f"UPDATE users SET {', '.join(updates)} WHERE id = ?"
            params = [
                param
                for param in [email, device_uuid, access_token, refresh_token]
                if param is not None
            ]
            params.append(user_id)

            cursor.execute(sql, params)
            conn.commit()
            return True
        except sqlite3.IntegrityError:
            return False
        finally:
            self.__pool.return_connection(conn)

    def delete_user(self, user_id: int) -> bool:
        """ユーザーを削除する"""
        conn = self.__pool.get_connection()
        try:
            cursor = conn.cursor()
            cursor.execute("DELETE FROM users WHERE id = ?", (user_id,))
            conn.commit()
            return True
        except sqlite3.IntegrityError:
            return False
        finally:
            self.__pool.return_connection(conn)


class State(Storage):
    """単一クライアントのステートを管理するクラス"""

    def __init__(
        self,
        *,
        storage_path: str,
        storage_pool_size=5,
        password: Optional[str] = None,
    ):
        super().__init__(storage_path, storage_pool_size)

        self.user_id = 0
        self.email = ""
        self.device_uuid = utils.generate_uuid(True)
        self.access_token = ""
        self.refresh_token = ""

        self.__crypto = Crypto(password)

    def set_user(self, user: LocalUser) -> None:
        """ユーザーを設定する

        Args:
            user (LocalUser):
        """
        self.user_id = user.user_id
        self.email = user.email
        self.access_token = user.access_token
        self.refresh_token = user.refresh_token
        self.device_uuid = user.device_uuid

    def get_user_by_email(self, email: str) -> Optional[LocalUser]:
        """メールアドレスからユーザーを取得する

        Args:
            email (str): ハッシュ化されていないメールアドレス

        Returns:
            Optional[LocalUser]: ユーザーが存在しない場合は None を返す
        """
        user = self.get_user(email=self.__crypto.hash(email))
        if user is None:
            return None
        user.email = email
        return user

    def set_encryption_key(self, password: str):
        """ローカルストレージ内のユーザーを暗号化するためのパスワードを設定する

        Args:
            password (str): 任意のパスワード
        """
        key = self.__crypto.generate_key(password)
        self.__crypto.set_encryption_key(key)

    def has_encryption_key(self) -> bool:
        """鍵が設定されているか確認する

        Returns:
            bool: 鍵が設定されている場合は True, そうでない場合 False
        """
        return self.__crypto.has_encryption_key()

    def decrypt(self, user: LocalUser) -> LocalUser:
        """設定された鍵からユーザー情報を復号化する

        Args:
            user (LocalUser): 暗号化されたユーザー

        Returns:
            LocalUser: 復号化されたユーザー
        """
        user.device_uuid = self.__crypto.decrypt(user.device_uuid)
        user.access_token = self.__crypto.decrypt(user.access_token)
        user.refresh_token = self.__crypto.decrypt(user.refresh_token)
        return user

    def save(self) -> bool:
        """設定されたユーザーをデータベースに保存する

        Note:
            自動的にメールアドレスのハッシュ化、その他認証情報を暗号化して保存する

        Returns:
            bool:
        """
        return self.create_user(
            LocalUser(
                self.user_id,
                email=self.__crypto.hash(self.email),
                device_uuid=self.__crypto.encrypt(self.device_uuid),
                access_token=self.__crypto.encrypt(self.access_token),
                refresh_token=self.__crypto.encrypt(self.refresh_token),
            )
        )

    def update(self) -> bool:
        """設定されたユーザー情報を元にデータベースをアップデートする

        Note:
            自動的にメールアドレスのハッシュ化、その他認証情報を暗号化して保存する

        Returns:
            bool:
        """
        return self.update_user(
            self.user_id,
            email=self.__crypto.hash(self.email),
            device_uuid=self.__crypto.encrypt(self.device_uuid),
            access_token=self.__crypto.encrypt(self.access_token),
            refresh_token=self.__crypto.encrypt(self.refresh_token),
        )

    def destory(self, user_id: int) -> bool:
        """データベース内のテーブルからユーザーを削除する

        Returns:
            bool:
        """
        return self.delete_user(user_id)
