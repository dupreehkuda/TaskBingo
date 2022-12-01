FROM node:14-alpine

WORKDIR /usr/src/app

COPY ./BingoApp/rollup.config.js ./
COPY ./BingoApp/package*.json ./

RUN npm install

COPY ./BingoApp/src ./src
COPY ./BingoApp/public ./public

RUN npm run-script build

EXPOSE 8080

ENV HOST=0.0.0.0

CMD [ "npm", "start" ]