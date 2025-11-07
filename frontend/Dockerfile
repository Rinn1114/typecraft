# --- 构建阶段 ---
FROM node:20-alpine AS builder

# 设置工作目录
WORKDIR /app

# 复制 package.json 和 package-lock.json
COPY package.json package-lock.json ./
# 安装项目依赖
RUN npm install

# 复制所有源代码
COPY . .

# 执行构建命令，生成生产环境的静态文件
RUN npm run build

# --- 运行阶段 ---
FROM nginx:1.25-alpine

# 将构建阶段生成的静态文件复制到 Nginx 的默认网站根目录
COPY --from=builder /app/build /usr/share/nginx/html

# 声明容器将暴露 80 端口 (Nginx 默认端口)
EXPOSE 80

# 定义容器启动时要执行的命令 (在前台启动 Nginx)
CMD ["nginx", "-g", "daemon off;"]
