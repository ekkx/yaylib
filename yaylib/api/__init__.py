"""

Yay! (nanameue, Inc.) API Client
~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~

An API wrapper for Yay! (yay.space) written in Python.

:copyright: (c) 2023-present Qvco, Konn
:license: MIT, see LICENSE for more details.

"""

from .api import *
from .call import *
from .cassandra import *
from .chat import *
from .group import *
from .login import *
from .misc import *
from .post import *
from .review import *
from .thread import *
from .user import *
from .websocket import *

__all__ = [
    "API",
    "block",
    "call",
    "cassandra",
    "chat",
    "common",
    "config",
    "dev",
    "file_upload",
    "game",
    "gift",
    "group",
    "hidden",
    "id_card_check",
    "internal_id_card_check",
    "login",
    "misc",
    "mute_keyword",
    "payment",
    "post",
    "review",
    "search",
    "thread",
    "twitter",
    "user_identity",
    "user",
    "wallet",
    "websocket",
]
