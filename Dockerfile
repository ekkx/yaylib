FROM python:3.11-slim

RUN apt-get update && apt-get install make

WORKDIR /yaylib

COPY Pipfile .

RUN pip install pipenv
RUN pipenv install --dev
RUN echo 'pipenv shell' >> ~/.bashrc

COPY . .
