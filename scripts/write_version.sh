#!/bin/bash

VERSION=$(grep '^version' ../pyproject.toml | head -n 1 | sed 's/version *= *"\(.*\)"/\1/')

sed -i -E "s/^__version__ = \".*\"/__version__ = \"${VERSION}\"/" ../yaylib/__init__.py

echo "Updated __version__ to ${VERSION} in yaylib/__init__.py"
