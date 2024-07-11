FROM node:20-alpine as build
WORKDIR /usr/app
COPY package.json .
RUN npm install
COPY . .
EXPOSE 3000
CMD ["sh", "entrypoint.sh"]