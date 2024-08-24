"""

Yay! (nanameue, Inc.) API Client
~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~

A powerful API wrapper for Yay! (yay.space) written in Python.

:copyright: (c) 2023 ekkx
:license: MIT, see LICENSE for more details.

https://github.com/ekkx/yaylib

"""

from .auth import *
from .call import *
from .chat import *
from .group import *
from .misc import *
from .notification import *
from .post import *
from .review import *
from .thread import *
from .user import *

__all__ = [
    "auth",
    "call",
    "chat",
    "group",
    "misc",
    "notification",
    "post",
    "review",
    "thread",
    "user",
]
