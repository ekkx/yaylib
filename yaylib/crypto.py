import hashlib
import base64


from typing import Any, Optional
from cryptography.fernet import Fernet


class CryptoManager:
    def __init__(self, password: Optional[str] = None) -> None:
        self.__encryption_key: Optional[Fernet] = None

        if password is not None:
            self.__encryption_key: Fernet = self._generate_key(password)

    def generate_key(self, password: str) -> Fernet:
        hashed = hashlib.sha256(password.encode()).digest()
        key: bytes = base64.urlsafe_b64encode(hashed[:32])
        return Fernet(key)

    def hash(self, text: str) -> str:
        return hashlib.sha256(text.encode()).hexdigest()

    def encrypt(self, text: str) -> str:
        encoded: bytes = text.encode()
        encrypted: bytes = self.__encryption_key.encrypt(encoded)
        return encrypted.decode()

    def decrypt(self, text: str) -> str:
        decrypted: bytes = self.__encryption_key.decrypt(text)
        return decrypted.decode()
