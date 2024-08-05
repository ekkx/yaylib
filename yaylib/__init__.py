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
__version__ = "1.5.0"

from .client import Client
from .state import State
from .errors import (
    YaylibError,
    HTTPError,
    BadRequestError,
    AuthenticationError,
    ForbiddenError,
    NotFoundError,
    RateLimitError,
    InternalServerError,
    ErrorCode,
)
from .models import *
from .responses import *
from .utils import mention
from .ws import Intents

__all__ = (
    "Client",
    "State",
    # errors - start
    "YaylibError",
    "HTTPError",
    "BadRequestError",
    "AuthenticationError",
    "ForbiddenError",
    "NotFoundError",
    "RateLimitError",
    "InternalServerError",
    "ErrorCode",
    # errors - end
    "models",
    "responses",
    "mention",
    "Intents",
)
