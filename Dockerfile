FROM golang:1.18-alpine AS builder

RUN apk update && apk add nmap
RUN mkdir /app
WORKDIR /app

COPY go.mod .
COPY go.sum .
RUN go mod download
COPY . .

ENV CGO_ENABLED=0
RUN go build cmd/main.go


FROM scratch
COPY --from=builder /app/main /
ENTRYPOINT ["/main"]