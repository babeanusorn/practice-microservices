###########
## Build ##
###########
FROM golang:1.17-alpine AS build

WORKDIR /app

RUN mkdir /pkg

COPY ./pkg /pkg
COPY ./update /app

RUN go mod download
RUN go build -o /update

############
## Deploy ##
############
FROM alpine:3.15

WORKDIR /

COPY --from=build /update /update

EXPOSE 8080

ENTRYPOINT ["/update"]