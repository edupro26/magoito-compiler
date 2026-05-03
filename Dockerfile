FROM golang:1.25-alpine AS builder

WORKDIR /app

COPY go.mod ./
RUN go mod download

COPY . .

RUN apk add --no-cache openjdk17-jre

RUN go generate ./...
RUN go build -o magoito ./cli

RUN go version | awk '{print $3}' > /go-version.txt

FROM alpine:3.21

WORKDIR /app

COPY --from=builder /app/magoito /usr/local/bin/magoito
COPY --from=builder /app/*.mag .
COPY --from=builder /app/test ./test

COPY --from=builder /go-version.txt /go-version.txt

RUN echo "printf '\033[H\033[2J\033[3J'" >> /root/.profile && \
    echo 'echo "🚀 Magoito Compiler CLI, built with Go, version: $(cat /go-version.txt)"' >> /root/.profile && \
    echo "" >> /root/.profile && \
    echo 'echo "This tool was developed as part of the Compilation Techniques course at Faculty of Sciences of the University of Lisbon."' >> /root/.profile && \
    echo "" >> /root/.profile && \
    echo 'echo "Author: Eduardo Proença"' >> /root/.profile && \
    echo "" >> /root/.profile && \
    echo "echo \"Type 'magoito --help' for more information, or run 'magoito run --example 1' to run an example\"" >> /root/.profile && \
    echo "" >> /root/.profile && \
    echo 'PS1="[magoito-compiler] \\w $ "' >> /root/.profile

CMD ["sh", "-l"]