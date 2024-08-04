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

import sqlite3

from queue import Queue
from typing import Optional
from dataclasses import dataclass

from . import utils
from .crypto import Crypto


@dataclass(slots=True)
class User:
    """ストレージ内のユーザー型"""

    user_id: int
    email: str
    device_uuid: str
    access_token: str
    refresh_token: str


class SQLiteConnectionPool:
    """`sqlite3` のコネクションマネージャー"""

    def __init__(self, db_path, pool_size=5):
        self.db_path = db_path
        self.pool = Queue(maxsize=pool_size)
        for _ in range(pool_size):
            self.pool.put(sqlite3.connect(db_path))

    def get_connection(self) -> sqlite3.Connection:
        """コネクションを取得する"""
        return self.pool.get()

    def return_connection(self, conn) -> None:
        """コネクションを返却する"""
        self.pool.put(conn)


class Storage:
    """`yaylib.Client` のステートを保存するローカルデータベースの操作を行う"""

    def __init__(self, path: str, pool_size=5):
        self.__pool = SQLiteConnectionPool(path, pool_size)

        with self.__pool.get_connection() as conn:
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

    def get_user(
        self, user_id: Optional[int] = None, email: Optional[str] = None
    ) -> Optional[User]:
        """ユーザーを取得する"""
        with self.__pool.get_connection() as conn:
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

        if user is None:
            return None

        return User(
            user_id=user[0],
            email=user[1],
            device_uuid=user[2],
            access_token=user[3],
            refresh_token=user[4],
        )

    def create_user(self, user: User) -> bool:
        """ユーザーを作成する"""
        with self.__pool.get_connection() as conn:
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

    def update_user(
        self,
        user_id: int,
        email: Optional[str] = None,
        device_uuid: Optional[str] = None,
        access_token: Optional[str] = None,
        refresh_token: Optional[str] = None,
    ) -> bool:
        """ユーザーを更新する"""
        with self.__pool.get_connection() as conn:
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

    def delete_user(self, user_id: int) -> bool:
        """ユーザーを削除する"""
        with self.__pool.get_connection() as conn:
            try:
                cursor = conn.cursor()
                cursor.execute("DELETE FROM users WHERE id = ?", (user_id,))
                conn.commit()
                return True
            except sqlite3.IntegrityError:
                return False


class State(Storage):
    """揮発性メモリ上にステートを保持するクラス"""

    user_id = 0
    email = ""
    device_uuid = ""
    access_token = ""
    refresh_token = ""

    def __init__(self, path: str, password: Optional[str] = None, pool_size=5):
        super().__init__(path, pool_size)
        self.__crypto = Crypto(password)
        self.device_uuid = utils.generate_uuid(True)

    def get_user_by_email(self, email: str) -> Optional[User]:
        return self.get_user(email=self.__crypto.hash(email))

    def create(self) -> bool:
        return self.create_user(
            User(
                self.user_id,
                email=self.__crypto.hash(self.email),
                device_uuid=self.__crypto.encrypt(self.device_uuid),
                access_token=self.__crypto.encrypt(self.access_token),
                refresh_token=self.__crypto.encrypt(self.refresh_token),
            )
        )

    def update(self) -> bool:
        return self.update_user(
            self.user_id,
            email=self.__crypto.hash(self.email),
            device_uuid=self.__crypto.encrypt(self.device_uuid),
            access_token=self.__crypto.encrypt(self.access_token),
            refresh_token=self.__crypto.encrypt(self.refresh_token),
        )

    def destory(self) -> bool:
        return self.delete_user(self.user_id)
