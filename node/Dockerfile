FROM node:lts-alpine

WORKDIR /app
COPY package.json .
COPY yarn.lock .
RUN yarn install
COPY . .
CMD node index.js
