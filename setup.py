from setuptools import setup, find_packages
from yaybot import __version__


with open('README.md', 'r', encoding='utf-8') as f:
    readme = f.read()
long_description = readme

setup(
    name='YayBot',
    version=__version__,
    description='Unofficial API for Yay! (yay.space)',
    long_description=long_description,
    long_description_content_type='text/markdown',
    url='https://github.com/qualia-5w4/Yay-Bot',
    author='qualia-5w4',
    author_email='nikola.desuga@gmail.com',
    license='MIT',
    classifiers=[
        'License :: OSI Approved :: MIT License',
        'Programming Language :: Python :: 3',
        'Programming Language :: Python :: 3.9',
        'Programming Language :: Python :: 3.10',
        'Programming Language :: Python :: 3.11'

    ],
    keywords='yay api, unofficial, library, modules',
    packages=find_packages(),
    install_requires=[
        'requests',
        'beautifulsoup4',
        'tqdm',
        'huepy',
    ]
)
