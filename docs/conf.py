# Configuration file for the Sphinx documentation builder.
#
# For the full list of built-in configuration values, see the documentation:
# https://www.sphinx-doc.org/en/master/usage/configuration.html

# -- Project information -----------------------------------------------------
# https://www.sphinx-doc.org/en/master/usage/configuration.html#project-information

import os
import sys

sys.path.insert(0, os.path.abspath(".."))

import yaylib

project = yaylib.__title__
copyright = f"Copyright © 2023, {yaylib.__author__}"
author = yaylib.__author__
release = yaylib.__version__

# -- General configuration ---------------------------------------------------
# https://www.sphinx-doc.org/en/master/usage/configuration.html#general-configuration

extensions = [
    "sphinx.ext.autodoc",
    "sphinx.ext.intersphinx",
    "sphinx.ext.extlinks",
    "sphinx.ext.todo",
    "sphinx.ext.viewcode",
    "sphinx.ext.autosummary",
    "sphinx_copybutton",
    "sphinx_design",
]

templates_path = ["_templates"]
exclude_patterns = ["_build", "Thumbs.db", ".DS_Store"]


# -- Options for HTML output -------------------------------------------------
# https://www.sphinx-doc.org/en/master/usage/configuration.html#options-for-html-output

html_theme = "shibuya"
html_static_path = ["_static"]

html_theme_options = {
    "accent_color": "blue",
    "github_url": "https://github.com/ekkx/yaylib",
    "nav_links": [
        {
            "title": "Examples",
            "url": "writing",
            "children": [
                {
                    "title": "Admonitions",
                    "url": "writing/admonition",
                    "summary": "Bring the attention of readers",
                },
                {
                    "title": "Code Blocks",
                    "url": "writing/code",
                    "summary": "Display code with highlights",
                },
                {
                    "title": "Autodoc",
                    "url": "writing/api",
                    "summary": "API documentation automatically",
                },
                {
                    "title": "Jupyter Notebook",
                    "url": "extensions/nbsphinx",
                    "summary": "Rendering .ipynb files",
                },
            ],
        },
        {"title": "Sponsor me", "url": "https://github.com/sponsors/ekkx"},
    ],
}
