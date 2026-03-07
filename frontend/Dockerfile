FROM node:20-alpine AS build
WORKDIR /app

COPY package.json package-lock.json* ./
RUN npm ci --no-audit --no-fund

COPY . .

# 如果 posts 目录不存在，尝试从 /workspace/data/posts 复制
RUN if [ ! -d "posts" ] && [ -d "/workspace/data/posts" ]; then \
      cp -r /workspace/data/posts posts; \
    fi

RUN npm run build

FROM nginx:alpine
COPY nginx.conf /etc/nginx/conf.d/default.conf
COPY --from=build /app/dist /usr/share/nginx/html
