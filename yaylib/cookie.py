import os
import hashlib
import uuid
import json


class Cookie(object):
    def __init__(
        self, save_cookie_file: bool, filepath: str, cookie_password: str
    ) -> None:
        pass

    @property
    def cookie(self) -> dict:
        pass

    @property
    def user_id(self) -> int:
        pass

    @property
    def uuid(self) -> str:
        pass

    @property
    def device_uuid(self) -> str:
        pass

    @property
    def access_token(self) -> str:
        pass

    @property
    def refresh_token(self) -> str:
        pass

    def __is_encrypted(self, cookie: dict) -> bool:
        pass

    def __generate_key(self, password: str):
        pass

    def __hash(self, text: str) -> str:
        pass

    def __encrypt(self, text: str) -> str:
        pass

    def __decrypt(self, text: str) -> str:
        pass

    def __encrypt_cookie(self, cookie: dict) -> dict:
        pass

    def __decrypt_cookie(self, cookie: dict) -> dict:
        pass

    def set(self, cookie: dict) -> None:
        pass

    def save(self, cookie: dict | None = None) -> None:
        pass

    def load(self, email: str) -> dict:
        pass

    def destroy(self) -> None:
        pass
