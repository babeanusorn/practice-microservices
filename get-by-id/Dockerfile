###########
## Build ##
###########
FROM golang:1.17-alpine AS build

WORKDIR /app

RUN mkdir /pkg

COPY ./pkg /pkg
COPY ./get-by-id /app

RUN go mod download
RUN go build -o /get-by-id

############
## Deploy ##
############
FROM alpine:3.15

WORKDIR /

COPY --from=build /get-by-id /get-by-id

EXPOSE 8080

ENTRYPOINT ["/get-by-id"]