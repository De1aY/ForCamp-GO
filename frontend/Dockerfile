FROM node:latest as build
WORKDIR /home
COPY package.json ./

RUN npm install
COPY . ./
ARG backend_endpoint=http://127.0.0.1:5000
ENV VUE_APP_BACKEND_URL=$backend_endpoint
RUN npm run build

FROM nginx:1.15-alpine
WORKDIR /home
COPY --from=build /home/dist /home/dist

# RUN apk add --no-cache apache2-utils
RUN rm /etc/nginx/conf.d/default.conf
ADD wplay.conf /etc/nginx/conf.d

EXPOSE 80
CMD ["nginx", "-g", "daemon off;"]
