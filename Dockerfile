ARG BUILDER_IMAGE=golang:1.17-buster
ARG BASE_IMAGE=bitnami/minideb:buster

FROM $BUILDER_IMAGE as builder

COPY . /root/go/src/app/

ARG BUILD_VERSION
ARG GOPROXY
ARG GOSUMDB=sum.golang.org

WORKDIR /root/go/src/app

ENV PATH="${PATH}:/usr/local/go/bin"
ENV BUILD_VERSION=$BUILD_VERSION
ENV GOPROXY=$GOPROXY
ENV GOSUMDB=$GOSUMDB

RUN go mod vendor

RUN CGO_ENABLED=0 GOOS=linux go build -a -v -ldflags "-X main.version=$(BUILD_VERSION)" -installsuffix cgo -o app ./server

FROM $BASE_IMAGE

WORKDIR /usr/app

COPY --from=builder /root/go/src/app/app /usr/app/app
COPY --from=builder /root/go/src/app/www /usr/app/www

LABEL authors="Fariz Rachman Yusuf"

LABEL org.opencontainers.image.authors="Fariz Rachman Yusuf <farizrachmanyusuf@gmail.com>"
LABEL org.opencontainers.image.title="swing"
LABEL org.opencontainers.image.description="Swing App"
LABEL org.opencontainers.image.vendor=""

EXPOSE 3000
EXPOSE 9090

ENTRYPOINT ["/usr/app/app"]
CMD ["grpc-gw-server", "--port1", "9090", "--port2", "3000", "--grpc-endpoint", ":9090"]
