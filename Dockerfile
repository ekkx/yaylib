FROM python:3.11-slim

RUN apt-get update && apt-get install make

WORKDIR /app

COPY Pipfile .
RUN pip install pipenv
RUN pipenv install
COPY . .
