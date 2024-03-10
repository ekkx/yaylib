from os import path
from setuptools import setup, find_packages

from yaylib import (
    __title__,
    __author__,
    __license__,
    __copyright__,
    __version__,
)

here = path.abspath(path.dirname(__file__))

with open(path.join(here, "README.md"), encoding="utf-8") as f:
    long_description = f.read()

author_email = "nikola.desuga@gmail.com"
description = "同世代と趣味の通話コミュニティ - Yay! (イェイ) で投稿やタイムラインの取得、リツイートやいいねの実行、フォローや投稿の検索など様々な機能をPythonプログラムから利用可能なAPIクライアントツールです。"
long_description_content_type = "text/markdown"
url = "https://github.com/ekkx/yaylib"

keywords = [
    "yay",
    "yaylib",
    "api",
    "bot",
    "tool",
    "client",
    "library",
    "wrapper",
    "ボット",
    "ライブラリ",
    "ツール",
]

install_requires = [
    "httpx[http2,socks]==0.25.2",
    "Pillow==9.3.0",
    "cryptography==41.0.1",
    "websocket-client==1.6.0",
    "PyJWT==2.8.0",
]

classifiers = [
    "License :: OSI Approved :: MIT License",
    "Programming Language :: Python :: 3",
    "Programming Language :: Python :: 3.10",
    "Programming Language :: Python :: 3.11",
]

setup(
    name=__title__,
    author=__author__,
    author_email=author_email,
    maintainer=__author__,
    maintainer_email=author_email,
    version=__version__,
    description=description,
    long_description=long_description,
    long_description_content_type=long_description_content_type,
    license=__license__,
    url=url,
    download_url=url,
    keywords=keywords,
    install_requires=install_requires,
    classifiers=classifiers,
    packages=find_packages(),
)
