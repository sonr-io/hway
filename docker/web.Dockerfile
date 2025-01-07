# Builder
FROM node:20-bullseye as builder

# Support custom branches of the react-sdk and js-sdk. This also helps us build
# images of element-web develop.
ARG USE_CUSTOM_SDKS=false
ARG REACT_SDK_REPO="https://github.com/matrix-org/matrix-react-sdk.git"
ARG REACT_SDK_BRANCH="master"
ARG JS_SDK_REPO="https://github.com/matrix-org/matrix-js-sdk.git"
ARG JS_SDK_BRANCH="master"
ARG GIT_SOURCE_REPO="https://github.com/element-hq/element-web.git"
ARG GIT_SOURCE_BRANCH="master"

RUN apt-get update && apt-get install -y git dos2unix
RUN git clone ${GIT_SOURCE_REPO} src
WORKDIR /src
RUN git checkout ${GIT_SOURCE_BRANCH}

RUN dos2unix /src/scripts/docker-link-repos.sh && bash /src/scripts/docker-link-repos.sh
RUN yarn --network-timeout=200000 install

RUN dos2unix /src/scripts/docker-package.sh && bash /src/scripts/docker-package.sh

# Copy the config now so that we don't create another layer in the app image
COPY config/element.json /src/webapp/config.json

# App
FROM nginx:alpine-slim

COPY --from=builder /src/webapp /app

# Override default nginx config
COPY config/nginx.conf /etc/nginx/conf.d/default.conf

RUN rm -rf /usr/share/nginx/html \
  && ln -s /app /usr/share/nginx/html

# Run nginx
CMD ["nginx", "-g", "daemon off;"]
