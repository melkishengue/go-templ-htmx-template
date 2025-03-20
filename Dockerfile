
FROM golang:1.23.3 AS builder

WORKDIR /app

COPY go.mod go.sum ./

RUN go mod download

COPY . .

ARG BUILDTIME=""
ARG REVISION=""
ARG CONFIG_FILE=""

RUN CGO_ENABLED=0 GOOS=linux go build -o main .

FROM alpine:latest

RUN apk --no-cache add ca-certificates

WORKDIR /root/

COPY --from=builder /app/main .

ARG CONFIG_FILE=""
COPY --from=builder /app/${CONFIG_FILE} ./ 

ARG BUILDTIME=""
ARG REVISION=""

# To expose as env variables to the application
ENV BUILDTIME=${BUILDTIME}
ENV REVISION=${REVISION}

EXPOSE 3000

ENTRYPOINT ["./main"]
