"""

Yay! (nanameue, Inc.) API Client
~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~

A powerful API wrapper for Yay! (yay.space) written in Python.

:copyright: (c) 2023 ekkx
:license: MIT, see LICENSE for more details.

https://github.com/ekkx/yaylib

"""

__title__ = "yaylib"
__author__ = "ekkx"
__license__ = "MIT"
__copyright__ = "Copyright (c) 2023 ekkx"

# 1.0.0.dev1  Development release
# 1.0.0a1     Alpha Release
# 1.0.0b1     Beta Release
# 1.0.0rc1    Release Candidate
# 1.0.0       Final Release
# 1.0.0.post1 Post Release
__version__ = "1.5.2"

from .client import Client
from .constants import *
from .errors import *
from .models import *
from .responses import *
from .state import State
from .utils import mention
from .ws import *

__all__ = (
    "Client",
    "constants",
    "errors",
    "models",
    "responses",
    "State",
    "mention",
    "ws",
)
