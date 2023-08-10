FROM node:18 AS builder

WORKDIR /app

COPY bingo-app/package.json bingo-app/yarn.lock ./

RUN yarn install --frozen-lockfile

COPY ./bingo-app .

RUN yarn build

FROM node:18-alpine

WORKDIR /app

COPY --from=builder /app .

CMD ["yarn", "preview", "--host"]
