FROM python:3.11-slim

WORKDIR /yaylib

RUN apt-get update && apt-get install make curl -y
RUN curl -sSL https://install.python-poetry.org | python3 -

ENV PATH /root/.local/bin:$PATH

COPY pyproject.toml* poetry.lock* ./

# RUN poetry config virtualenvs.in-project true
RUN poetry install

COPY . .
