from setuptools import setup, find_packages
from yaylib import __version__

name = "yaylib"
author = "Qvco, Konn"
author_email = "nikola.desuga@gmail.com"
description = "This Python package provides an easy-to-use interface for accessing data from Yay! (https://yay.space/). With this API Client, you can retrieve user profiles, posts, comments, and other content from Yay!, as well as perform common tasks like liking and commenting on posts."
long_description_content_type = "text/markdown"
license = "MIT"
url = "https://github.com/qvco/yaylib"

keywords = [
    "yay",
    "yaylib",
    "api",
    "bot",
    "client",
    "library",
    "wrapper",
    "ボット",
    "ライブラリ",
]

install_requires = [
    "httpx>=0.17.1",
    "Pillow>=9.3.0",
    "cryptography>=41.0.1",
    "websocket-client>=1.6.0",
]

classifiers = [
    "License :: OSI Approved :: MIT License",
    "Programming Language :: Python :: 3",
    "Programming Language :: Python :: 3.11",
]

with open("README.md", "r", encoding="utf-8") as f:
    long_description = f.read()

setup(
    name=name,
    author=author,
    author_email=author_email,
    maintainer=author,
    maintainer_email=author_email,
    version=__version__,
    description=description,
    long_description=long_description,
    long_description_content_type=long_description_content_type,
    license=license,
    url=url,
    download_url=url,
    keywords=keywords,
    install_requires=install_requires,
    classifiers=classifiers,
    packages=find_packages(),
)
