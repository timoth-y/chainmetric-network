# This image is a microservice in golang for the Degree chaincode
FROM golang:1.15-alpine AS build

WORKDIR /go/src/github.com/timoth-y/iot-blockchain-contracts/
COPY .. .

# Build application
RUN go build -o chaincode -v ./src/devices

# Production ready image
FROM alpine:3.11 as prod

COPY --from=build /go/src/github.com/timoth-y/iot-blockchain-contracts/chaincode /app/chaincode

USER 1000

WORKDIR /app
CMD ./chaincode
