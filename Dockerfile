FROM python:3.11-slim

WORKDIR /yaylib

RUN apt-get update && apt-get install -y make curl
RUN curl -sSL https://install.python-poetry.org | python3 - --version 1.8.2

ENV PATH /root/.local/bin:$PATH

COPY . .

RUN poetry install
RUN poetry lock

RUN ./poetry_plugins.sh
