FROM python:3.11-slim

WORKDIR /yaylib

RUN apt-get update && apt-get install make curl -y
RUN curl -sSL https://install.python-poetry.org | python3 - --version 1.8.2

ENV PATH /root/.local/bin:$PATH

COPY . .

# RUN poetry config virtualenvs.in-project true
RUN poetry lock
RUN poetry install

RUN ./poetry_plugins.sh
