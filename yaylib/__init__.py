"""

Yay! (nanameue, Inc.) API Client
~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~

An API wrapper for Yay! (yay.space) written in Python.

:copyright: (c) 2023 qvco
:license: MIT, see LICENSE for more details.

"""

__title__ = "yaylib"
__author__ = "qvco"
__license__ = "MIT"
__copyright__ = "Copyright (c) 2023 qvco"
__version__ = "1.3.4"

from .client import Client
from .cookie import Cookie
from .errors import *
from .models import *
from .responses import *
from .ws import Intents

__all__ = (
    "Client",
    "Cookie",
    "errors",
    "models",
    "responses",
    "Intents",
)
