FROM docker.io/golang:1.20-alpine AS base

WORKDIR /build
RUN mkdir -p /out
RUN apk --update --no-cache add build-base git
ARG GIT_SOURCE_REPO="https://github.com/matrix-org/sliding-sync.git"
ARG GIT_SOURCE_BRANCH="main"

RUN git clone ${GIT_SOURCE_REPO} .

ARG BINARYNAME=syncv3
RUN go build -o /out/syncv3 "./cmd/$BINARYNAME"

FROM alpine:3.17

RUN apk --update --no-cache add curl
COPY --from=base /out/* /usr/bin/

ENV SYNCV3_SERVER="https://sonr.chat"
ENV SYNCV3_DB="postgresql://neondb_owner:4zHmfZpEB6wQ@ep-square-wildflower-a52z3tuw.us-east-2.aws.neon.tech/neondb?sslmode=require"
ENV SYNCV3_SECRET="d1caeb5d86bee467e51a8650aa44e21ad68103d9495872015fbd36c0b9de4957"
ENV SYNCV3_BINDADDR="0.0.0.0:8008"

EXPOSE 8008

WORKDIR /usr/bin
# It would be nice if the binary we exec was called $BINARYNAME here, but build args
# aren't expanded in ENTRYPOINT directives. Instead, we always call the output binary
# "syncv3". (See https://github.com/moby/moby/issues/18492)
ENTRYPOINT /usr/bin/syncv3
