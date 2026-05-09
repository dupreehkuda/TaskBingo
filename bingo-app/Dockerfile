FROM node:20-alpine AS builder

WORKDIR /app

COPY package.json package-lock.json ./
RUN npm ci

COPY . .

# Vite inlines VITE_* env vars at build time. Pass them in via --build-arg
# so the final bundle bakes in the production API URL, not localhost.
ARG VITE_API_URL
ARG VITE_WS_URL
ARG VITE_WEB_URL
ENV VITE_API_URL=$VITE_API_URL \
    VITE_WS_URL=$VITE_WS_URL \
    VITE_WEB_URL=$VITE_WEB_URL

RUN npm run build

# ---------------------------------------------------------------------------
FROM node:20-alpine AS runner

WORKDIR /app

# Copy only what `node build/index.js` needs at runtime.
COPY --from=builder /app/build ./build
COPY --from=builder /app/package.json ./
COPY --from=builder /app/package-lock.json ./
COPY --from=builder /app/node_modules ./node_modules

ENV NODE_ENV=production \
    HOST=0.0.0.0 \
    PORT=3000

EXPOSE 3000

CMD ["node", "build"]
