FROM node:18.13 AS builder

WORKDIR /app

COPY package.json package-lock.json ./

RUN npm ci

COPY . .

RUN npm run build

FROM node:18.13-alpine

WORKDIR /app

COPY --from=builder /app .

CMD ["npm", "run", "preview", "--host"]