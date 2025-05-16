FROM python:3.13-slim

WORKDIR /yaylib

RUN apt-get update && apt-get install -y make curl openssh-client
RUN curl -sSL https://install.python-poetry.org | python3 - --version 2.1.3

ENV PATH /root/.local/bin:$PATH

COPY . .

RUN poetry install
RUN poetry lock
