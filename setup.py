from setuptools import setup, find_packages
from yaybot import __version__


with open('README.md', 'r', encoding='utf-8') as f:
    readme = f.read()
long_description = readme

setup(
    name='yaybot',
    version=__version__,
    description='This Python package provides an easy-to-use interface for accessing data from Yay!, a social networking platform. With this API, you can retrieve user profiles, posts, comments, and other content from Yay!, as well as perform common tasks like liking and commenting on posts. Compatible with Yay! version 3.0 and later. Please note that some API calls may be subject to rate limits or require authentication.',
    long_description=long_description,
    long_description_content_type='text/markdown',
    url='https://github.com/qvco/Yay-Bot',
    author='qvco',
    author_email='nikola.desuga@gmail.com',
    license='MIT',
    classifiers=[
        'License :: OSI Approved :: MIT License',
        'Programming Language :: Python :: 3',
        'Programming Language :: Python :: 3.9',
        'Programming Language :: Python :: 3.10',
        'Programming Language :: Python :: 3.11',
    ],
    keywords='yay, api, unofficial, library, modules',
    packages=find_packages(),
    install_requires=[
        'requests',
        'beautifulsoup4',
        'tqdm',
        'huepy',
        'python-dotenv',
        'Pillow',
    ]
)
