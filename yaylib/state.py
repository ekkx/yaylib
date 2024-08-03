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

from typing import Optional

from .crypto import Crypto
from .storage import Storage, User


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

    def exists(self, email: str) -> bool:
        return bool(self.get_user(email=self.__crypto.hash(email)))

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
