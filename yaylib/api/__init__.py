"""

Yay! (nanameue, Inc.) API Client
~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~

An API wrapper for Yay! (yay.space) written in Python.

:copyright: (c) 2023 ekkx
:license: MIT, see LICENSE for more details.

"""

from .auth import *
from .post import *

__all__ = [
    "auth",
    "post",
]
