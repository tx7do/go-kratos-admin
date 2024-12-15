############################
# STEP 1 build optimized executable binary
############################

FROM golang:1.20 AS builder

COPY . /src

WORKDIR /src/app/${APP_RELATIVE_PATH}

RUN GOPROXY=https://goproxy.cn make build

############################
# STEP 2 build a small image
############################

FROM debian:stable-slim

ARG APP_RELATIVE_PATH

RUN apt update && \
    apt install -y --no-install-recommends ca-certificates netbase && \
    rm -rf /var/lib/apt/lists/ && \
    apt autoremove -y && \
    apt autoclean -y

COPY --from=builder /src/app/${APP_RELATIVE_PATH}/bin /app

WORKDIR /app

EXPOSE ${GRPC_PORT} ${REST_PORT}
VOLUME /data/conf

USER nonroot:nonroot

CMD ["./server", "-conf", "/data/conf"]
