###########
## Build ##
###########
FROM golang:1.17-alpine AS build

WORKDIR /app

RUN mkdir /pkg

COPY ./pkg /pkg
COPY ./create /app

RUN go mod download
RUN go build -o /create

############
## Deploy ##
############
FROM alpine:3.15

WORKDIR /

COPY --from=build /create /create

EXPOSE 8080

ENTRYPOINT ["/create"]