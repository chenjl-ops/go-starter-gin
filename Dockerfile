# Build stage
FROM golang:1.19.7 AS builder
ENV GOPROXY=https://proxy.golang.com.cn
#ENV GOPROXY=$GOPROXY

WORKDIR /go-starter-gin

COPY . .

RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o /go/bin/go-starter-gin ./cmd/app
#RUN bash ./scripts/build.sh

# Production stage
FROM amd64/alpine:latest

# Set System TimeZone
RUN apk add --no-cache tzdata
ENV TZ=Asia/Shanghai
RUN ln -snf /usr/share/zoneinfo/$TZ /etc/localtime && echo $TZ > /etc/timezone
COPY --from=builder /go/bin/go-starter-gin /chj/app/go-starter-gin
ENV WORKDIR=/chj/app/
CMD /chj/app/go-starter-gin start
