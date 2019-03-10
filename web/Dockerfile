FROM node:10-alpine

WORKDIR /getting-started-docker-compose/web
COPY . .

RUN npm install --prod

CMD ["node", "index.js"]
