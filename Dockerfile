FROM golang:alpine AS builder

WORKDIR /speakbuddy

COPY . .

RUN go mod download

RUN go mod tidy

RUN go build speakbuddy

ENTRYPOINT [ "./speakbuddy" ]

# Minimize the image size
FROM alpine:latest AS release

WORKDIR /speakbuddy

RUN apk upgrade -U \ 
    && apk add ca-certificates ffmpeg

COPY --from=builder /speakbuddy/speakbuddy .

ENTRYPOINT [ "./speakbuddy" ]

