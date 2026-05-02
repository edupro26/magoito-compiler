FROM golang:1.25-alpine AS builder

WORKDIR /app

COPY go.mod ./
RUN go mod download

COPY . .

RUN apk add --no-cache openjdk17-jre

RUN go generate ./...
RUN go build -o magoito ./cmd

FROM alpine:3.21

WORKDIR /app

COPY --from=builder /app/magoito /usr/local/bin/magoito
COPY --from=builder /app/*.mag .
COPY --from=builder /app/test ./test

RUN echo "clear" >> /root/.profile
RUN echo 'PS1="[magoito-compiler] \\w $ "' >> /root/.profile
CMD ["sh", "-l"]