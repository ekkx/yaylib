FROM python:3.13-slim

WORKDIR /yaylib

# locales はドキュメント生成に必要
RUN apt-get update && \
    apt-get install -y make curl openssh-client locales && \
    echo "ja_JP.UTF-8 UTF-8" > /etc/locale.gen && \
    locale-gen && \
    update-locale LANG=ja_JP.UTF-8
RUN curl -sSL https://install.python-poetry.org | python3 - --version 2.1.3

ENV LANG ja_JP.UTF-8
ENV LANGUAGE ja_JP:ja
ENV LC_ALL ja_JP.UTF-8
ENV PATH /root/.local/bin:$PATH

COPY . .

RUN poetry install
RUN poetry lock
