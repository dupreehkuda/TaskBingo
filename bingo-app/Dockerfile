FROM node:18.12.0-alpine3.16 AS builder
WORKDIR /app
COPY ./bingo-app/package*.json ./
RUN npm ci
COPY . .
RUN npm run build
RUN npm ci --prod

FROM node:18.12.0-alpine3.16
USER node:node
WORKDIR /app
COPY --from=builder --chown=node:node /app/build ./build
COPY --from=builder --chown=node:node /app/node_modules ./node_modules
COPY --chown=node:node ./bingo-app/package.json .
CMD ["node","build"]