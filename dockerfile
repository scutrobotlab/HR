FROM node:lts as build-env
RUN npm install --global npm@6.12.1
WORKDIR /app
COPY package.json package-lock.json /app/
RUN npm ci

FROM build-env as build
COPY . /app
RUN npm run build

FROM nginx:stable as runtime
COPY nginx/nginx.conf /etc/nginx/
COPY --from=build /app/dist/ /usr/share/nginx/html