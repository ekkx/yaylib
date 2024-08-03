import sqlite3

from typing import Optional
from dataclasses import dataclass


@dataclass(slots=True)
class User:
    user_id: int
    email: str
    device_uuid: str
    access_token: str
    refresh_token: str


class Storage:

    instance = None

    def __new__(cls):
        if cls.instance is None:
            cls.instance = super().__new__(cls)
        return cls.instance

    def __init__(self, path: str):
        self.__con = sqlite3.connect(f"{path}.db")
        self.__cur = self.__con.cursor()
        self.__cur.execute(
            """
            CREATE TABLE IF NOT EXISTS "users" (
            "id" INTEGER NOT NULL UNIQUE,
            "email" TEXT NOT NULL,
            "device_uuid" TEXT NOT NULL,
            "access_token" TEXT NOT NULL,
            "refresh_token" TEXT NOT NULL,
            PRIMARY KEY("id"));
            """
        )
        self.__con.commit()

    def get_user(self, user_id: int) -> Optional[User]:
        res = self.__cur.execute("SELECT * FROM users WHERE id = ?", (user_id))
        result = res.fetchone()
        if result is None:
            return
        else:
            return User(
                user_id=result[0],
                email=result[1],
                device_uuid=result[2],
                access_token=result[3],
                refresh_token=result[4],
            )

    def create_user(self, user: User) -> bool:
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
        email: Optional[str],
        device_uuid: Optional[str],
        access_token: Optional[str],
        refresh_token: Optional[str],
    ) -> bool:
        try:
            sql = "UPDATE users SET "
            if email is not None:
                sql += "email = ?"
            if device_uuid is not None:
                sql += "device_uuid = ?"
            if access_token is not None:
                sql += "access_token = ?"
            if refresh_token is not None:
                sql += "refresh_token = ?"
            sql += " WHERE id = ?"
            self.__cur.execute(
                sql, (email, device_uuid, access_token, refresh_token, user_id)
            )
            self.__con.commit()
            return True
        except sqlite3.IntegrityError:
            return False

    def delete_user(self, user_id: int) -> bool:
        try:
            self.__cur("DELETE FROM users WHERE id = ? ", (user_id))
            return True
        except sqlite3.IntegrityError:
            return False
