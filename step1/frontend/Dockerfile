FROM node:16.13-slim

MAINTAINER Egor Zyuskin <ezyuskin@navicons.ru>

ENV PROJECT_DIR=/src/frontend
RUN mkdir -p $PROJECT_DIR
WORKDIR $PROJECT_DIR
COPY package.json package-lock.json $PROJECT_DIR
RUN npm install

COPY ./src/ $PROJECT_DIR/src/
COPY ./public/ $PROJECT_DIR/public/

RUN npm run build

FROM nginx:1.23-alpine-slim

ADD ./nginx.conf /etc/nginx/

COPY --from=0 /src/frontend/dist/ /var/www/

EXPOSE 80

WORKDIR /var/www/

CMD ["nginx", "-g", "daemon off;"]
