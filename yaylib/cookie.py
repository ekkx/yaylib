import base64
import hashlib
import json
import os
import uuid

from cryptography.fernet import Fernet
from typing import Optional

from .errors import YayError


class CookieAuthentication(object):
    __slots__ = ("access_token", "refresh_token")

    def __init__(self, data: dict) -> None:
        self.access_token = data.get("access_token")
        self.refresh_token = data.get("refresh_token")


class CookieUser(object):
    __slots__ = ("email", "user_id", "uuid")

    def __init__(self, data: dict) -> None:
        self.email: str = data.get("email")
        self.user_id: int = data.get("user_id")
        self.uuid: str = data.get("uuid")


class CookieDevice(object):
    __slots__ = "device_uuid"

    def __init__(self, data: dict) -> None:
        self.device_uuid: str = data.get("device_uuid")


class CookieProps(object):
    __slots__ = ("authentication", "user", "device")

    def __init__(self, data: dict) -> None:
        self.authentication: CookieAuthentication = CookieAuthentication(
            data.get("authentication")
        )
        self.user: CookieUser = CookieUser(data.get("user"))
        self.device: CookieDevice = CookieDevice(data.get("device"))


class Cookie(object):
    def __init__(
        self,
        save_cookie_file: bool,
        filepath: str,
        cookie_password: Optional[str] = None,
    ) -> None:
        self.__save_cookie_file: bool = save_cookie_file
        self.__filepath: str = filepath
        self.__email: str = ""
        self.__user_id: int = 0
        self.__uuid: str = str(uuid.uuid4())
        self.__device_uuid: str = str(uuid.uuid4())
        self.__access_token: str = ""
        self.__refresh_token: str = ""
        self.__encryption_key: Optional[Fernet] = None

        if cookie_password is not None:
            self.__encryption_key: Fernet = self.__generate_key(cookie_password)

    def __is_encrypted(self, cookie: CookieProps) -> bool:
        return cookie.authentication.access_token.startswith("encrypted:")

    def __generate_key(self, password: str) -> Fernet:
        key: bytes = base64.urlsafe_b64encode(password.ljust(32))
        return Fernet(key)

    def __hash(self, text: str) -> str:
        return hashlib.sha256(text.encode()).hexdigest()

    def __encrypt(self, text: str) -> str:
        encoded: bytes = text.encode()
        encrypted: bytes = self.__encryption_key.encrypt(encoded)
        return "encrypted:" + encrypted.decode()

    def __decrypt(self, text: str) -> str:
        if text.startswith("encrypted:"):
            text = text[len("encrypted:") :]
        decrypted: bytes = self.__encryption_key.decrypt(text)
        return decrypted.decode()

    def __encrypt_cookie(self, cookie: CookieProps) -> CookieProps:
        return CookieProps(
            {
                **cookie,
                "user": {
                    **cookie.user,
                    "uuid": self.__encrypt(cookie.user.uuid),
                },
                "device": {
                    **cookie.device,
                    "device_uuid": self.__encrypt(cookie.device.device_uuid),
                },
                "authentication": {
                    **cookie.authentication,
                    "access_token": self.__encrypt(cookie.authentication.access_token),
                    "refresh_token": self.__encrypt(
                        cookie.authentication.refresh_token
                    ),
                },
            }
        )

    def __decrypt_cookie(self, cookie: CookieProps) -> CookieProps:
        return CookieProps(
            {
                **cookie,
                "user": {
                    **cookie.user,
                    "uuid": self.__decrypt(cookie.user.uuid),
                },
                "device": {
                    **cookie.device,
                    "device_uuid": self.__decrypt(cookie.device.device_uuid),
                },
                "authentication": {
                    **cookie.authentication,
                    "access_token": self.__decrypt(cookie.authentication.access_token),
                    "refresh_token": self.__decrypt(
                        cookie.authentication.refresh_token
                    ),
                },
            }
        )

    def get(self) -> CookieProps:
        return CookieProps(
            {
                "authentication": {
                    "access_token": self.__access_token,
                    "refresh_token": self.__refresh_token,
                },
                "user": {
                    "email": self.__email,
                    "user_id": self.__user_id,
                    "uuid": self.__uuid,
                },
                "device": {"device_uuid": self.__device_uuid},
            }
        )

    def set(self, cookie: Optional[CookieProps] = None) -> None:
        self.__email = cookie.user.email
        self.__user_id = cookie.user.user_id
        self.__uuid = cookie.user.uuid
        self.__device_uuid = cookie.device.device_uuid
        self.__access_token = cookie.authentication.access_token
        self.__refresh_token = cookie.authentication.refresh_token

    def save(self, cookie: Optional[CookieProps] = None) -> None:
        if not self.__save_cookie_file:
            return

        if cookie is None:
            cookie = self.get()
            if self.__encryption_key:
                cookie = self.__encrypt_cookie(cookie)

        cookie.user.email = self.__hash(cookie.user.email)

        with open(self.__filepath, "w") as f:
            json.dump(cookie, f, indent=4)

    def load(self, email: str) -> CookieProps:
        with open(self.__filepath, "r") as f:
            loaded_cookie: CookieProps = json.load(f)

        if self.__hash(email) != loaded_cookie.user.email:
            raise YayError("クッキーのメールアドレスが一致しませんでした。")

        loaded_cookie.user.email = email

        is_encrypted: bool = self.__is_encrypted(loaded_cookie)

        if is_encrypted and not self.__encryption_key:
            raise YayError("このクッキーは暗号化されています。")

        if self.__encryption_key:
            if not is_encrypted:
                loaded_cookie = self.__encrypt_cookie(loaded_cookie)
                self.save(loaded_cookie)

            loaded_cookie = self.__decrypt_cookie(loaded_cookie)

        self.set(loaded_cookie)

        return self.get()

    def destroy(self) -> None:
        os.remove(self.__filepath)
