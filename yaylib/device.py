import random
from typing import Dict, Optional, Self

DEVICES = [
    {
        "device_type": "android",
        "os_version": "11",
        "screen_density": "3.5",
        "screen_size": "1440x2960",
        "model": "Galaxy S9",
    },
]


class Device:
    def __init__(
        self,
        device_type: str,
        os_version: str,
        screen_density: str,
        screen_size: str,
        model: str,
    ):
        self.device_type = device_type
        self.os_version = os_version
        self.screen_density = screen_density
        self.screen_size = screen_size
        self.model = model

    @classmethod
    def instance(cls, device: Optional[Dict] = None) -> Self:
        if device is None:
            device = random.choice(DEVICES)
        return cls(**device)

    def generate_user_agent(self) -> str:
        return f"{self.device_type} {self.os_version} ({self.screen_density}x {self.screen_size} {self.model})"

    def generate_device_info(self, yay_version_name: str) -> str:
        return f"yay {yay_version_name} {self.generate_user_agent()}"
