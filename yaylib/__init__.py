"""

Yay! (nanameue, Inc.) API Client
~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~

An API wrapper for Yay! (yay.space) written in Python.

:copyright: (c) 2023 ekkx
:license: MIT, see LICENSE for more details.

"""

__title__ = "yaylib"
__author__ = "ekkx"
__license__ = "MIT"
__copyright__ = "Copyright (c) 2023 ekkx"
__version__ = "1.4.15"

from .client import Client
from .cookie import Cookie
from .errors import *
from .models import *
from .responses import *
from .types import *
from .utils import mention
from .ws import Intents

__all__ = (
    "Client",
    "Cookie",
    "errors",
    "models",
    "responses",
    "types",
    "mention",
    "Intents",
)
