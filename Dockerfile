FROM golang:1.10.2 as builder

ARG VERSION=unknown

#disable crosscompiling
ENV CGO_ENABLED=0

#compile linux only
ENV GOOS=linux

RUN curl -fsSL -o /usr/local/bin/dep https://github.com/golang/dep/releases/download/v0.4.1/dep-linux-amd64 && chmod +x /usr/local/bin/dep
ADD Gopkg.lock /go/src/github.com/dangkaka/leaderboard/
ADD Gopkg.toml /go/src/github.com/dangkaka/leaderboard/
WORKDIR /go/src/github.com/dangkaka/leaderboard/
RUN dep ensure -vendor-only -v
ADD . /go/src/github.com/dangkaka/leaderboard/
RUN go install ./...

FROM alpine
ARG VERSION=unknown
ENV VERSION $VERSION
RUN apk add --no-cache ca-certificates
ENV AWS_REGION ap-southeast-1
WORKDIR /app
COPY --from=builder /go/src/github.com/dangkaka/leaderboard/config/config.json /app/config/config.json
COPY --from=builder /go/bin/leaderboard .

CMD ./leaderboard
