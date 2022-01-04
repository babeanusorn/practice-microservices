###########
## Build ##
###########
FROM golang:1.17-alpine AS build

WORKDIR /app

RUN mkdir /pkg

COPY ./pkg /pkg
COPY ./get /app

RUN go mod download
RUN go build -o /get

############
## Deploy ##
############
FROM alpine:3.15

WORKDIR /

COPY --from=build /get /get

EXPOSE 8080

ENTRYPOINT ["/get"]