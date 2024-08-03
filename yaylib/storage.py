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

from typing import Optional
from dataclasses import dataclass


@dataclass(slots=True)
class User:
    """ストレージ内のユーザー型"""

    user_id: int
    email: str
    device_uuid: str
    access_token: str
    refresh_token: str


class Storage:
    """`yaylib.Client` のステートを保存するローカルデータベースの操作を行う"""

    def __init__(self, path: str):
        self.path = path
        self.__con = None
        self.__cur = None

    def __enter__(self):
        self.__con = sqlite3.connect(self.path)
        self.__cur = self.__con.cursor()
        self.__cur.execute(
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
        self.__con.commit()
        return self

    def __exit__(self, exc_type, exc_val, exc_tb):
        self.__cur.close()
        self.__con.close()

    def get_user(self, user_id: int) -> Optional[User]:
        """ユーザーを取得する"""
        res = self.__cur.execute(f"SELECT * FROM users WHERE id = {user_id}")
        user = res.fetchone()

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
        try:
            self.__cur.execute(
                "INSERT INTO users (id, email, device_uuid, access_token, refresh_token) VALUES (?, ?, ?, ?, ?)",
                (
                    user.user_id,
                    user.email,
                    user.device_uuid,
                    user.access_token,
                    user.refresh_token,
                ),
            )
            self.__con.commit()
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
        try:
            sql = "UPDATE users SET "
            updates = []

            if email is not None:
                updates.append(f"email = '{email}'")
            if device_uuid is not None:
                updates.append(f"device_uuid = '{device_uuid}'")
            if access_token is not None:
                updates.append(f"access_token = '{access_token}'")
            if refresh_token is not None:
                updates.append(f"refresh_token = '{refresh_token}'")

            sql += ", ".join(updates)
            sql += f" WHERE id = {user_id};"

            self.__cur.execute(sql)
            self.__con.commit()
            return True
        except sqlite3.IntegrityError:
            return False

    def delete_user(self, user_id: int) -> bool:
        """ユーザーを削除する"""
        try:
            self.__cur.execute(f"DELETE FROM users WHERE id = {user_id} ")
            self.__con.commit()
            return True
        except sqlite3.IntegrityError:
            return False
