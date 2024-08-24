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
from .constants import *
from .errors import *
from .models import *
from .responses import *
from .state import State
from .utils import mention

__all__ = (
    "Client",
    "constants",
    "errors",
    "models",
    "responses",
    "State",
    "mention",
)
