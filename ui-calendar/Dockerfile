# pull official base image
FROM node:18.10.0-alpine as builder

# set working directory
WORKDIR /app

# install app dependencies
COPY package.json ./
COPY package-lock.json ./
RUN npm install

# add app
COPY . ./

# build time args used as node build envs
# ARG REACT_APP_API_ENDPOINT
# ENV REACT_APP_API_ENDPOINT $REACT_APP_API_ENDPOINT

# fixes JS out of memory error
ENV NODE_OPTIONS --max-old-space-size=3072

# build project
RUN npm run build

# Start a new stage from scratch
FROM nginx:1.23.2-alpine

COPY --from=builder /app/nginx.conf /etc/nginx/conf.d/default.conf
COPY --from=builder /app/build /usr/share/nginx/html

EXPOSE 80

CMD ["nginx", "-g", "daemon off;"] 
