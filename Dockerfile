FROM golang as build

ENV GOPROXY=https://goproxy.io

ADD . /bdas

WORKDIR /bdas

RUN CGO_ENABLED=1 GOOS=linux GOARCH=amd64 go build -o api_server

FROM jeanblanchard/alpine-glibc:3.7_2.30-r0

ENV REDIS_ADDR=""
ENV REDIS_PW=""
ENV REDIS_DB=""
ENV MysqlDSN=""
ENV GIN_MODE="release"
ENV PORT=3000

RUN echo "http://mirrors.aliyun.com/alpine/v3.7/main/" > /etc/apk/repositories && \
    apk update && \
    apk add ca-certificates && \
    echo "hosts: files dns" > /etc/nsswitch.conf && \
    mkdir -p /www/conf && \
    mkdir -p /www/static

WORKDIR /www

COPY --from=build /bdas/api_server /usr/bin/api_server
ADD ./conf /www/conf
ADD ./static /www/static

RUN chmod +x /usr/bin/api_server

ENTRYPOINT ["api_server"]