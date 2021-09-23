FROM node:lts-alpine AS build

WORKDIR /app
COPY package.json .
COPY yarn.lock .
RUN yarn install
COPY . .
RUN yarn build


FROM nginx:latest

COPY nginx/nginx.conf /etc/nginx/nginx.conf
COPY  --from=build /app/build/ /usr/share/nginx/html
