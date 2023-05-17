from setuptools import setup, find_packages
from yaylib import __version__

with open("README", "r", encoding="utf-8") as f:
    long_description = f.read()

setup(
    name="yaylib",
    version=__version__,
    description="",
    long_description=long_description,
    long_description_content_type="text/markdown",
    license="MIT",
    url="https://github.com/qualia-5w4/yaylib",
    keywords=["yay", "api", "bot", "library", "wrapper"],
    install_requires=[
        "asyncio",
        "fake_useragent",
        "logging",
        "huepy",
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
