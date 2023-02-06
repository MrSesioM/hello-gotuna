FROM golang:alpine AS build
WORKDIR /usr/src/app
COPY . .
RUN go build examples/fullapp/cmd/main.go

FROM alpine:3.17
WORKDIR /opt/gotuna
COPY --from=build /usr/src/app .
EXPOSE 8888
CMD ["./main"]
