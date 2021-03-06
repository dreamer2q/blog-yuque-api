FROM alpine

EXPOSE 8008

ENV TIME_ZONE=Asia/Shanghai
RUN apk add tzdata && cp /usr/share/zoneinfo/Asia/Shanghai /etc/localtime \
    && echo "Asia/Shanghai" > /etc/timezone \
    && apk del tzdata

WORKDIR /root/

COPY main app

CMD ["./app"]
