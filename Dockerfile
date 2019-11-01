# compile
FROM golang:1.12.12 as build-env
ADD . /ws
WORKDIR /ws
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o main .


FROM alpine:3.10
RUN apk add --no-cache --update tzdata ca-certificates
ENV TZ 'Asia/Shanghai'
WORKDIR /app
COPY --from=build-env /ws/main /app/
EXPOSE 8080
CMD ["/app/main"]