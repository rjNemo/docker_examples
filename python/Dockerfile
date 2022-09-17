FROM python:3-slim

ENV PYTHONUNBUFFERED 1

WORKDIR /src

COPY requirements.txt ./

RUN pip install --no-cache-dir -r requirements.txt

COPY ./app /src/app

EXPOSE 80
CMD uvicorn app.main:app --host 0.0.0.0 --port 80