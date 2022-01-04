###########
## Build ##
###########
FROM golang:1.17-alpine AS build

WORKDIR /app

RUN mkdir /pkg

COPY ./pkg /pkg
COPY ./delete /app

RUN go mod download
RUN go build -o /delete

############
## Deploy ##
############
FROM alpine:3.15

WORKDIR /

COPY --from=build /delete /delete

EXPOSE 8080

ENTRYPOINT ["/delete"]