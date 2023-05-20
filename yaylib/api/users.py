from ..api import (Yay, AuthenticationError)
from ..config import (Endpoints, Configs)
from typing import (Dict)


async def contact_friends(self: Yay, header: Dict[str, str | int]):
    if header.get("Authorization") is None:
        AuthenticationError("Authorization is not present in the header.")
    response = await Yay.send_request(self=self, method="get", endpoint=f"https://{Endpoints.USER_V1}/contact_friends", headers=header)
