FROM golang:1.12.5 as builder

ARG VERSION=unknown

#disable crosscompiling
ENV CGO_ENABLED=0

#compile linux only
ENV GOOS=linux

WORKDIR /dangkaka/leaderboard/
ADD . /dangkaka/leaderboard/
RUN go build -o app

FROM alpine
ARG VERSION=unknown
ENV VERSION $VERSION
RUN apk add --no-cache ca-certificates
ENV AWS_REGION ap-southeast-1
WORKDIR /app
COPY --from=builder /dangkaka/leaderboard/config/config.json /app/config/config.json
COPY --from=builder /dangkaka/leaderboard/app .

CMD ./leaderboard
