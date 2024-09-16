FROM node:16.20.2-alpine3.18 as builder

WORKDIR /app

COPY package*.json ./

RUN npm install

COPY . .

RUN npm run build


# FROM node:16.20.2-alpine3.18 as runner

# WORKDIR /app

# COPY --from=builder /app/dist ./dist

# EXPOSE 3000

# CMD ["node", "dist/index.js"]
FROM nginx:1.25.4-alpine3.18

COPY ./nginx.conf /etc/nginx/conf.d/default.conf
COPY --from=builder /app/dist /var/www/html/

EXPOSE 3000

ENTRYPOINT ["nginx","-g","daemon off;"]