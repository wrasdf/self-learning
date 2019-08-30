FROM node:12.9-alpine

COPY package.json *yarn* /app/
RUN yarn install

COPY . /app
EXPOSE 8082
