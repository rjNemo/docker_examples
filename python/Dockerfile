FROM python:3-slim

ENV PYTHONUNBUFFERED 1

WORKDIR /app

COPY requirements.txt .

RUN pip install -r requirements.txt

COPY . .
CMD uvicorn main:app  --host 0.0.0.0 --port 80