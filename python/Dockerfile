FROM python:3.11.0-slim

ENV PYTHONUNBUFFERED 1

RUN useradd -s /bin/bash -m me && \
    groupadd app && \
    usermod -aG app me

WORKDIR /src

RUN apt-get update && \
    apt-get upgrade -y && \
    rm -rf /var/lib/apt/lists/*

COPY requirements.txt ./

RUN pip install --no-cache-dir -r requirements.txt

COPY ./app /src/app

EXPOSE 80
CMD uvicorn app.main:app --host 0.0.0.0 --port 80