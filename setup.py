from setuptools import setup, find_packages
from yaylib import __version__

name = "yaylib"
author = "qvco, Konn"
author_email = "nikola.desuga@gmail.com"
description = "This Python package provides an easy-to-use interface for accessing data from Yay!, a social networking platform. With this API Client, you can retrieve user profiles, posts, comments, and other content from Yay!, as well as perform common tasks like liking and commenting on posts. Compatible with Yay! version 3.16 and later. Please note that some API calls may be subject to rate limits or require authentication."
long_description_content_type = "text/markdown"
license = "MIT"
url = "https://github.com/qvco/yaylib"

keywords = [
    "yay",
    "yaylib",
    "api",
    "bot",
    "library",
    "wrapper"
]

install_requires = [
    "httpx>=0.17.1"
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
    # version=__version__,
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
