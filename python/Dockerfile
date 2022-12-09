FROM python:3.11.0-slim AS builder

WORKDIR /tmp

RUN pip install poetry

COPY pyproject.toml poetry.lock* ./

RUN poetry export -f requirements.txt --output requirements.txt --without-hashes


FROM python:3.11.0-slim

ENV PYTHONUNBUFFERED 1

RUN useradd -s /bin/bash -m me && \
    groupadd app && \
    usermod -aG app me

RUN apt-get update && \
    apt-get upgrade -y && \
    rm -rf /var/lib/apt/lists/*

WORKDIR /src

COPY --from=builder /tmp/requirements.txt ./

RUN pip install --no-cache-dir --upgrade -r requirements.txt

COPY ./app /src/app

EXPOSE 80
CMD uvicorn app.main:app --host 0.0.0.0 --port 80