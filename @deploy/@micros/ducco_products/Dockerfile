FROM    golang:alpine3.19 AS dependencies
RUN     mkdir -p /ducco/source-code
COPY    ./go.mod /ducco/source-code
COPY    ./go.mod /ducco/source-code
WORKDIR /ducco/source-code
RUN     go mod download

FROM    dependencies AS builder
COPY    ./ /ducco/source-code
WORKDIR /ducco/source-code/microservices/ducco_products
RUN     go build -v -o ./dist/app *.go


FROM    alpine:3.14 AS dist
ARG     micro
COPY    --from=builder /ducco/source-code/microservices/ducco_products/dist/. /ducco/source-code
WORKDIR /ducco/source-code
RUN     touch .env
CMD     ["./app"]
