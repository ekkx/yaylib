from setuptools import setup, find_packages
from yaylib import __version__

with open("README.md", "r", encoding="utf-8") as f:
    long_description = f.read()

setup(
    name="yaylib",
    version=__version__,
    description="",
    long_description=long_description,
    long_description_content_type="text/markdown",
    license="MIT",
    url="https://github.com/qvco/yaylib",
    keywords=["yay", "yaylib", "api", "bot", "library", "wrapper"],
    install_requires=[
        "logging",
        "httpx",
        "tqdm",
    ],
    classifiers=[
        "License :: OSI Approved :: MIT License",
        "Programming Language :: Python :: 3",
        "Programming Language :: Python :: 3.9",
        "Programming Language :: Python :: 3.10",
        "Programming Language :: Python :: 3.11",
    ],
    packages=find_packages(),
)
