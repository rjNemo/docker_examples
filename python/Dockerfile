FROM python:3.10-slim

ENV PYTHONUNBUFFERED 1

WORKDIR /app

COPY Pipfile* .
RUN pip install pipenv
RUN pipenv install --system --deploy --ignore-pipfile

COPY . .
CMD uvicorn main:app  --host 0.0.0.0 --port 80