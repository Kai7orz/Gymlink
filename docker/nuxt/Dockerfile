  GNU nano 8.3                                  Dockerfile
# ---- Builder ----
FROM node:22 AS builder
WORKDIR /app
ENV NODE_ENV=development
COPY ./package.json .
COPY ./package-lock.json .
RUN npm install
COPY . .
RUN npm run build

# ---- Runner ----
FROM node:22 AS runner
ENV NODE_ENV=production
ENV PORT=3000
COPY --from=builder /app/.output ./.output
EXPOSE 3000
CMD ["node", ".output/server/index.mjs"]
