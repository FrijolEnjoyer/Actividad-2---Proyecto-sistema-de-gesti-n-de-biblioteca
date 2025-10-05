# Build stage
FROM node:20-alpine as build
WORKDIR /app
COPY package.json package-lock.json* pnpm-lock.yaml* yarn.lock* ./
RUN npm ci || npm install
COPY . .
RUN npm run build

# Serve dev server for simplicity (Vite preview)
FROM node:20-alpine
WORKDIR /app
COPY --from=build /app/dist /app/dist
RUN npm i -g serve
ENV PORT=5173
EXPOSE 5173
CMD ["serve", "-s", "dist", "-l", "5173"]
