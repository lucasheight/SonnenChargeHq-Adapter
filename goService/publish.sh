#!/bin/bash
# docker build -t go-sonnen-chargehq-service:prod .

docker build -t go-sonnen-chargehq-service-amd64:prod --platform=linux/amd64 .

docker build -t go-sonnen-chargehq-service-arm64:prod --platform=linux/arm64 .

docker save --output ./dist/go-sonnen-amd64.tar go-sonnen-chargehq-service-amd64:prod

docker save --output ./dist/go-sonnen-arm64.tar go-sonnen-chargehq-service-arm64:prod

# docker build -t go-sonnen-chargehq-service-amd64:prod --output type=tar,dest=./dist/go-sonnen-amd64.tar --platform linux/amd64 .

# docker build -t go-sonnen-chargehq-service-arm64:prod --output type=tar,dest=./dist/go-sonnen-arm64.tar --platform linux/arm64 .