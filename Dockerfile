FROM    golang:alpine3.19 AS dependencies
RUN     mkdir -p /opt/app-code
COPY    ./go.mod /opt/app-code
COPY    ./go.sum /opt/app-code
WORKDIR /opt/app-code
RUN     go mod download

FROM    dependencies AS builder
COPY    ./ /opt/app-code
WORKDIR /opt/app-code/microservices/ducco_products
RUN     go build -v -o ./dist/app *.go


FROM    alpine:3.14 AS dist
ARG     micro
COPY    --from=builder /opt/app-code/microservices/ducco_products/dist/. /opt/app-code
WORKDIR /opt/app-code
RUN     touch .env
CMD     ["./app"]
