FROM node:16.20.0-alpine3.18 as builder
WORKDIR /usr/src/app
COPY . .
RUN npm install
ARG VUE_APP_API_URL
ENV VUE_APP_API_URL $VUE_APP_API_URL
RUN npm run build

FROM nginx:1.27.0
COPY --from=builder /usr/src/app/dist /usr/share/nginx/html
EXPOSE 8080

