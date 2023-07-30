from setuptools import setup, find_packages
from yaylib import __version__

name = "yaylib"
author = "Qvco, Konn"
author_email = "nikola.desuga@gmail.com"
description = "同世代と趣味の通話コミュニティ - Yay! (イェイ) で、投稿やタイムラインの取得、リツイートやいいねの実行、フォローや投稿の検索など様々な機能を利用可能なAPIクライアントツールです。"
long_description_content_type = "text/markdown"
license = "MIT"
url = "https://github.com/qvco/yaylib"

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
    "httpx>=0.17.1",
    "Pillow>=9.3.0",
    "cryptography>=41.0.1",
    "websocket-client>=1.6.0",
]

classifiers = [
    "License :: OSI Approved :: MIT License",
    "Programming Language :: Python :: 3",
    "Programming Language :: Python :: 3.10",
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
